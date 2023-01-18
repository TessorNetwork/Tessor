package keeper

import (
	"github.com/TessorNetwork/tessor/v4/x/stakeibc/types"
)

var _ types.QueryServer = Keeper{}
