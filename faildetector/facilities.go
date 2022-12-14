package faildetector

import (
	"github.com/maxpoletaev/kv/clust"
	"github.com/maxpoletaev/kv/membership"
)

type memberRegistry interface {
	Self() membership.Member
	Members() []membership.Member
	SetStatus(id membership.NodeID, status membership.Status) (membership.Status, error)
}

type connectionRegistry interface {
	Get(id membership.NodeID) (clust.Conn, error)
}
