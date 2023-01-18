package keeper

import (
	"github.com/TessorNetwork/Tessor/x/records/types"
)

var _ types.QueryServer = Keeper{}
