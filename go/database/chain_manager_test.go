package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
)

func newTestChainManager(t *testing.T) *ChainManager {
	manager := NewChainManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addChains(manager *ChainManager) []*ent.Chain {
	var chains []*ent.Chain
	chainDto, _ := manager.Create(&types.Chain{
		ChainId:      "cosmoshub-3",
		Name:         "Cosmos",
		PrettyName:   "Cosmos Hub",
		Path:         "cosmos",
		Display:      "Cosmos Hub",
		NetworkType:  "mainnet",
		Image:        "",
		Bech32Prefix: "cosmos",
	}, "")
	chains = append(chains, chainDto)
	chainDto, _ = manager.Create(&types.Chain{
		ChainId:      "osmosis-1",
		Name:         "Osmosis",
		Path:         "osmosis",
		Display:      "Osmosis",
		NetworkType:  "mainnet",
		Image:        "",
		Bech32Prefix: "osmo",
	}, "")
	chains = append(chains, chainDto)
	chainDto, _ = manager.Create(&types.Chain{
		ChainId:      "comdex-1",
		Name:         "Comdex",
		PrettyName:   "Comdex",
		Path:         "comdex",
		Display:      "Comdex",
		NetworkType:  "mainnet",
		Image:        "",
		Bech32Prefix: "comdex",
	}, "")
	chains = append(chains, chainDto)
	return chains
}
