package types

type Chain struct {
	ChainId      string `json:"chain_id"`
	Name         string `json:"name"`
	PrettyName   string `json:"pretty_name"`
	Path         string `json:"path"`
	Display      string `json:"display"`
	NetworkType  string `json:"network_type"`
	Image        string `json:"image"`
	Bech32Prefix string `json:"bech32_prefix"`
}

type ChainInfo struct {
	Chains []Chain `json:"chains"`
}
