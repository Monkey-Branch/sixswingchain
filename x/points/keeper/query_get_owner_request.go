package keeper

import (
	"context"

	"sixswingchain/x/points/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetOwnerRequest(goCtx context.Context, req *types.QueryGetOwnerRequestRequest) (*types.QueryGetOwnerRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, found := k.GetOwner(ctx)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "owner not set")
	}

	return &types.QueryGetOwnerRequestResponse{Owner: owner}, nil
}
