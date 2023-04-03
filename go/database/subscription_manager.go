package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/subscription_service"
)

type SubscriptionManager struct {
	client                *ent.Client
	ctx                   context.Context
	userManager           *UserManager
	chainManager          *ChainManager
	contractManager       IContractManager
	telegramChatManager   ITelegramChatManager
	discordChannelManager IDiscordChannelManager
}

func NewSubscriptionManager(
	client *ent.Client,
	ctx context.Context,
	userManager *UserManager,
	chainManager *ChainManager,
	contractManager IContractManager,
	telegramChatManager ITelegramChatManager,
	discordChannelManager IDiscordChannelManager,
) *SubscriptionManager {
	return &SubscriptionManager{
		client:                client,
		ctx:                   ctx,
		userManager:           userManager,
		chainManager:          chainManager,
		contractManager:       contractManager,
		telegramChatManager:   telegramChatManager,
		discordChannelManager: discordChannelManager,
	}
}

func (m *SubscriptionManager) querySubscriptions(ofUser []int, qType queryType, isAdmin bool) []*pb.Subscription {
	var subs []*pb.Subscription
	if qType == chainQuery {
		query := m.chainManager.QueryEnabled()
		if isAdmin {
			query = m.chainManager.QueryAll()
		}
		for _, c := range query {
			var subscription = pb.Subscription{
				Id:              int64(c.ID),
				Name:            c.PrettyName,
				IsSubscribed:    false,
				IsEnabled:       c.IsEnabled,
				ThumbnailUrl:    c.ThumbnailURL,
				ContractAddress: "",
			}
			for _, nc := range ofUser { // check if user gets notified for this contract
				if nc == c.ID {
					subscription.IsSubscribed = true
				}
			}
			subs = append(subs, &subscription)
		}
	} else {
		for _, c := range m.contractManager.QueryAll() {
			var subscription = pb.Subscription{
				Id:              int64(c.ID),
				Name:            c.Name,
				IsSubscribed:    false,
				IsEnabled:       true,
				ThumbnailUrl:    c.ThumbnailURL,
				ContractAddress: c.Address,
			}
			for _, nc := range ofUser { // check if user gets notified for this contract
				if nc == c.ID {
					subscription.IsSubscribed = true
				}
			}
			subs = append(subs, &subscription)
		}
	}
	return subs
}

func (m *SubscriptionManager) UpdateToggleChainSubscription(entUser *ent.User, chatRoomId int64, chainId int) (bool, error) {
	if entUser.Type == user.TypeTelegram {
		return m.telegramChatManager.UpdateAddOrRemoveChain(chatRoomId, chainId)
	} else {
		return m.discordChannelManager.UpdateAddOrRemoveChain(chatRoomId, chainId)
	}
}

func (m *SubscriptionManager) UpdateToggleContractSubscription(entUser *ent.User, chatRoomId int64, contractId int) (bool, error) {
	if entUser.Type == user.TypeTelegram {
		return m.telegramChatManager.UpdateAddOrRemoveContract(chatRoomId, contractId)
	} else {
		return m.discordChannelManager.UpdateAddOrRemoveContract(chatRoomId, contractId)
	}
}

type queryType string

const (
	chainQuery    queryType = "chain"
	contractQuery queryType = "contract"
)

func (m *SubscriptionManager) queryTelegram(entUser *ent.User, qType queryType) []*ent.TelegramChat {
	q := entUser.
		QueryTelegramChats().
		Order(ent.Asc(telegramchat.FieldName)).
		Order(ent.Asc(contract.FieldName))
	if qType == chainQuery {
		q = q.WithChains()
	} else {
		q = q.WithContracts()
	}
	tgChats, err := q.All(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying telegram chats of user %v (%v): %v", entUser.Name, entUser.ID, err)
	}
	return tgChats
}

func (m *SubscriptionManager) queryDiscord(entUser *ent.User, qType queryType) []*ent.DiscordChannel {
	q := entUser.
		QueryDiscordChannels().
		Order(ent.Asc(telegramchat.FieldName)).
		Order(ent.Asc(contract.FieldName))
	if qType == chainQuery {
		q = q.WithChains()
	} else {
		q = q.WithContracts()
	}
	discordChannels, err := q.All(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying discord channels of user %v (%v): %v", entUser.Name, entUser.ID, err)
	}
	return discordChannels
}

func (m *SubscriptionManager) getIdsOfTelegramChat(tgChat *ent.TelegramChat, qType queryType) []int {
	var ids []int
	if qType == chainQuery {
		for _, c := range tgChat.Edges.Chains {
			ids = append(ids, c.ID)
		}
	} else {
		for _, c := range tgChat.Edges.Contracts {
			ids = append(ids, c.ID)
		}
	}
	return ids
}

