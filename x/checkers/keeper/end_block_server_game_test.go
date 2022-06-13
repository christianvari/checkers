package keeper_test

import (
	"time"

	"github.com/christianvari/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestForfeitUnplayed() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[createEventCount:])

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitOlderUnplayed() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		IdValue:  3,
		FifoHead: "2",
		FifoTail: "2",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[2*createEventCount:])
}

func (suite *IntegrationTestSuite) TestForfeit2OldestUnplayedIn1Call() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   13,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	game2, found := keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().True(found)
	game2.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game2)
	keeper.ForfeitExpiredGames(goCtx)

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)
	_, found = keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		IdValue:  4,
		FifoHead: "3",
		FifoTail: "3",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	forfeitAttributes := forfeitEvent.Attributes[3*createEventCount:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitAttributes[:4])
	forfeitAttributes = forfeitAttributes[4:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "2"},
		{Key: "Winner", Value: "*"},
	}, forfeitAttributes)
}
