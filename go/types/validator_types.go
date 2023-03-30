package types

import "time"

type Validator struct {
	OperatorAddress string `json:"operator_address"`
	ConsensusPubkey struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	} `json:"consensus_pubkey"`
	Jailed          bool   `json:"jailed"`
	Status          string `json:"status"`
	Tokens          string `json:"tokens"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Moniker         string `json:"moniker"`
		Identity        string `json:"identity"`
		Website         string `json:"website"`
		SecurityContact string `json:"security_contact"`
		Details         string `json:"details"`
	} `json:"description"`
	UnbondingHeight string    `json:"unbonding_height"`
	UnbondingTime   time.Time `json:"unbonding_time"`
	Commission      struct {
		CommissionRates struct {
			Rate          string `json:"rate"`
			MaxRate       string `json:"max_rate"`
			MaxChangeRate string `json:"max_change_rate"`
		} `json:"commission_rates"`
		UpdateTime time.Time `json:"update_time"`
	} `json:"commission"`
	MinSelfDelegation       string `json:"min_self_delegation"`
	UnbondingOnHoldRefCount string `json:"unbonding_on_hold_ref_count"`
	UnbondingIds            []any  `json:"unbonding_ids"`
}

type ValidatorsResponse struct {
	Validators []Validator `json:"validators"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}

type ValidatorSetValidator struct {
	Address string `json:"address"`
	PubKey  struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	} `json:"pub_key"`
	VotingPower      string `json:"voting_power"`
	ProposerPriority string `json:"proposer_priority"`
}

type ValidatorSetResponse struct {
	BlockHeight string                  `json:"block_height"`
	Validators  []ValidatorSetValidator `json:"validators"`
	Pagination  struct {
		NextKey any    `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}
