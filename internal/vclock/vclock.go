package vclock

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maxpoletaev/kv/internal/generic"
)

// Causality is a type that represents the causality relationship between two vectors.
type Causality int

const (
	Before Causality = iota + 1
	Concurrent
	After
	Equal
)

func (c Causality) String() string {
	switch c {
	case Before:
		return "Before"
	case Concurrent:
		return "Concurrent"
	case After:
		return "After"
	case Equal:
		return "Equal"
	default:
		return ""
	}
}

type V map[uint32]uint32

// Vector represents a vector clock, which is a mapping of node IDs to clock values.
// The clock value is a monotonically increasing counter that is incremented every time
// the node makes an update. The clock value is used to determine the causality relationship
// between two events. If the clock value of a node is greater than the clock value of
// another node, it means that the first event happened after the second event.
// The implementation uses a 32-bit unsigned integer to store the clock value,
// which means that the clock value will roll over to zero after it reaches the maximum
// value of 2^32-1. The implementation keeps track of such rollover events, so that
// the causality relationship between two events can be determined even if the clock
// value has rolled over.
type Vector struct {
	clocks    V
	rollovers map[uint32]bool
}

// New returns a new vector clock. If the given values are not empty, the new vector
// clock is initialized with the given values. Otherwise, the new vector clock is
// initialized with an empty map.
func New(values ...V) *Vector {
	if len(values) > 1 {
		panic("too many arguments")
	}

	var clocks V
	if len(values) == 1 {
		clocks = values[0]
	} else {
		clocks = make(V)
	}

	return &Vector{
		clocks:    clocks,
		rollovers: make(map[uint32]bool, len(clocks)),
	}
}

// Get returns the clock value for the given node ID.
func (v *Vector) Get(id uint32) uint32 {
	return v.clocks[id]
}

// Rollover inverts the rollover flag for the given node ID.
func (v *Vector) Rollover(id uint32) {
	v.rollovers[id] = !v.rollovers[id]
}

// Update increments the clock value for the given node ID.
// If the clock value has rolled over, the rollover flag is inverted.
func (vc *Vector) Update(id uint32) {
	old := vc.clocks[id]
	vc.clocks[id]++

	if old > vc.clocks[id] {
		// Clock value has rolled over.
		vc.rollovers[id] = !vc.rollovers[id]
	}
}

// Clone returns a copy of the vector clock. The copy is a deep copy,
// so that the original vector clock can be modified without affecting
// the copy.
func (v *Vector) Clone() *Vector {
	newvec := &Vector{
		clocks:    make(map[uint32]uint32, len(v.clocks)),
		rollovers: make(map[uint32]bool, len(v.rollovers)),
	}

	generic.MapCopy(v.clocks, newvec.clocks)
	generic.MapCopy(v.rollovers, newvec.rollovers)

	return newvec
}

// String returns a string representation of the vector clock.
// The string representation is a comma-separated list of key=value pairs, where the
// key is the node ID and the value is the clock value: {1=1, 2=2}. If the clock value
// has rolled over, the value is prefixed with an exclamation mark: {1=1, 2=!2}.
func (v Vector) String() string {
	b := strings.Builder{}

	b.WriteString("{")

	keys := generic.MapKeys(v.clocks)

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for i, key := range keys {
		if i > 0 {
			b.WriteString(", ")
		}

		b.WriteString(fmt.Sprint(key))
		b.WriteString("=")

		if v.rollovers[key] {
			b.WriteString("!")
		}

		b.WriteString(fmt.Sprint(v.clocks[key]))
	}

	b.WriteString("}")

	return b.String()
}

// Compare returns the causality relationship between two vectors.
// Compare(a, b) == After means that a happened after b, and so on.
// Comparing values that have rolled over is tricky, so the implementation
// uses the following rules: if the clock value of a node is greater than
// the clock value of another node, and the rollover flags are different,
// it means that the value has wrapped around and we need to invert the
// comparison. For example, if a clock value is 2^32-1 and the other clock
// value is 0, and the rollover flags are different, it means that the clock
// value of the first node has wrapped around and the second node has not.
// In this case, the first node is considered to be less than the second node.
func Compare(a, b *Vector) Causality {
	var greater, less bool

	for _, key := range generic.MapKeys(a.clocks, b.clocks) {
		// If the rollover flags are different, it means that the value
		// has wrapped around and we need to invert the comparison.
		wrapped := a.rollovers[key] != b.rollovers[key]

		if a.clocks[key] > b.clocks[key] {
			if !wrapped {
				greater = true
			} else {
				less = true
			}
		} else if a.clocks[key] < b.clocks[key] {
			if !wrapped {
				less = true
			} else {
				greater = true
			}
		}
	}

	switch {
	case greater && !less:
		return After
	case less && !greater:
		return Before
	case !less && !greater:
		return Equal
	default:
		return Concurrent
	}
}

// IsEqual returns true if the two vectors are equal.
func IsEqual(a, b *Vector) bool {
	return Compare(a, b) == Equal
}

// Merge returns a new vector that is the result of merging two vectors.
// The merge operation is commutative and associative, so that
// Merge(a, Merge(b, c)) == Merge(Merge(a, b), c).
func Merge(a, b *Vector) *Vector {
	keys := generic.MapKeys(a.clocks, b.clocks)

	clock := &Vector{
		clocks:    make(map[uint32]uint32, len(keys)),
		rollovers: make(map[uint32]bool, len(keys)),
	}

	for _, key := range keys {
		// TODO: this is a bit ugly, but it works.
		if a.rollovers[key] == b.rollovers[key] {
			if a.clocks[key] > b.clocks[key] {
				clock.clocks[key] = a.clocks[key]
				if rollover, ok := a.rollovers[key]; ok {
					clock.rollovers[key] = rollover
				}
			} else {
				clock.clocks[key] = b.clocks[key]
				if rollover, ok := b.rollovers[key]; ok {
					clock.rollovers[key] = rollover
				}
			}
		} else {
			if a.clocks[key] < b.clocks[key] {
				clock.clocks[key] = a.clocks[key]
				if rollover, ok := a.rollovers[key]; ok {
					clock.rollovers[key] = rollover
				}
			} else {
				clock.clocks[key] = b.clocks[key]
				if rollover, ok := b.rollovers[key]; ok {
					clock.rollovers[key] = rollover
				}
			}
		}
	}

	return clock
}
