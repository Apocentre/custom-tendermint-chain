package keeper_test

import (
	"testing"

	"github.com/apocentre/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)


func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)

	response, err := msgServer.CreateGame(context, &types.MsgCreateGame {
		Creator: alice,
		Black: bob,
		Red: carol,
	})

	require.Nil(t, err)
	require.Equal(t, response.GameIndex, "1")
	require.EqualValues(t, types.MsgCreateGameResponse {GameIndex: "1"}, *response)
}
