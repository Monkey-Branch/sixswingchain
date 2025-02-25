package keeper

import (
	"context"

	"sixswingchain/x/points/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetOwner(goCtx context.Context, msg *types.MsgSetOwner) (*types.MsgSetOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Obtém o owner atual do estado
	currentOwner, found := k.GetOwner(ctx)
	if !found {
		// Se não houver owner, define o remetente como owner
		k.SetNewOwner(ctx, msg.Creator)
	} else {
		// Se já houver owner, somente o owner atual pode atualizar
		if currentOwner != msg.Creator {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the current owner can change the owner")
		}
		k.SetNewOwner(ctx, msg.NewOwner)
	}

	return &types.MsgSetOwnerResponse{}, nil
}
