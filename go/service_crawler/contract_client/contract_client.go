package contract_client

import (
	"bytes"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/shifty11/dao-dao-notifier/log"
	"net/http"
	"strconv"
	"time"
)

type ContractConfig struct {
	Name                  string
	Description           string
	Threshold             interface{}
	MaxVotingPeriod       interface{} `json:"max_voting_period"`
	ProposalDeposit       string      `json:"proposal_deposit"`
	RefundFailedProposals bool        `json:"refund_failed_proposals"`
	ImageUrl              string      `json:"image_url"`
}

type Config struct {
	Config          ContractConfig `json:"config"`
	GovToken        string         `json:"gov_token"`
	StakingContract string         `json:"staking_contract"`
}

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

type Proposal struct {
	Id            int
	Proposer      string
	Expires       Expires
	Title         string
	Description   string
	Msgs          interface{}
	Status        string
	Threshold     interface{}
	DepositAmount int
}

type ProposalList struct {
	Proposals []Proposal `json:"proposals"`
}

type ContractClient struct {
	ContractAddress string
}

func NewContractClient(contractAddress string) *ContractClient {
	return &ContractClient{
		ContractAddress: contractAddress,
	}
}

func (cc *ContractClient) Config() *Config {
	values := map[string]string{"contractAddress": cc.ContractAddress}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	resp, err := http.Post("http://localhost:8080/config", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querieng /config (%v): %v", resp.StatusCode, resp.Body)
		return nil
	}

	var config Config
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	return &config
}

func (cc *ContractClient) Proposals() *ProposalList {
	values := map[string]string{"contractAddress": cc.ContractAddress}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	resp, err := http.Post("http://localhost:8080/proposals", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querieng /proposals (%v): %v", resp.StatusCode, resp.Body)
		return nil
	}

	var proposals ProposalList
	err = json.NewDecoder(resp.Body).Decode(&proposals)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	return &proposals
}

func (cc *ContractClient) Proposal(proposalId int) *Proposal {
	values := map[string]string{"contractAddress": cc.ContractAddress, "id": strconv.Itoa(proposalId)}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	resp, err := http.Post("http://localhost:8080/proposal", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querieng /proposals (%v): %v", resp.StatusCode, resp.Body)
		return nil
	}

	var proposal Proposal
	err = json.NewDecoder(resp.Body).Decode(&proposal)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	return &proposal
}
