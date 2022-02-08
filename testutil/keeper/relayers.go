package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/app"
	"github.com/stafiprotocol/stafihub/x/relayers/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	sudokeeper "github.com/stafiprotocol/stafihub/x/sudo/keeper"
	sudotypes "github.com/stafiprotocol/stafihub/x/sudo/types"
)

func RelayersKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	sudoStoreKey := sdk.NewKVStoreKey(sudotypes.StoreKey)
	sudoMemStoreKey := storetypes.NewMemoryStoreKey(sudotypes.MemStoreKey)

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(bankStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(sudoStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(sudoMemStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	encCfg := app.MakeTestEncodingConfig()
	paramsKeeper := ParamsKeeper(&encCfg)
	accountKeeper := AccountKeeper(&encCfg, &paramsKeeper)
	bankKeeper := BankKeeper(&encCfg, &paramsKeeper, &accountKeeper)
	sudoKeeper := SimpleSudoKeeper(sudoStoreKey, sudoMemStoreKey, cdc, bankKeeper)

	k := SimpleRelayersKeeper(storeKey, memStoreKey, cdc, bankKeeper, sudoKeeper)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}

func SimpleRelayersKeeper(storeKey *sdk.KVStoreKey, memStoreKey *sdk.MemoryStoreKey, cdc *codec.ProtoCodec, bankKeeper bankkeeper.Keeper, sudoKeeper *sudokeeper.Keeper) *keeper.Keeper {
	return keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		sudoKeeper,
		bankKeeper,
	)
}
