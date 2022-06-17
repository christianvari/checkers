package keeper

import (
	"context"
	"strconv"
	"strings"

	"github.com/christianvari/checkers/x/checkers/rules"
	"github.com/christianvari/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.GetStoredGame(ctx, msg.IdValue)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found")
	}

	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, types.ErrGameFinished
	}

	err := k.Keeper.CollectWager(ctx, &storedGame)
	if err != nil {
		return nil, err
	}

	isRed := strings.Compare(storedGame.Red, msg.Creator) == 0
	isBlack := strings.Compare(storedGame.Black, msg.Creator) == 0

	var player rules.Player

	if !isRed && !isBlack {
		return nil, types.ErrCreatorNotPlayer
	} else if isRed && isBlack {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isRed {
		player = rules.RED_PLAYER
	} else {
		player = rules.BLACK_PLAYER
	}

	game, err := storedGame.ParseGame()
	if err != nil {
		panic(err.Error())
	}

	if !game.TurnIs(player) {
		return nil, types.ErrNotPlayerTurn
	}

	captured, moveErr := game.Move(
		rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
		},
		rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
		},
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrWrongMove, moveErr.Error())
	}

	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}

	storedGame.MoveCount++
	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx))
	storedGame.Winner = rules.PieceStrings[game.Winner()]
	storedGame.Game = game.String()
	storedGame.Turn = rules.PieceStrings[game.Turn]

	if storedGame.Winner == rules.PieceStrings[rules.NO_PLAYER] {
		k.Keeper.SendToFifoTail(ctx, &storedGame, &nextGame)
	} else {
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &nextGame)
		k.Keeper.MustPayWinnings(ctx, &storedGame)
	}

	k.Keeper.SetStoredGame(ctx, storedGame)
	k.Keeper.SetNextGame(ctx, nextGame)

	ctx.GasMeter().ConsumeGas(types.PlayMoveGas, "Play a move")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "checkers"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.PlayMoveEventKey),
			sdk.NewAttribute(types.PlayMoveEventCreator, msg.Creator),
			sdk.NewAttribute(types.PlayMoveEventIdValue, msg.IdValue),
			sdk.NewAttribute(types.PlayMoveEventCapturedX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.PlayMoveEventCapturedY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.PlayMoveEventWinner, rules.PieceStrings[game.Winner()]),
		),
	)

	return &types.MsgPlayMoveResponse{IdValue: msg.IdValue, CapturedX: int64(captured.X), CapturedY: int64(captured.Y), Winner: rules.PieceStrings[game.Winner()]}, nil
}
