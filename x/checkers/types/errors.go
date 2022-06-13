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
	ErrInvalidCreator          = sdkerrors.Register(ModuleName, 1101, "creator address is invalid: %s")
	ErrInvalidRed              = sdkerrors.Register(ModuleName, 1102, "red address is invalid: %s")
	ErrInvalidBlack            = sdkerrors.Register(ModuleName, 1103, "black address is invalid: %s")
	ErrGameNotParseable        = sdkerrors.Register(ModuleName, 1104, "game cannot be parsed")
	ErrGameNotFound            = sdkerrors.Register(ModuleName, 1105, "game by id not found: %s")
	ErrCreatorNotPlayer        = sdkerrors.Register(ModuleName, 1106, "message creator is not a player: %s")
	ErrNotPlayerTurn           = sdkerrors.Register(ModuleName, 1107, "player tried to play out of turn: %s")
	ErrWrongMove               = sdkerrors.Register(ModuleName, 1108, "wrong move")
	ErrRedAlreadyPlayed        = sdkerrors.Register(ModuleName, 1109, "red player has already played")
	ErrBlackAlreadyPlayed      = sdkerrors.Register(ModuleName, 1110, "black player has already played")
	ErrInvalidDeadline         = sdkerrors.Register(ModuleName, 1111, "deadline cannot be parsed: %s")
	ErrGameFinished            = sdkerrors.Register(ModuleName, 1112, "game is already finished")
	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1113, "cannot find winner by color %s")
	ErrBlackCannotPay          = sdkerrors.Register(ModuleName, 1114, "black cannot pay the wager")
	ErrNothingToPay            = sdkerrors.Register(ModuleName, 1115, "there is nothing to pay, should not have been called")
	ErrCannotRefundWager       = sdkerrors.Register(ModuleName, 1116, "cannot refund wager to: %s")
	ErrCannotPayWinnings       = sdkerrors.Register(ModuleName, 1117, "cannot pay winnings to winner")
	ErrNotInRefundState        = sdkerrors.Register(ModuleName, 1118, "game is not in a state to refund, move count: %d")
	ErrRedCannotPay            = sdkerrors.Register(ModuleName, 1119, "red cannot pay the wager")
)
