package node

import (
	"github.com/theQRL/zond/beacon-chain/blockchain"
	"github.com/theQRL/zond/beacon-chain/builder"
	"github.com/theQRL/zond/beacon-chain/execution"
)

// Option for beacon node configuration.
type Option func(bn *BeaconNode) error

// WithBlockchainFlagOptions includes functional options for the blockchain service related to CLI flags.
func WithBlockchainFlagOptions(opts []blockchain.Option) Option {
	return func(bn *BeaconNode) error {
		bn.serviceFlagOpts.blockchainFlagOpts = opts
		return nil
	}
}

// WithExecutionChainOptions includes functional options for the execution chain service related to CLI flags.
func WithExecutionChainOptions(opts []execution.Option) Option {
	return func(bn *BeaconNode) error {
		bn.serviceFlagOpts.executionChainFlagOpts = opts
		return nil
	}
}

// WithBuilderFlagOptions includes functional options for the builder service related to CLI flags.
func WithBuilderFlagOptions(opts []builder.Option) Option {
	return func(bn *BeaconNode) error {
		bn.serviceFlagOpts.builderOpts = opts
		return nil
	}
}
