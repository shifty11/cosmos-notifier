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
	Address         string
	Name            string
	Description     string
	ImageUrl        string
	RpcEndpoint     string
	ContractVersion ContractVersion
}

type ContractVersion string

const (
	ContractVersionUnknown ContractVersion = "unknown"
	ContractVersionV1      ContractVersion = "v1"
	ContractVersionV2      ContractVersion = "v2"
)

func (cv ContractVersion) String() string {
	return string(cv)
}

func (c *Config) ToContractData(address string, rpcEndpoint string) *ContractData {
	data := &ContractData{
		Address:         address,
		Name:            c.Config.Name,
		Description:     c.Config.Description,
		ImageUrl:        c.Config.ImageUrl,
		RpcEndpoint:     rpcEndpoint,
		ContractVersion: ContractVersionV2,
	}
	if data.Name == "" {
		data.Name = address
	}
	return data
}

func (c *ConfigV1) ToContractData(address string, rpcEndpoint string) *ContractData {
	data := &ContractData{
		Address:         address,
		Name:            c.Name,
		Description:     c.Description,
		ImageUrl:        c.ImageUrl,
		RpcEndpoint:     rpcEndpoint,
		ContractVersion: ContractVersionV1,
	}
	if data.Name == "" {
		data.Name = address
	}
	return data
}
