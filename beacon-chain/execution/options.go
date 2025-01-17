package execution

import (
	"github.com/theQRL/zond/beacon-chain/cache/depositcache"
	statefeed "github.com/theQRL/zond/beacon-chain/core/feed/state"
	"github.com/theQRL/zond/beacon-chain/db"
	"github.com/theQRL/zond/beacon-chain/state"
	"github.com/theQRL/zond/beacon-chain/state/stategen"
	"github.com/theQRL/zond/common"
)

type Option func(s *Service) error

// WithHttpEndpoint parse http endpoint for the powchain service to use.
func WithHttpEndpoint(endpointString string) Option {
	return func(s *Service) error {
		s.cfg.currHttpEndpoint = HttpEndpoint(endpointString)
		return nil
	}
}

// WithHttpEndpointAndJWTSecret for authenticating the execution node JSON-RPC endpoint.
func WithHttpEndpointAndJWTSecret(endpointString string, secret []byte) Option {
	return func(s *Service) error {
		if len(secret) == 0 {
			return nil
		}
		// Overwrite authorization type for all endpoints to be of a bearer type.
		hEndpoint := HttpEndpoint(endpointString)
		// hEndpoint.Auth.Method = authorization.Bearer
		hEndpoint.Auth.Value = string(secret)

		s.cfg.currHttpEndpoint = hEndpoint
		return nil
	}
}

// WithHeaders adds headers to the execution node JSON-RPC requests.
func WithHeaders(headers []string) Option {
	return func(s *Service) error {
		s.cfg.headers = headers
		return nil
	}
}

// WithDepositContractAddress for the deposit contract.
func WithDepositContractAddress(addr common.Address) Option {
	return func(s *Service) error {
		s.cfg.depositContractAddr = addr
		return nil
	}
}

// WithDatabase for the beacon chain database.
func WithDatabase(database db.HeadAccessDatabase) Option {
	return func(s *Service) error {
		s.cfg.beaconDB = database
		return nil
	}
}

// WithDepositCache for caching deposits.
func WithDepositCache(cache *depositcache.DepositCache) Option {
	return func(s *Service) error {
		s.cfg.depositCache = cache
		return nil
	}
}

// WithStateNotifier for subscribing to state changes.
func WithStateNotifier(notifier statefeed.Notifier) Option {
	return func(s *Service) error {
		s.cfg.stateNotifier = notifier
		return nil
	}
}

// WithStateGen to regenerate beacon states from checkpoints.
func WithStateGen(gen *stategen.State) Option {
	return func(s *Service) error {
		s.cfg.stateGen = gen
		return nil
	}
}

// WithEth1HeaderRequestLimit to set the upper limit of eth1 header requests.
func WithEth1HeaderRequestLimit(limit uint64) Option {
	return func(s *Service) error {
		s.cfg.eth1HeaderReqLimit = limit
		return nil
	}
}

// WithBeaconNodeStatsUpdater to set the beacon node stats updater.
func WithBeaconNodeStatsUpdater(updater BeaconNodeStatsUpdater) Option {
	return func(s *Service) error {
		s.cfg.beaconNodeStatsUpdater = updater
		return nil
	}
}

// WithFinalizedStateAtStartup to set the beacon node's finalized state at startup.
func WithFinalizedStateAtStartup(st state.BeaconState) Option {
	return func(s *Service) error {
		s.cfg.finalizedStateAtStartup = st
		return nil
	}
}
