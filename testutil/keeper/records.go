package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	tessorapp "github.com/TessorNetwork/Tessor/app"
	"github.com/TessorNetwork/Tessor/x/records/keeper"
)

func RecordsKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := tessorapp.InitTessorTestApp(true)
	recordKeeper := app.RecordsKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "tessor-1", Time: time.Now().UTC()})

	return &recordKeeper, ctx
}
