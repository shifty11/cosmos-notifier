package types

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

type ConfigV1 struct {
	Name        string
	Description string
	ImageUrl    string `json:"image_url"`
}

type ContractData struct {
	Address     string
	Name        string
	Description string
	ImageUrl    string
}

func (c *Config) ToContractData(address string) *ContractData {
	data := &ContractData{
		Address:     address,
		Name:        c.Config.Name,
		Description: c.Config.Description,
		ImageUrl:    c.Config.ImageUrl,
	}
	if data.Name == "" {
		data.Name = address
	}
	return data
}

func (c *ConfigV1) ToContractData(address string) *ContractData {
	data := &ContractData{
		Address:     address,
		Name:        c.Name,
		Description: c.Description,
		ImageUrl:    c.ImageUrl,
	}
	if data.Name == "" {
		data.Name = address
	}
	return data
}
