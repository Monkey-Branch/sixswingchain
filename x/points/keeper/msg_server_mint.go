package keeper

import (
	"context"
	"fmt"

	"sixswingchain/x/points/types"

	"cosmossdk.io/math"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verifica se o owner está definido e se o remetente é o owner
	currentOwner, found := k.GetOwner(ctx)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "owner not set")
	}
	if currentOwner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can mint tokens")
	}

	// Converte o Amount em Coins com denom "swing"
	swingCoins := sdk.NewCoins(sdk.NewCoin("swing", math.NewInt(int64(msg.Amount))))

	// Mintar as moedas no module account
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, swingCoins); err != nil {
		return nil, err
	}

	// Enviar do module account para o recipient
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	recipientAcc, err := sdk.AccAddressFromBech32(msg.Wallet)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address: %v", err)
	}
	if err := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAcc, swingCoins); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("mint_swing_tokens",
			sdk.NewAttribute("owner", msg.Creator),
			sdk.NewAttribute("recipient", msg.Wallet),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
		),
	)

	return &types.MsgMintResponse{}, nil
}