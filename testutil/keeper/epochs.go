package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	tessorapp "github.com/TessorNetwork/Tessor/app"
	"github.com/TessorNetwork/Tessor/x/epochs/keeper"
)

func EpochsKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := tessorapp.InitTessorTestApp(true)
	epochsKeeper := app.EpochsKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "tessor-1", Time: time.Now().UTC()})

	return &epochsKeeper, ctx
}
