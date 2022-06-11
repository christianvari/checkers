package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)

var (
	ErrInvalidCreator   = sdkerrors.Register(ModuleName, 1101, "creator address is invalid: %s")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1102, "red address is invalid: %s")
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1103, "black address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1104, "game cannot be parsed")
)
