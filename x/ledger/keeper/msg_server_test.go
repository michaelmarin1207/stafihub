package keeper_test

import (
	"context"
	"github.com/stafiprotocol/stafihub/testutil/sample"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/ledger/keeper"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LedgerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func addNewPool(s types.MsgServer, ctx context.Context) *types.MsgAddNewPool {
	msg := types.NewMsgAddNewPool(sample.TestAdminAcc, sample.TestDenom, sample.AccAddress())
	_, err := s.AddNewPool(ctx, msg)
	if err != nil {
		panic(err)
	}
	return msg
}

func TestMsgServer_AddNewPool(t *testing.T) {
	s, wctx := setupMsgServer(t)
	msg := addNewPool(s, wctx)
	t.Log(msg)
}
