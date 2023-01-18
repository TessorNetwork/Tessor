package client

import (
	"github.com/TessorNetwork/tessor/v4/x/stakeibc/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	AddValidatorProposalHandler = govclient.NewProposalHandler(cli.CmdAddValidatorProposal)
)
