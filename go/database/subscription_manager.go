package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/discordchannel"
	"github.com/shifty11/dao-dao-notifier/ent/telegramchat"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
	pb "github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/subscription_service"
)

type SubscriptionStats struct {
	Total    int
	Telegram int
	Discord  int
}

type Subscription struct {
	Id              int64
	Name            string
	Notify          bool
	ThumbnailUrl    string
	ContractAddress string
	Stats           SubscriptionStats
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
	telegramChatManager   ITelegramChatManager
	discordChannelManager IDiscordChannelManager
}

func NewSubscriptionManager(
	client *ent.Client,
	ctx context.Context,
	userManager *UserManager,
	contractManager IContractManager,
	telegramChatManager ITelegramChatManager,
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

func (m *SubscriptionManager) getSubscriptions(ofUser []*ent.Contract) []*pb.Subscription {
	contracts := m.contractManager.All()
	var subs []*pb.Subscription
	for _, c := range contracts {
		var subscription = pb.Subscription{
			Id:              int64(c.ID),
			Name:            c.Name,
			IsSubscribed:    false,
			ThumbnailUrl:    c.ThumbnailURL,
			ContractAddress: c.Address,
		}
		for _, nc := range ofUser { // check if user gets notified for this contract
			if nc.ID == c.ID {
				subscription.IsSubscribed = true
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

func (m *SubscriptionManager) GetSubscriptions(entUser *ent.User) []*pb.ChatRoom {
	if entUser.Type == user.TypeTelegram {
		tgChats, err := entUser.
			QueryTelegramChats().
			Order(ent.Asc(telegramchat.FieldName)).
			WithContracts().
			Order(ent.Asc(contract.FieldName)).
			All(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying telegram chats of user %v (%v): %v", entUser.Name, entUser.ID, err)
		}

		var chats []*pb.ChatRoom
		for _, tgChat := range tgChats {
			chats = append(chats, &pb.ChatRoom{
				Id:            tgChat.ChatID,
				Name:          tgChat.Name,
				TYPE:          pb.ChatRoom_TELEGRAM,
				Subscriptions: m.getSubscriptions(tgChat.Edges.Contracts),
			})
			if entUser.Role == user.RoleAdmin {
				m.collectStats(chats)
			}
		}
		return chats
	} else {
		dChannels, err := entUser.
			QueryDiscordChannels().
			Order(ent.Asc(discordchannel.FieldName)).
			WithContracts().
			Order(ent.Asc(contract.FieldName)).
			All(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying discord channels of user %v (%v): %v", entUser.Name, entUser.ID, err)
		}

		var chats []*pb.ChatRoom
		for _, dChannel := range dChannels {
			chats = append(chats, &pb.ChatRoom{
				Id:            dChannel.ChannelID,
				Name:          dChannel.Name,
				TYPE:          pb.ChatRoom_DISCORD,
				Subscriptions: m.getSubscriptions(dChannel.Edges.Contracts),
			})
			if entUser.Role == user.RoleAdmin {
				m.collectStats(chats)
			}
		}
		return chats
	}
}

// collectStats takes a list of ChatRoom and fills the stats for each subscription
func (m *SubscriptionManager) collectStats(chats []*pb.ChatRoom) {
	type stats []struct {
		ID  int
		Cnt int
	}
	var tgStats = stats{}
	err := m.client.Contract.
		Query().
		GroupBy(contract.FieldID).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(contract.TelegramChatsTable)
			s.Join(t).On(s.C(contract.FieldID), t.C(contract.TelegramChatsPrimaryKey[1]))
			return sql.As(sql.Count(t.C(contract.TelegramChatsPrimaryKey[1])), "cnt")
		}).
		Scan(m.ctx, &tgStats)
	if err != nil {
		log.Sugar.Errorf("Error while getting tgStats: %v", err)
	}

	var dStats = stats{}
	err = m.client.Contract.
		Query().
		GroupBy(contract.FieldID).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(contract.DiscordChannelsTable)
			s.Join(t).On(s.C(contract.FieldID), t.C(contract.DiscordChannelsPrimaryKey[1]))
			return sql.As(sql.Count(t.C(contract.DiscordChannelsPrimaryKey[1])), "cnt")
		}).
		Scan(m.ctx, &dStats)
	if err != nil {
		log.Sugar.Errorf("Error while getting dStats: %v", err)
	}

	for _, chat := range chats {
		for _, sub := range chat.Subscriptions {
			for _, s := range tgStats {
				if s.ID == int(sub.Id) {
					sub.Stats.Telegram = int32(s.Cnt)
				}
			}
			for _, s := range dStats {
				if s.ID == int(sub.Id) {
					sub.Stats.Discord = int32(s.Cnt)
				}
			}
			sub.Stats.Total = sub.Stats.Telegram + sub.Stats.Discord
		}
	}
}
