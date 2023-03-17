package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
)

func newTestChainManager(t *testing.T) *ChainManager {
	manager := NewChainManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addChains(manager *ChainManager) {
	manager.Create(&types.Chain{
		ChainId:      "cosmoshub-3",
		Name:         "Cosmos",
		PrettyName:   "Cosmos Hub",
		Path:         "cosmos",
		Display:      "Cosmos Hub",
		NetworkType:  "mainnet",
		Image:        "",
		Bech32Prefix: "cosmos",
	}, "")
	manager.Create(&types.Chain{
		ChainId:      "osmosis-1",
		Name:         "Osmosis",
		Path:         "osmosis",
		Display:      "Osmosis",
		NetworkType:  "mainnet",
		Image:        "",
		Bech32Prefix: "osmo",
	}, "")
	manager.Create(&types.Chain{
		ChainId:      "comdex-1",
		Name:         "Comdex",
		PrettyName:   "Comdex",
		Path:         "comdex",
		Display:      "Comdex",
		NetworkType:  "mainnet",
		Image:        "",
		Bech32Prefix: "comdex",
	}, "")
}
