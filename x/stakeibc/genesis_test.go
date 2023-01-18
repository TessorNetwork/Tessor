package stakeibc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/TessorNetwork/Tessor/testutil/keeper"
	"github.com/TessorNetwork/Tessor/testutil/nullify"
	"github.com/TessorNetwork/Tessor/x/stakeibc"
	"github.com/TessorNetwork/Tessor/x/stakeibc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		IcaAccount: &types.ICAAccount{
			Address: "78",
		},
		EpochTrackerList: []types.EpochTracker{
			{EpochIdentifier: "tessor_epoch"},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StakeibcKeeper(t)
	stakeibc.InitGenesis(ctx, *k, genesisState)
	got := stakeibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)
	require.Equal(t, genesisState.IcaAccount, got.IcaAccount)
	require.Equal(t, genesisState.EpochTrackerList, got.EpochTrackerList)
	require.Equal(t, genesisState.Params, got.Params)
	// this line is used by starport scaffolding # genesis/test/assert
}
