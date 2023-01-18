package keeper

import (
	"github.com/TessorNetwork/Tessor/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
