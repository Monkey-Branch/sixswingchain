package keeper

import (
	"context"
	"sixswingchain/x/points/types"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	if ctx.TxBytes() != nil {
		if msg.Creator != msg.Owner {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "tx must be signed by owner")
		}
	}

	denom := "swing"
	burnCoins := sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(int64(msg.Amount))))

	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err = k.bankKeeper.SendCoins(ctx, ownerAddr, moduleAcct, burnCoins)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "failed to send coins to module account: %v", err)
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to burn coins: %v", err)
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("burn_tokens",
			sdk.NewAttribute("owner", msg.Owner),
			sdk.NewAttribute("amount", Int32ToString(msg.Amount)),
		),
	)

	return &types.MsgBurnResponse{}, nil
}

func Int32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10) // base 10
}
