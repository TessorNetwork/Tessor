package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/TessorNetwork/Tessor/app/apptesting"
	"github.com/TessorNetwork/Tessor/x/records/keeper"
	"github.com/TessorNetwork/Tessor/x/records/types"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
}

// Dynamically gets the MsgServer for this module's keeper
// this function must be used so that the MsgServer is always created with the most updated App context
//	which can change depending on the type of test
//	(e.g. tests with only one Tessor chain vs tests with multiple chains and IBC support)
func (s *KeeperTestSuite) GetMsgServer() types.MsgServer {
	return keeper.NewMsgServerImpl(s.App.RecordsKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
