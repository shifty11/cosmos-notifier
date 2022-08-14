package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Expires struct {
	AtTime time.Time `json:"at_time"`
}

func (s *Expires) UnmarshalJSON(data []byte) error {
	var v struct {
		AtTime string `json:"at_time"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	timestamp, err := strconv.ParseInt(v.AtTime, 10, 64)
	if err != nil {
		return err
	}
	s.AtTime = time.Unix(0, timestamp)
	return nil
}

type ProposalStatus string

const (
	StatusOpen            ProposalStatus = "open"
	StatusRejected        ProposalStatus = "rejected"
	StatusPassed          ProposalStatus = "passed"
	StatusExecuted        ProposalStatus = "executed"
	StatusClosed          ProposalStatus = "closed"
	StatusExecutionFailed ProposalStatus = "execution_failed"
)

var ProposalStatusValues = []ProposalStatus{
	StatusOpen, StatusRejected, StatusPassed, StatusExecuted, StatusClosed, StatusExecutionFailed,
}

func (s *ProposalStatus) UnmarshalJSON(data []byte) error {
	*s = ProposalStatus(strings.Replace(string(data), "\"", "", -1))
	return nil
}

type Proposal struct {
	Id          int
	Proposer    string
	Expires     Expires
	Title       string
	Description string
	//Msgs          interface{}
	Status ProposalStatus
	//Threshold     interface{}
	//DepositAmount int
}

type ProposalList struct {
	Proposals []Proposal `json:"proposals"`
}

type ProposalV1Object struct {
	Proposer    string
	Title       string
	Description string
	Expiration  Expires
	Msgs        interface{}
	Status      ProposalStatus
	Threshold   interface{}
}

type ProposalV1 struct {
	Id       int
	Proposal ProposalV1Object
}

type ProposalListV1 struct {
	Proposals []ProposalV1 `json:"proposals"`
}

func (p *ProposalV1) ToProposal() *Proposal {
	return &Proposal{
		Id:          p.Id,
		Proposer:    p.Proposal.Proposer,
		Title:       p.Proposal.Title,
		Description: p.Proposal.Description,
		Expires:     p.Proposal.Expiration,
		Status:      p.Proposal.Status,
	}
}

func (l *ProposalListV1) ToProposalList() *ProposalList {
	proposals := make([]Proposal, len(l.Proposals))
	for i, p := range l.Proposals {
		proposals[i] = *p.ToProposal()
	}
	return &ProposalList{
		Proposals: proposals,
	}
}
