package types_test

import (
	"testing"

	"github.com/apocentre/checkers/x/checkers/rules"
	"github.com/apocentre/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
			Black: alice,
			Red:   bob,
			Index: "1",
			Board: rules.New().String(),
			Turn:  "b",
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	black, err2 := GetStoredGame1().GetBlackAddress()

	require.Equal(t, aliceAddress, black)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestCanGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4" // Bad last digit
	
	black, err := storedGame.GetBlackAddress()

	require.Nil(t, black)
	require.EqualError(
		t, err,
		"black address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)",
	)
	require.EqualError(t, storedGame.Validate(), err.Error())
}
