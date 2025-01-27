package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sge-network/sge/app/params"
	"github.com/sge-network/sge/testutil/network"
	"github.com/sge-network/sge/testutil/nullify"
	"github.com/sge-network/sge/x/bet/client/cli"
	"github.com/sge-network/sge/x/bet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithBetObjects(t *testing.T, n int) (*network.Network, []types.Bet) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		bet := types.Bet{
			UID:       strconv.Itoa(i),
			OddsValue: sdk.NewDec(10),
			Amount:    sdk.NewInt(10),
			BetFee:    sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(1)),
		}
		nullify.Fill(&bet)
		state.BetList = append(state.BetList, bet)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.BetList
}

func TestShowBet(t *testing.T) {
	net, objs := networkWithBetObjects(t, 5)

	t.Run("ShowBet", func(t *testing.T) {
		ctx := net.Validators[0].ClientCtx
		common := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		for _, tc := range []struct {
			desc string
			uid  string

			args []string
			err  error
			obj  types.Bet
		}{
			{
				desc: "found",
				uid:  objs[0].UID,

				args: common,
				obj:  objs[0],
			},
			{
				desc: "not found",
				uid:  strconv.Itoa(100000),

				args: common,
				err:  status.Error(codes.NotFound, "not found"),
			},
		} {
			tc := tc
			t.Run(tc.desc, func(t *testing.T) {
				args := []string{
					tc.uid,
				}
				args = append(args, tc.args...)
				out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowBet(), args)
				if tc.err != nil {
					stat, ok := status.FromError(tc.err)
					require.True(t, ok)
					require.ErrorIs(t, stat.Err(), tc.err)
				} else {
					require.NoError(t, err)
					var resp types.QueryBetResponse
					require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
					require.NotNil(t, resp.Bet)
					require.Equal(t,
						nullify.Fill(&tc.obj),
						nullify.Fill(&resp.Bet),
					)
				}
			})
		}
	})

	t.Run("ListBet", func(t *testing.T) {
		ctx := net.Validators[0].ClientCtx
		request := func(next []byte, offset, limit uint64, total bool) []string {
			args := []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			}
			if next == nil {
				args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
			} else {
				args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
			}
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
			if total {
				args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
			}
			return args
		}
		t.Run("ByOffset", func(t *testing.T) {
			step := 2
			for i := 0; i < len(objs); i += step {
				args := request(nil, uint64(i), uint64(step), false)
				out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListBet(), args)
				require.NoError(t, err)
				var resp types.QueryListBetAllResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.LessOrEqual(t, len(resp.Bet), step)
				require.Subset(t,
					nullify.Fill(objs),
					nullify.Fill(resp.Bet),
				)
			}
		})
		t.Run("ByKey", func(t *testing.T) {
			step := 2
			var next []byte
			for i := 0; i < len(objs); i += step {
				args := request(next, 0, uint64(step), false)
				out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListBet(), args)
				require.NoError(t, err)
				var resp types.QueryListBetAllResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.LessOrEqual(t, len(resp.Bet), step)
				require.Subset(t,
					nullify.Fill(objs),
					nullify.Fill(resp.Bet),
				)
				next = resp.Pagination.NextKey
			}
		})
		t.Run("Total", func(t *testing.T) {
			args := request(nil, 0, uint64(len(objs)), true)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListBet(), args)
			require.NoError(t, err)
			var resp types.QueryListBetAllResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, err)
			require.Equal(t, len(objs), int(resp.Pagination.Total))
			require.ElementsMatch(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Bet),
			)
		})
	})
}
