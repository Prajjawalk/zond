// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package ethapi implements the general Zond API functions.
package zondapi

import (
	"context"
	"time"

	"github.com/theQRL/zond/common"
	"github.com/theQRL/zond/core"
	"github.com/theQRL/zond/core/state"
	"github.com/theQRL/zond/core/types"
	"github.com/theQRL/zond/core/vm"
	"github.com/theQRL/zond/metadata"
	"github.com/theQRL/zond/params"
	"github.com/theQRL/zond/protos"
	"github.com/theQRL/zond/rpc"
	"github.com/theQRL/zond/transactions"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Ethereum API
	//SyncProgress() ethereum.SyncProgress

	//SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	//FeeHistory(ctx context.Context, blockCount int, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error)
	//ChainDb() ethdb.Database
	//AccountManager() *accounts.Manager
	//ExtRPCEnabled() bool
	RPCGasCap() uint64            // global gas cap for eth_call over rpc: DoS protection
	RPCEVMTimeout() time.Duration // global timeout for eth_call over rpc: DoS protection
	//RPCTxFeeCap() float64         // global tx fee cap for all transaction related APIs
	//UnprotectedAllowed() bool     // allows only for EIP155 transactions.

	// Blockchain API
	//SetHead(number uint64)
	GetValidators(ctx context.Context) (*metadata.EpochMetaData, error)
	HeaderByNumberV1(ctx context.Context, number rpc.BlockNumber) (*protos.BlockHeader, error)
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	HeaderByHashV1(ctx context.Context, hash common.Hash) (*protos.BlockHeader, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error)
	//CurrentHeader() *types.Header
	//CurrentBlock() *types.Block
	BlockByNumberV1(ctx context.Context, number rpc.BlockNumber) (*protos.Block, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error)
	BlockByHashV1(ctx context.Context, hash common.Hash) (*protos.Block, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumberOrHashV1(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*protos.Block, error)
	BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	StateAndHeaderByNumberOrHashV1(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *protos.BlockHeader, error)
	StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error)
	PendingBlockAndReceipts() (*types.Block, types.Receipts)
	GetReceiptsV1(ctx context.Context, hash common.Hash, isProtocolTransaction bool) (types.Receipts, error)
	GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error)
	//GetTd(ctx context.Context, hash common.Hash) *big.Int
	GetEVMV1(ctx context.Context, msg core.Message, state *state.StateDB, header *protos.BlockHeader, vmConfig *vm.Config) (*vm.EVM, func() error, error)
	GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config) (*vm.EVM, func() error, error)
	//SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	//SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	//SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription

	// Transaction pool API
	SendTxV1(ctx context.Context, signedTx transactions.TransactionInterface) error
	GetTransactionV1(ctx context.Context, txHash common.Hash) (*protos.Transaction, common.Hash, uint64, uint64, error)
	GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonceV1(ctx context.Context, addr common.Address) (uint64, error)
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	//Stats() (pending int, queued int)
	//TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	//TxPoolContentFrom(addr common.Address) (types.Transactions, types.Transactions)
	//SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription

	// Filter API
	//BloomStatus() (uint64, uint64)
	GetLogsV1(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error)
	GetLogs(ctx context.Context, hash common.Hash, number uint64) ([][]*types.Log, error)
	//ServiceFilter(ctx context.Context, session *bloombits.MatcherSession)
	//SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription
	//SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription
	//SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription

	ChainConfig() *params.ChainConfig
	//Engine() consensus.Engine
}

func GetAPIs(apiBackend Backend) []rpc.API {
	//nonceLock := new(AddrLocker)
	return []rpc.API{
		{
			Namespace: "zond",
			Version:   "0.1",
			Service:   NewBlockChainAPI(apiBackend),
		},
		{
			Namespace: "zond",
			Version:   "0.1",
			Service:   NewTransactionAPI(apiBackend, nil),
		},
	}
}
