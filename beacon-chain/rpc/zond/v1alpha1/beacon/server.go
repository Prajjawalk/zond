// Package beacon defines a gRPC beacon service implementation, providing
// useful endpoints for checking fetching chain-specific data such as
// blocks, committees, validators, assignments, and more.
package beacon

import (
	"context"
	"time"

	"github.com/theQRL/zond/beacon-chain/blockchain"
	"github.com/theQRL/zond/beacon-chain/cache/depositcache"
	blockfeed "github.com/theQRL/zond/beacon-chain/core/feed/block"
	"github.com/theQRL/zond/beacon-chain/core/feed/operation"
	statefeed "github.com/theQRL/zond/beacon-chain/core/feed/state"
	"github.com/theQRL/zond/beacon-chain/db"
	"github.com/theQRL/zond/beacon-chain/execution"
	"github.com/theQRL/zond/beacon-chain/operations/attestations"
	"github.com/theQRL/zond/beacon-chain/operations/slashings"
	"github.com/theQRL/zond/beacon-chain/p2p"
	"github.com/theQRL/zond/beacon-chain/state/stategen"
	"github.com/theQRL/zond/beacon-chain/sync"
	ethpb "github.com/theQRL/zond/protos/zond/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum beacon chain.
type Server struct {
	BeaconDB                    db.ReadOnlyDatabase
	Ctx                         context.Context
	ChainStartFetcher           execution.ChainStartFetcher
	HeadFetcher                 blockchain.HeadFetcher
	CanonicalFetcher            blockchain.CanonicalFetcher
	FinalizationFetcher         blockchain.FinalizationFetcher
	DepositFetcher              depositcache.DepositFetcher
	BlockFetcher                execution.POWBlockFetcher
	GenesisTimeFetcher          blockchain.TimeFetcher
	StateNotifier               statefeed.Notifier
	BlockNotifier               blockfeed.Notifier
	AttestationNotifier         operation.Notifier
	Broadcaster                 p2p.Broadcaster
	AttestationsPool            attestations.Pool
	SlashingsPool               slashings.PoolManager
	ChainStartChan              chan time.Time
	ReceivedAttestationsBuffer  chan *ethpb.Attestation
	CollectedAttestationsBuffer chan []*ethpb.Attestation
	StateGen                    stategen.StateManager
	SyncChecker                 sync.Checker
	ReplayerBuilder             stategen.ReplayerBuilder
	HeadUpdater                 blockchain.HeadUpdater
	OptimisticModeFetcher       blockchain.OptimisticModeFetcher
}
