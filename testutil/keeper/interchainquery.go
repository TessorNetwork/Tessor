package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	tessorapp "github.com/TessorNetwork/tessor/v4/app"
	"github.com/TessorNetwork/tessor/v4/x/interchainquery/keeper"
)

func InterchainqueryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := tessorapp.InitTessorTestApp(true)
	icqKeeper := app.InterchainqueryKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "tessor-1", Time: time.Now().UTC()})

	return &icqKeeper, ctx
}
