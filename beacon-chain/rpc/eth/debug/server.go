// Package debug defines a gRPC beacon service implementation,
// following the official API standards https://ethereum.github.io/beacon-apis/#/.
// This package includes the beacon and config endpoints.
package debug

import (
	"github.com/theQRL/zond/beacon-chain/blockchain"
	"github.com/theQRL/zond/beacon-chain/db"
	"github.com/theQRL/zond/beacon-chain/rpc/statefetcher"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum Beacon Chain.
type Server struct {
	BeaconDB              db.ReadOnlyDatabase
	HeadFetcher           blockchain.HeadFetcher
	StateFetcher          statefetcher.Fetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	ForkFetcher           blockchain.ForkFetcher
}