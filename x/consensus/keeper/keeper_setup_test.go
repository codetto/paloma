package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	"github.com/palomachain/paloma/x/consensus/types"
	"github.com/palomachain/paloma/x/consensus/types/mocks"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

type mockedServices struct {
	ValsetKeeper *mocks.ValsetKeeper
}

func newConsensusKeeper(t testing.TB) (*Keeper, mockedServices, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	types.RegisterInterfaces(registry)

	ss := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"ConsensusSubSpace",
	)
	IBCKeeper := ibckeeper.NewKeeper(
		appCodec,
		storeKey,
		ss,
		nil,
		nil,
		capabilityKeeper.ScopeToModule("ConsensusIBCKeeper"),
	)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"ConsensusParams",
	)
	ms := mockedServices{
		ValsetKeeper: mocks.NewValsetKeeper(t),
	}
	k := NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		IBCKeeper.ChannelKeeper,
		&IBCKeeper.PortKeeper,
		capabilityKeeper.ScopeToModule("ConsensusScopedKeeper"),
		ms.ValsetKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ms, ctx
}