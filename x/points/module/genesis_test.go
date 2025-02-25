package points_test

import (
	"testing"

	keepertest "sixswingchain/testutil/keeper"
	"sixswingchain/testutil/nullify"
	points "sixswingchain/x/points/module"
	"sixswingchain/x/points/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PointsKeeper(t)
	points.InitGenesis(ctx, k, genesisState)
	got := points.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
