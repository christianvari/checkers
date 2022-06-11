package types_test

import (
	"testing"

	"github.com/christianvari/checkers/x/checkers/rules"
	"github.com/christianvari/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *types.StoredGame {
	return &types.StoredGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Index:   "1",
		Game:    rules.New().String(),
		Turn:    "a"}
}

func TestCanGetAddressRed(t *testing.T) {
	carolAddress, err1 := sdk.AccAddressFromBech32(carol)
	redAddress, err2 := GetStoredGame1().GetRedAddress()
	require.Equal(t, carolAddress, redAddress)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	redAddress, err := storedGame.GetRedAddress()
	require.Nil(t, redAddress)
	require.EqualError(t,
		err,
		"red address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
}

func TestCanGetAddressBlack(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	blackAddress, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, bobAddress, blackAddress)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	blackAddress, err := storedGame.GetBlackAddress()
	require.Nil(t, blackAddress)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
}
