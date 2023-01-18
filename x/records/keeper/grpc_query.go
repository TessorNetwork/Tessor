package keeper

import (
	"github.com/TessorNetwork/tessor/v4/x/records/types"
)

var _ types.QueryServer = Keeper{}
