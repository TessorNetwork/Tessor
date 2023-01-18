package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	tessorapp "github.com/TessorNetwork/Tessor/app"
	"github.com/TessorNetwork/Tessor/x/stakeibc/keeper"
)

func StakeibcKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := tessorapp.InitTessorTestApp(true)
	stakeibcKeeper := app.StakeibcKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "tessor-1", Time: time.Now().UTC()})

	return &stakeibcKeeper, ctx
}
