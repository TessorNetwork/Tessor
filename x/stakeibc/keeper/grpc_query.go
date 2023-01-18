package keeper

import (
	"github.com/TessorNetwork/Tessor/x/stakeibc/types"
)

var _ types.QueryServer = Keeper{}
