package database

import (
	"context"
	"errors"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
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

func (manager *AddressTrackerManager) AddTracker(userEnt *ent.User, address string, discordChannelId int, telegramChatId int) (*ent.AddressTracker, error) {
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
		exist, err := userEnt.QueryDiscordChannels().
			Where(discordchannel.IDEQ(discordChannelId)).
			Exist(manager.ctx)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.New("discord channel not found")
		}
		createQuery.SetDiscordChannelID(discordChannelId)
	} else {
		exist, err := userEnt.QueryTelegramChats().
			Where(telegramchat.IDEQ(telegramChatId)).
			Exist(manager.ctx)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.New("telegram chat not found")
		}
		createQuery.SetTelegramChatID(telegramChatId)
	}

	addressTrackerDto, err := createQuery.Save(manager.ctx)
	if err != nil {
		return nil, err
	}

	return addressTrackerDto, err
}
