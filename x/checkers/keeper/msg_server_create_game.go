package keeper

import (
	"context"
	"strconv"

	"github.com/apocentre/checkers/x/checkers/rules"
	"github.com/apocentre/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the next game id from the store via Keeper.GetSystemInfo
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	gameId := strconv.FormatUint(systemInfo.NextId, 10)

	// Create the object to be stored:
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index: gameId,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
		MoveCount: 0,
		BeforeIndex: types.NoFifoIndex,
    AfterIndex:  types.NoFifoIndex,
		Deadline: types.FormatDeadline(types.GetNextDeadline(ctx)),
		Winner:    rules.PieceStrings[rules.NO_PLAYER],
		Wager: msg.Wager,
	}

	// Confirm that the values in the object are correct by checking the validity
	err := storedGame.Validate()

	if err != nil {
		return nil, err
	}

	// Send the new game to the tail because it is freshly created
	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	// Save the StoredGame object using the Keeper.SetStoredGame
	k.Keeper.SetStoredGame(ctx, storedGame)

	// Increase the next game id
	systemInfo.NextId += 1
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	// Now you must implement this correspondingly in the GUI, or include a server to listen for such events.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, gameId),
			sdk.NewAttribute(types.GameCreatedEventBlack, msg.Black),
			sdk.NewAttribute(types.GameCreatedEventRed, msg.Red),
			sdk.NewAttribute(types.GameCreatedEventWager, strconv.FormatUint(msg.Wager, 10)),
		),
	)

	// Return the newly created ID
	return &types.MsgCreateGameResponse{GameIndex: gameId}, nil
}
