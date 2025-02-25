package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const OwnerKey = "owner"

func (k Keeper) GetOwner(ctx sdk.Context) (string, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Use um prefixo consistente, por exemplo, "points:" ou outro
	store := prefix.NewStore(storeAdapter, []byte("points:"))
	bz := store.Get([]byte(OwnerKey))
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetNewOwner(ctx sdk.Context, owner string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte("points:"))
	store.Set([]byte(OwnerKey), []byte(owner))
}
