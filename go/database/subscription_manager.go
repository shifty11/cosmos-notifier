package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
)

type Subscription struct {
	Id              int64
	Name            string
	Notify          bool
	ThumbnailUrl    string
	ContractAddress string
}

type ChatRoom struct {
	Id            int64
	Name          string
	Subscriptions []*Subscription
}

type SubscriptionManager struct {
	client                *ent.Client
	ctx                   context.Context
	userManager           *UserManager
	contractManager       IContractManager
	telegramChatManager   *TelegramChatManager
	discordChannelManager *DiscordChannelManager
}

func NewSubscriptionManager(
	client *ent.Client,
	ctx context.Context,
	userManager *UserManager,
	contractManager IContractManager,
	telegramChatManager *TelegramChatManager,
	discordChannelManager *DiscordChannelManager,
) *SubscriptionManager {
	return &SubscriptionManager{
		client:                client,
		ctx:                   ctx,
		userManager:           userManager,
		contractManager:       contractManager,
		telegramChatManager:   telegramChatManager,
		discordChannelManager: discordChannelManager,
	}
}

func getSubscriptions(contractManager IContractManager, ofUser []*ent.Contract) []*Subscription {
	contracts := contractManager.All()
	var subs []*Subscription
	for _, c := range contracts {
		var subscription = Subscription{
			Id:              int64(c.ID),
			Name:            c.Name,
			Notify:          false,
			ThumbnailUrl:    c.ThumbnailURL,
			ContractAddress: c.Address,
		}
		for _, nc := range ofUser { // check if user gets notified for this contract
			if nc.ID == c.ID {
				subscription.Notify = true
			}
		}
		subs = append(subs, &subscription)
	}
	return subs
}

func (m *SubscriptionManager) ToggleSubscription(entUser *ent.User, chatRoomId int64, contractId int) (bool, error) {
	if entUser.Type == user.TypeTelegram {
		return m.telegramChatManager.AddOrRemoveContract(chatRoomId, contractId)
	} else {
		return m.discordChannelManager.AddOrRemoveContract(chatRoomId, contractId)
	}
}

func (m *SubscriptionManager) GetSubscriptions(entUser *ent.User) []*ChatRoom {
	if entUser.Type == user.TypeTelegram {
		tgChats, err := entUser.
			QueryTelegramChats().
			WithContracts().
			All(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying telegram chats of user %v (%v): %v", entUser.Name, entUser.ID, err)
		}

		var chats []*ChatRoom
		for _, tgChat := range tgChats {
			chats = append(chats, &ChatRoom{
				Id:            tgChat.ChatID,
				Name:          tgChat.Name,
				Subscriptions: getSubscriptions(m.contractManager, tgChat.Edges.Contracts),
			})
		}
		return chats
	} else {
		dChannels, err := entUser.
			QueryDiscordChannels().
			WithContracts().
			All(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying discord channels of user %v (%v): %v", entUser.Name, entUser.ID, err)
		}

		var chats []*ChatRoom
		for _, dChannel := range dChannels {
			chats = append(chats, &ChatRoom{
				Id:            dChannel.ChannelID,
				Name:          dChannel.Name,
				Subscriptions: getSubscriptions(m.contractManager, dChannel.Edges.Contracts),
			})
		}
		return chats
	}
}
