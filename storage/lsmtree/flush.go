package lsmtree

import (
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"

	protobuf "google.golang.org/protobuf/proto"

	"github.com/maxpoletaev/kv/internal/bloom"
	"github.com/maxpoletaev/kv/internal/opengroup"
	"github.com/maxpoletaev/kv/internal/protoio"
	"github.com/maxpoletaev/kv/storage/lsmtree/proto"
)

type flushOpts struct {
	prefix    string
	tableID   int64
	indexGap  int64
	useMmap   bool
	bloomProb float64
}

// flushToDisk writes the contents of the memtable to disk and returns an SSTable
// that can be used to read the data. The memtable must be closed before calling
// this function to guarantee that it is not modified while the flush. The parameters
// of the bloom filter are calculated based on the number of entries in the memtable.
func flushToDisk(mem *Memtable, opts flushOpts) (*SSTable, error) {
	og := opengroup.New()
	defer og.CloseAll()

	info := &SSTableInfo{
		ID:         opts.tableID,
		NumEntries: int64(mem.Len()),
		IndexFile:  fmt.Sprintf("sst-%d.index", opts.tableID),
		DataFile:   fmt.Sprintf("sst-%d.data", opts.tableID),
		BloomFile:  fmt.Sprintf("sst-%d.bloom", opts.tableID),
	}

	indexFile := og.Open(filepath.Join(opts.prefix, info.IndexFile), os.O_CREATE|os.O_WRONLY, 0o644)
	bloomFile := og.Open(filepath.Join(opts.prefix, info.BloomFile), os.O_CREATE|os.O_WRONLY, 0o644)
	dataFile := og.Open(filepath.Join(opts.prefix, info.DataFile), os.O_CREATE|os.O_WRONLY, 0o644)
	if err := og.Err(); err != nil {
		return nil, fmt.Errorf("failed to open files: %w", err)
	}

	bf := bloom.NewWithProbability(mem.Len(), opts.bloomProb)
	indexWriter := protoio.NewWriter(indexFile)
	dataWriter := protoio.NewWriter(dataFile)

	var lastOffset int64

	for it := mem.entries.Scan(); it.HasNext(); {
		key, entry := it.Next()

		bf.Add([]byte(key))

		offset := dataWriter.Offset()

		if _, err := dataWriter.Append(entry); err != nil {
			return nil, fmt.Errorf("failed to write entry: %w", err)
		}

		// The index file is sparse, so we only write an index entry if the gap between
		// the current offset and the last offset is larger than the threshold.
		if lastOffset == 0 || offset-lastOffset >= opts.indexGap {
			indexEntry := &proto.IndexEntry{
				DataOffset: int64(offset),
				Key:        key,
			}

			if _, err := indexWriter.Append(indexEntry); err != nil {
				return nil, fmt.Errorf("failed to write index entry: %w", err)
			}

			lastOffset = offset
		}
	}

	bloomData, err := protobuf.Marshal(&proto.BloomFilter{
		Crc32:     crc32.ChecksumIEEE(bf.Bytes()),
		NumHashes: int32(bf.Hashes()),
		NumBytes:  int32(bf.Size()),
		Data:      bf.Bytes(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bloom filter: %w", err)
	}

	if _, err := bloomFile.Write(bloomData); err != nil {
		return nil, fmt.Errorf("failed to write bloom filter: %w", err)
	}

	// Use the size of the data file as the size of the table,
	// as it includes both the size of the keys and the values.
	info.Size = dataWriter.Offset()

	// Open the flushed table for reading. This should be done before discarding
	// the memtable as we want to ensure that the table is readable.
	sst, err := OpenTable(info, opts.prefix, opts.useMmap)
	if err != nil {
		_ = og.RemoveAll() // Cleanup so that we don???t generate garbage in case of error.
		return nil, fmt.Errorf("failed to open table: %w", err)
	}

	return sst, nil
}
