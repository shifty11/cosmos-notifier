package types

type Chain struct {
	ChainId     string `json:"chain_id"`
	Name        string `json:"name"`
	PrettyName  string `json:"pretty_name"`
	NetworkType string `json:"network_type"`
	Image       string `json:"image"`
}

type ChainInfo struct {
	Chains []Chain `json:"chains"`
}
