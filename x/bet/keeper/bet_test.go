package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sge-network/sge/app/params"
	"github.com/sge-network/sge/testutil/nullify"
	"github.com/sge-network/sge/x/bet/keeper"
	"github.com/sge-network/sge/x/bet/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Bet {
	items := make([]types.Bet, n)
	for i := range items {
		items[i].UID = strconv.Itoa(i)
		items[i].OddsValue = sdk.NewDec(10)
		items[i].Amount = sdk.NewInt(10)
		items[i].BetFee = sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(1))

		keeper.SetBet(ctx, items[i])
	}
	return items
}

func TestBetGet(t *testing.T) {
	k, ctx := setupKeeper(t)
	items := createNBet(k, ctx, 10)

	rst, found := k.GetBet(ctx,
		"NotExistUid",
	)
	var expectedResp types.Bet
	require.False(t, found)
	require.Equal(t,
		nullify.Fill(&expectedResp),
		nullify.Fill(&rst),
	)

	for _, item := range items {
		rst, found := k.GetBet(ctx,
			item.UID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestBetGetAll(t *testing.T) {
	k, ctx := setupKeeper(t)
	items := createNBet(k, ctx, 10)

	bets, err := k.GetBetAll(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(bets),
	)
}
