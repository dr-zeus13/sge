package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sge-network/sge/consts"
	"github.com/sge-network/sge/testutil/nullify"
	"github.com/sge-network/sge/x/bet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBetQuerySingle(t *testing.T) {
	k, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBet(k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryBetRequest
		response *types.QueryBetResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryBetRequest{
				Uid: msgs[0].UID,
			},
			response: &types.QueryBetResponse{Bet: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryBetRequest{
				Uid: msgs[1].UID,
			},
			response: &types.QueryBetResponse{Bet: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryBetRequest{
				Uid: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.Bet(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestBetQueryPaginated(t *testing.T) {
	k, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBet(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryListBetAllRequest {
		return &types.QueryListBetAllRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := k.ListAllBet(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Bet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Bet),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.ListAllBet(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Bet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Bet),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.ListAllBet(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Bet),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.ListAllBet(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest))
	})
}
