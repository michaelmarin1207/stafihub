package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/x/ledger/keeper"
	//"github.com/stafihub/stafihub/x/ledger/types"
)

func setupSettings(t testing.TB) {
	k, ctx := keepertest.LedgerKeeper(t)
	s, wctx := keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
	msg := addNewPool(s, wctx)
	t.Log(msg)
}
