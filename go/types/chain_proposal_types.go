package types

import (
	"encoding/json"
	cosmossdktypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"time"
)

type ChainProposalStatus cosmossdktypes.ProposalStatus

func (s *ChainProposalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(cosmossdktypes.ProposalStatus(*s).String())
}

func (s *ChainProposalStatus) UnmarshalJSON(data []byte) error {
	var name = ""
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	*s = ChainProposalStatus(cosmossdktypes.ProposalStatus_value[name])
	return nil
}

func (s *ChainProposalStatus) String() string {
	return cosmossdktypes.ProposalStatus(*s).String()
}

func (s *ChainProposalStatus) ToEntStatus() chainproposal.Status {
	return chainproposal.Status(s.String())
}

type ChainProposalContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ChainProposal struct {
	ProposalId      int                  `json:"proposal_id,string"`
	Content         ChainProposalContent `json:"content"`
	Status          ChainProposalStatus  `json:"status"`
	VotingStartTime time.Time            `json:"voting_start_time"`
	VotingEndTime   time.Time            `json:"voting_end_time"`
}

type Pagination struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	NextKey    int `json:"next_key"`
}

type ChainProposalsResponse struct {
	Proposals  []ChainProposal `json:"proposals"`
	Pagination Pagination      `json:"pagination"`
}

type ChainProposalResponse struct {
	Proposal ChainProposal `json:"proposal"`
}

type ChainProposalVoteOption cosmossdktypes.VoteOption

func (o *ChainProposalVoteOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(cosmossdktypes.VoteOption(*o).String())
}

func (o *ChainProposalVoteOption) UnmarshalJSON(data []byte) error {
	var name string
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	*o = ChainProposalVoteOption(cosmossdktypes.VoteOption_value[name])
	return nil
}

func (o ChainProposalVoteOption) ToCosmosType() cosmossdktypes.VoteOption {
	return cosmossdktypes.VoteOption(o)
}

type ChainProposalVoteResponse struct {
	Vote struct {
		ProposalID string                  `json:"proposal_id"`
		Voter      string                  `json:"voter"`
		Option     ChainProposalVoteOption `json:"option"`
		Options    []struct {
			Option ChainProposalVoteOption `json:"option"`
			Weight string                  `json:"weight"`
		} `json:"options"`
	} `json:"vote"`
}
