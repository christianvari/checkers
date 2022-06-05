package keeper

import (
	"github.com/christianvari/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
