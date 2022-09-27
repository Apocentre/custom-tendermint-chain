package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/apocentre/checkers/testutil/keeper"
	"github.com/apocentre/checkers/x/checkers"
	"github.com/apocentre/checkers/x/checkers/keeper"
	"github.com/apocentre/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	// Initialize the keeper with the default genesis
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)

	response, err := msgServer.CreateGame(context, &types.MsgCreateGame {
		Creator: alice,
		Black: bob,
		Red: carol,
	})

	require.Nil(t, err)
	require.Equal(t, response.GameIndex, "1")
	require.EqualValues(t, types.MsgCreateGameResponse {GameIndex: "1"}, *response)
}
