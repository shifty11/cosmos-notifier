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
	Name        string
	Description string
	ImageUrl    string
}

func (c *Config) ToContractData() *ContractData {
	return &ContractData{
		Name:        c.Config.Name,
		Description: c.Config.Description,
		ImageUrl:    c.Config.ImageUrl,
	}
}

func (c *ConfigV1) ToContractData() *ContractData {
	return &ContractData{
		Name:        c.Name,
		Description: c.Description,
		ImageUrl:    c.ImageUrl,
	}
}
