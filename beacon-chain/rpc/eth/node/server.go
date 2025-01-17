// Package node defines a gRPC node service implementation, providing
// useful endpoints for checking a node's sync status, peer info,
// genesis data, and version information.
package node

import (
	"github.com/theQRL/zond/beacon-chain/blockchain"
	"github.com/theQRL/zond/beacon-chain/db"
	"github.com/theQRL/zond/beacon-chain/p2p"
	"github.com/theQRL/zond/beacon-chain/sync"
	"google.golang.org/grpc"
)

// Server defines a server implementation of the gRPC Node service,
// providing RPC endpoints for verifying a beacon node's sync status, genesis and
// version information.
type Server struct {
	SyncChecker           sync.Checker
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	Server                *grpc.Server
	BeaconDB              db.ReadOnlyDatabase
	PeersFetcher          p2p.PeersProvider
	PeerManager           p2p.PeerManager
	MetadataProvider      p2p.MetadataProvider
	GenesisTimeFetcher    blockchain.TimeFetcher
	HeadFetcher           blockchain.HeadFetcher
}
