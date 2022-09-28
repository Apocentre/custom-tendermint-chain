package types

import (
	"errors"
	"fmt"
	"time"

	"github.com/apocentre/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// / Your stored game's black field is only string, but they represent sdk.AccAddress
func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

// / Your stored game's red field is only string, but they represent sdk.AccAddress
func (storedGame StoredGame) GetRedAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidRed.Error(), storedGame.Red)
}

// / Parse the game so that it can be played. The Turn has to be set by hand:
func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)

	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}

	return board, nil
}

// / checks a game's validity:
func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}

	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}
	_, err = storedGame.GetDeadlineAsTime()

	return err
}

func (storedGame *StoredGame) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, errDeadline := time.Parse(DeadlineLayout, storedGame.Deadline)
	return deadline, sdkerrors.Wrapf(errDeadline, ErrInvalidDeadline.Error(), storedGame.Deadline)
}

// encapsulates how the next deadline is calculated
func GetNextDeadline(ctx sdk.Context) time.Time {
	return ctx.BlockTime().Add(MaxTurnDuration)
}

func FormatDeadline(deadline time.Time) string {
	return deadline.UTC().Format(DeadlineLayout)
}

func (storedGame StoredGame) GetPlayerAddress(color string) (address sdk.AccAddress, found bool, err error) {
	black, err := storedGame.GetBlackAddress()
	if err != nil {
		return nil, false, err
	}
	red, err := storedGame.GetRedAddress()
	if err != nil {
		return nil, false, err
	}
	address, found = map[string]sdk.AccAddress{
		rules.PieceStrings[rules.BLACK_PLAYER]: black,
		rules.PieceStrings[rules.RED_PLAYER]:   red,
	}[color]

	return address, found, nil
}

func (storedGame StoredGame) GetWinnerAddress() (address sdk.AccAddress, found bool, err error) {
	return storedGame.GetPlayerAddress(storedGame.Winner)
}

func (storedGame *StoredGame) GetWagerCoin() (wager sdk.Coin) {
	return sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(storedGame.Wager)))
}
