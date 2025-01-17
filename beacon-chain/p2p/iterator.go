package p2p

import (
	"context"

	"github.com/theQRL/zond/p2p/znode"
)

// filterNodes wraps an iterator such that Next only returns nodes for which
// the 'check' function returns true. This custom implementation also
// checks for context deadlines so that in the event the parent context has
// expired, we do exit from the search rather than  perform more network
// lookups for additional peers.
func filterNodes(ctx context.Context, it znode.Iterator, check func(*znode.Node) bool) znode.Iterator {
	return &filterIter{ctx, it, check}
}

type filterIter struct {
	context.Context
	znode.Iterator
	check func(*znode.Node) bool
}

// Next looks up for the next valid node according to our
// filter criteria.
func (f *filterIter) Next() bool {
	for f.Iterator.Next() {
		if f.Context.Err() != nil {
			return false
		}
		if f.check(f.Node()) {
			return true
		}
	}
	return false
}
