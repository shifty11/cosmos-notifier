package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/log"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/cosmos-notifier/ent"
)

type AddressTrackerManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewAddressTrackerManager(client *ent.Client, ctx context.Context) *AddressTrackerManager {
	return &AddressTrackerManager{client: client, ctx: ctx}
}

func (manager *AddressTrackerManager) IsValid(address string) bool {
	if address == "" {
		return false
	}

	chains, err := manager.client.Chain.
		Query().
		Where(chain.Bech32PrefixHasPrefix(address[:1])).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Error(err)
		return false
	}
	for _, chainDto := range chains {
		if _, err := cosmossdk.GetFromBech32(address, chainDto.Bech32Prefix); err == nil {
			return true
		}
	}
	return false
}
