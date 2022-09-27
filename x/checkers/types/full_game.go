package types

import (
	"errors"
	"fmt"

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
	return err
}
