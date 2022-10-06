package milestone_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/heimdall/app"
	"github.com/maticnetwork/heimdall/milestone"
	hmTypes "github.com/maticnetwork/heimdall/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	app *app.HeimdallApp
	ctx sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app, suite.ctx, _ = createTestApp(false)
}

func TestKeeperTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestAddCheckpoint() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.MilestoneKeeper

	startBlock := uint64(0)
	endBlock := uint64(63)
	rootHash := hmTypes.HexToHeimdallHash("123")
	proposerAddress := hmTypes.HexToHeimdallAddress("123")
	timestamp := uint64(time.Now().Unix())
	borChainId := "1234"
	milestoneID := "0000"

	milestone := hmTypes.CreateMilestone(
		startBlock,
		endBlock,
		rootHash,
		proposerAddress,
		borChainId,
		milestoneID,
		timestamp,
	)
	err := keeper.AddMilestone(ctx, milestone)
	require.NoError(t, err)

	result, err := keeper.GetLastMilestone(ctx)
	require.NoError(t, err)
	require.Equal(t, startBlock, result.StartBlock)
	require.Equal(t, endBlock, result.EndBlock)
	require.Equal(t, rootHash, result.RootHash)
	require.Equal(t, borChainId, result.BorChainID)
	require.Equal(t, proposerAddress, result.Proposer)
	require.Equal(t, timestamp, result.TimeStamp)
}

func (suite *KeeperTestSuite) TestGetCount() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.MilestoneKeeper

	result := keeper.GetCount(ctx)
	require.Equal(t, uint64(0), result)

	startBlock := uint64(0)
	endBlock := uint64(63)
	rootHash := hmTypes.HexToHeimdallHash("123")
	proposerAddress := hmTypes.HexToHeimdallAddress("123")
	timestamp := uint64(time.Now().Unix())
	borChainId := "1234"
	milestoneID := "0000"

	milestone := hmTypes.CreateMilestone(
		startBlock,
		endBlock,
		rootHash,
		proposerAddress,
		borChainId,
		milestoneID,
		timestamp,
	)
	err := keeper.AddMilestone(ctx, milestone)
	require.NoError(t, err)

	result = keeper.GetCount(ctx)
	require.Equal(t, uint64(1), result)

}

func (suite *KeeperTestSuite) TestGetNoAckMilestone() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.MilestoneKeeper

	result := keeper.GetCount(ctx)
	require.Equal(t, uint64(0), result)

	milestoneID := "0000"

	keeper.SetNoAckMilestone(ctx, milestoneID)

	val := keeper.GetNoAckMilestone(ctx, "0000")
	require.True(t, val)

	val = keeper.GetNoAckMilestone(ctx, "00001")
	require.False(t, val)

	val = keeper.GetNoAckMilestone(ctx, "")
	require.False(t, val)

	milestoneID = "0001"
	keeper.SetNoAckMilestone(ctx, milestoneID)

	val = keeper.GetNoAckMilestone(ctx, "0001")
	require.True(t, val)

	val = keeper.GetNoAckMilestone(ctx, "0000")
	require.True(t, val)

}

func (suite *KeeperTestSuite) TestLastNoAckMilestone() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.MilestoneKeeper

	result := keeper.GetCount(ctx)
	require.Equal(t, uint64(0), result)

	milestoneID := "0000"

	val := keeper.GetLastNoAckMilestone(ctx)
	require.NotEqual(t, val, milestoneID)

	keeper.SetNoAckMilestone(ctx, milestoneID)

	val = keeper.GetLastNoAckMilestone(ctx)
	require.Equal(t, val, milestoneID)

	milestoneID = "0001"

	keeper.SetNoAckMilestone(ctx, milestoneID)

	val = keeper.GetLastNoAckMilestone(ctx)
	require.Equal(t, val, milestoneID)

}

func (suite *KeeperTestSuite) TestHasStoreValue() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.CheckpointKeeper
	key := milestone.MilestoneKey
	result := keeper.HasStoreValue(ctx, key)
	require.True(t, result)
}
