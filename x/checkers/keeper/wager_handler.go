package keeper

import (
	"fmt"

	"github.com/apocentre/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k *Keeper) CollectWager(ctx sdk.Context, storedGame *types.StoredGame) error {
	// Collecting wagers happens on a player's first move.
	// Therefore, differentiate between players:

	if storedGame.MoveCount == 0 {
		black, err := storedGame.GetBlackAddress()
		if err != nil {
			panic(err.Error())
		}

		// try to transfer into the escrow:
		// The reason why we can trnansfer from the black user account to the module by simply providing the module name
		// is because the bank module creates an address for your module's escrow account. When you have the full app, you can access it with:
		// checkersModuleAddress := app.AccountKeeper.GetModuleAddress(types.ModuleName)
		k.bank.SendCoinsFromAccountToModule(ctx, black, types.ModuleName, sdk.NewCoins(storedGame.GetWagerCoin()))
		if err != nil {
			return sdkerrors.Wrapf(err, types.ErrBlackCannotPay.Error())
		}
	} else if storedGame.MoveCount == 1 {
		// Red plays second
		red, err := storedGame.GetRedAddress()
		if err != nil {
			panic(err.Error())
		}

		err = k.bank.SendCoinsFromAccountToModule(ctx, red, types.ModuleName, sdk.NewCoins(storedGame.GetWagerCoin()))
		if err != nil {
			return sdkerrors.Wrapf(err, types.ErrRedCannotPay.Error())
		}
	}

	return nil
}

func (k *Keeper) MustPayWinnings(ctx sdk.Context, storedGame *types.StoredGame) {
	winnerAddress, found, err := storedGame.GetWinnerAddress()
	if err != nil {
		panic(err.Error())
	}
	if !found {
		panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Winner))
	}

	// calculate the winnings to pay
	winnings := storedGame.GetWagerCoin()
	if storedGame.MoveCount == 0 {
		panic(types.ErrNothingToPay.Error())
	} else if 1 < storedGame.MoveCount {
		// You double the wager only if the red player has also played and therefore both players have paid their wagers
		winnings = winnings.Add(winnings)
	}

	// pay the winner
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winnerAddress, sdk.NewCoins(winnings))
	if err != nil {
		panic(fmt.Sprintf(types.ErrCannotPayWinnings.Error(), err.Error()))
	}
}

// refunding wagers takes place when the game has partially started, i.e. only one party has paid,
// or when the game ends in a draw.
func (k *Keeper) MustRefundWager(ctx sdk.Context, storedGame *types.StoredGame) {
	if storedGame.MoveCount == 1 {
		// Refund
		black, err := storedGame.GetBlackAddress()
		if err != nil {
			panic(err.Error())
		}
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, black, sdk.NewCoins(storedGame.GetWagerCoin()))
		if err != nil {
			panic(fmt.Sprintf(types.ErrCannotRefundWager.Error(), err.Error()))
		}
	} else if storedGame.MoveCount == 0 {
		// Do nothing
	} else {
		// TODO Implement a draw mechanism.
		panic(fmt.Sprintf(types.ErrNotInRefundState.Error(), storedGame.MoveCount))
	}
}

