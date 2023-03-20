package database

import (
	"context"
	"errors"
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

func (manager *AddressTrackerManager) IsValid(address string) (bool, *ent.Chain) {
	if address == "" {
		return false, nil
	}

	chains, err := manager.client.Chain.
		Query().
		Where(chain.Bech32PrefixHasPrefix(address[:1])).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Error(err)
		return false, nil
	}
	for _, chainDto := range chains {
		if _, err := cosmossdk.GetFromBech32(address, chainDto.Bech32Prefix); err == nil {
			return true, chainDto
		}
	}
	return false, nil
}

func (manager *AddressTrackerManager) AddTracker(address string, discordChannelId int, telegramChatId int) (*ent.AddressTracker, error) {
	isValid, chainEnt := manager.IsValid(address)
	if !isValid {
		return nil, errors.New("invalid address")
	}
	if discordChannelId == 0 && telegramChatId == 0 {
		return nil, errors.New("at least one of discordChannelId or telegramChatId must be non-zero")
	}
	if discordChannelId != 0 && telegramChatId != 0 {
		return nil, errors.New("only one of discordChannelId or telegramChatId must be non-zero")
	}

	createQuery := manager.client.AddressTracker.
		Create().
		SetChain(chainEnt).
		SetAddress(address)

	if discordChannelId != 0 {
		createQuery.SetDiscordChannelID(discordChannelId)
	} else {
		createQuery.SetTelegramChatID(telegramChatId)
	}

	addressTrackerDto, err := createQuery.Save(manager.ctx)
	if err != nil {
		return nil, err
	}

	return addressTrackerDto, err
}
