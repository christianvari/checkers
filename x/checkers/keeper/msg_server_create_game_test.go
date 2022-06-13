package keeper_test

import (
	"github.com/christianvari/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestCreateGame() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	createResponse, err := suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   12,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgCreateGameResponse{
		IdValue: "1",
	}, *createResponse)
}

func (suite *IntegrationTestSuite) TestCreateGameDidNotPay() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   12,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}
