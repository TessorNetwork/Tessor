package keeper

import (
	"github.com/TessorNetwork/tessor/v4/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
