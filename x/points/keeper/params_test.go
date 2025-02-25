package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "sixswingchain/testutil/keeper"
	"sixswingchain/x/points/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PointsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