func (m *SubscriptionManager) getIdsOfDiscordChannel(discordChannel *ent.DiscordChannel, qType queryType) []int {
	var ids []int
	if qType == chainQuery {
		for _, c := range discordChannel.Edges.Chains {
			ids = append(ids, c.ID)
		}
	} else {
		for _, c := range discordChannel.Edges.Contracts {
			ids = append(ids, c.ID)
		}
	}
	return ids
}

func (m *SubscriptionManager) telegramChatsToChatRoom(tgChats []*ent.TelegramChat, qType queryType, isAdmin bool) []*pb.ChatRoom {
	var chats []*pb.ChatRoom
	for _, tgChat := range tgChats {
		chats = append(chats, &pb.ChatRoom{
			Id:            tgChat.ChatID,
			Name:          tgChat.Name,
			Type:          pb.ChatRoom_TELEGRAM,
			Subscriptions: m.querySubscriptions(m.getIdsOfTelegramChat(tgChat, qType), qType, isAdmin),
		})
		if isAdmin && qType == chainQuery {
			m.collectChainStats(chats)
		} else if isAdmin && qType == contractQuery {
			m.collectContractStats(chats)
		}
	}
	return chats
}

func (m *SubscriptionManager) discordChannelsToChatRoom(discordChannels []*ent.DiscordChannel, qType queryType, isAdmin bool) []*pb.ChatRoom {
	var chats []*pb.ChatRoom
	for _, dChannel := range discordChannels {
		chats = append(chats, &pb.ChatRoom{
			Id:            dChannel.ChannelID,
			Name:          dChannel.Name,
			Type:          pb.ChatRoom_DISCORD,
			Subscriptions: m.querySubscriptions(m.getIdsOfDiscordChannel(dChannel, qType), qType, isAdmin),
		})
		if isAdmin && qType == chainQuery {
			m.collectChainStats(chats)
		} else if isAdmin && qType == contractQuery {
			m.collectContractStats(chats)
		}
	}
	return chats
}

func (m *SubscriptionManager) QuerySubscriptions(entUser *ent.User) *pb.ListSubscriptionsResponse {
	isAdmin := entUser.Role == user.RoleAdmin
	if entUser.Type == user.TypeTelegram {
		return &pb.ListSubscriptionsResponse{
			ChainChatRooms:    m.telegramChatsToChatRoom(m.queryTelegram(entUser, chainQuery), chainQuery, isAdmin),
			ContractChatRooms: m.telegramChatsToChatRoom(m.queryTelegram(entUser, contractQuery), contractQuery, isAdmin),
		}
	} else {
		return &pb.ListSubscriptionsResponse{
			ChainChatRooms:    m.discordChannelsToChatRoom(m.queryDiscord(entUser, chainQuery), chainQuery, isAdmin),
			ContractChatRooms: m.discordChannelsToChatRoom(m.queryDiscord(entUser, contractQuery), contractQuery, isAdmin),
		}
	}
}

// collectContractStats takes a list of ChatRooms and fills the stats for each subscription
func (m *SubscriptionManager) collectContractStats(chats []*pb.ChatRoom) {
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
			if sub.Stats == nil {
				sub.Stats = &pb.SubscriptionStats{}
			}
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

// collectChainStats takes a list of ChatRooms and fills the stats for each subscription
func (m *SubscriptionManager) collectChainStats(chats []*pb.ChatRoom) {
	type stats []struct {
		ID  int
		Cnt int
	}
	var tgStats = stats{}
	err := m.client.Chain.
		Query().
		GroupBy(chain.FieldID).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(chain.TelegramChatsTable)
			s.Join(t).On(s.C(chain.FieldID), t.C(chain.TelegramChatsPrimaryKey[1]))
			return sql.As(sql.Count(t.C(chain.TelegramChatsPrimaryKey[1])), "cnt")
		}).
		Scan(m.ctx, &tgStats)
	if err != nil {
		log.Sugar.Errorf("Error while getting tgStats: %v", err)
	}

	var dStats = stats{}
	err = m.client.Chain.
		Query().
		GroupBy(chain.FieldID).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(chain.DiscordChannelsTable)
			s.Join(t).On(s.C(chain.FieldID), t.C(chain.DiscordChannelsPrimaryKey[1]))
			return sql.As(sql.Count(t.C(chain.DiscordChannelsPrimaryKey[1])), "cnt")
		}).
		Scan(m.ctx, &dStats)
	if err != nil {
		log.Sugar.Errorf("Error while getting dStats: %v", err)
	}

	for _, chat := range chats {
		for _, sub := range chat.Subscriptions {
			if sub.Stats == nil {
				sub.Stats = &pb.SubscriptionStats{}
			}
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
