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

	// 1. Validar e converter endereços
	ownerAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	// 2. (Opcional) Verificar se quem assina a Tx é o msg.Owner
	//    (Ignite normalmente gera esse boilerplate em `MsgBurn.ValidateBasic()`,
	//    mas pode incluir verificação extra aqui)
	if ctx.TxBytes() != nil {
		// Em cosmos sdk 0.45 / 0.46, você pode checar a signature do Tx
		// ou utilizar "msg.GetSigners()" comparando com msg.Owner.
		// Exemplo simples:

		if msg.Creator != msg.Owner {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "tx must be signed by owner")
		}
	}

	// 3. Converter msg.Amount para sdk.Coins (supondo 1 denom, ex. "swing")
	//    Se você quiser permitir denom dinâmica, use outro campo no Msg
	//    ou se for fixo "points"/"swing", faça algo como:
	denom := "swing" // ou "points"
	burnCoins := sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(int64(msg.Amount))))

	// 4. Enviar as moedas do usuário => para a module account do seu módulo
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err = k.bankKeeper.SendCoins(ctx, ownerAddr, moduleAcct, burnCoins)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "failed to send coins to module account: %v", err)
	}

	// 5. Queimar as moedas que agora estão no module account
	//    Importante: garanta que seu modulo tenha permissão "Burner" no app.go/config
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to burn coins: %v", err)
	}

	// 6. (Opcional) emit event, atualizar stores, etc.
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
