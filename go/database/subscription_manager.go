package database

import (
	"context"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/chain"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/telegramchat"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/ent/userwithzeroid"
	"github.com/shifty11/dao-dao-notifier/log"
	pb "github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/subscription_service"
	"io"
	"os"
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

func (m *SubscriptionManager) getSubscriptions(ofUser []int, qType queryType, isAdmin bool) []*pb.Subscription {
	var subs []*pb.Subscription
	if qType == chainQuery {
		query := m.chainManager.Enabled()
		if isAdmin {
			query = m.chainManager.All()
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
		for _, c := range m.contractManager.All() {
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

func (m *SubscriptionManager) ToggleChainSubscription(entUser *ent.User, chatRoomId int64, chainId int) (bool, error) {
	if entUser.Type == user.TypeTelegram {
		return m.telegramChatManager.AddOrRemoveChain(chatRoomId, chainId)
	} else {
		return m.discordChannelManager.AddOrRemoveChain(chatRoomId, chainId)
	}
}

func (m *SubscriptionManager) ToggleContractSubscription(entUser *ent.User, chatRoomId int64, contractId int) (bool, error) {
	if entUser.Type == user.TypeTelegram {
		return m.telegramChatManager.AddOrRemoveContract(chatRoomId, contractId)
	} else {
		return m.discordChannelManager.AddOrRemoveContract(chatRoomId, contractId)
	}
}

type queryType string

const (
	chainQuery    queryType = "chain"
	contractQuery queryType = "contract"
)

func (m *SubscriptionManager) telegramQuery(entUser *ent.User, qType queryType) []*ent.TelegramChat {
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

func (m *SubscriptionManager) discordQuery(entUser *ent.User, qType queryType) []*ent.DiscordChannel {
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
			TYPE:          pb.ChatRoom_TELEGRAM,
			Subscriptions: m.getSubscriptions(m.getIdsOfTelegramChat(tgChat, qType), qType, isAdmin),
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
			TYPE:          pb.ChatRoom_DISCORD,
			Subscriptions: m.getSubscriptions(m.getIdsOfDiscordChannel(dChannel, qType), qType, isAdmin),
		})
		if isAdmin && qType == chainQuery {
			m.collectChainStats(chats)
		} else if isAdmin && qType == contractQuery {
			m.collectContractStats(chats)
		}
	}
	return chats
}

func (m *SubscriptionManager) GetSubscriptions(entUser *ent.User) *pb.GetSubscriptionsResponse {
	isAdmin := entUser.Role == user.RoleAdmin
	if entUser.Type == user.TypeTelegram {
		return &pb.GetSubscriptionsResponse{
			ChainChatRooms:    m.telegramChatsToChatRoom(m.telegramQuery(entUser, chainQuery), chainQuery, isAdmin),
			ContractChatRooms: m.telegramChatsToChatRoom(m.telegramQuery(entUser, contractQuery), contractQuery, isAdmin),
		}
	} else {
		return &pb.GetSubscriptionsResponse{
			ChainChatRooms:    m.discordChannelsToChatRoom(m.discordQuery(entUser, chainQuery), chainQuery, isAdmin),
			ContractChatRooms: m.discordChannelsToChatRoom(m.discordQuery(entUser, contractQuery), contractQuery, isAdmin),
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

// TODO: remove after migration
type TelegramUser struct {
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	ChatId   int64  `json:"chat_id"`
	ChatName string `json:"chat_name"`
	IsGroup  bool   `json:"is_group"`
	ChainId  string `json:"chain_id"`
}

// TODO: remove after migration
type DiscordUser struct {
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	ChannelId   int64  `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	IsGroup     bool   `json:"is_group"`
	ChainId     string `json:"chain_id"`
}

// TODO: remove after migration
func (manager *SubscriptionManager) ImportDb() {
	var tgUsers []TelegramUser
	var discordUsers []DiscordUser
	readFromFile("telegram_users.json", &tgUsers)
	readFromFile("discord_users.json", &discordUsers)

	chains := manager.chainManager.All()

	for _, tgUser := range tgUsers {
		if tgUser.ChatId != 0 {
			tc, _ := manager.telegramChatManager.CreateOrUpdateChat(tgUser.UserId, tgUser.Name, tgUser.ChatId, tgUser.ChatName, tgUser.IsGroup)
			if tgUser.ChainId != "" {
				for _, c := range chains {
					if c.ChainID == tgUser.ChainId {
						err := tc.Update().AddChains(c).Exec(context.Background())
						if err != nil && ent.IsConstraintError(err) {
							log.Sugar.Infof("Telegram chat %v already in c %v", tc.ID, c.ID)
						} else if err != nil {
							log.Sugar.Panicf("Could not add c to telegram chat: %v", err)
						}
					}
				}
			}
		}
	}
	for _, discordUser := range discordUsers {
		if discordUser.ChannelId != 0 {
			dc, _ := manager.discordChannelManager.CreateOrUpdateChannel(discordUser.UserId, discordUser.Name, discordUser.ChannelId, discordUser.ChannelName, discordUser.IsGroup)
			if discordUser.ChainId != "" {
				for _, c := range chains {
					if c.ChainID == discordUser.ChainId {
						err := dc.Update().AddChains(c).Exec(context.Background())
						if err != nil && ent.IsConstraintError(err) {
							log.Sugar.Infof("Discord channel %v already in c %v", dc.ID, c.ID)
						} else if err != nil {
							log.Sugar.Panicf("Could not add c to discord channel: %v", err)
						}
					}
				}
			}
		}
	}

	for _, tgUser := range tgUsers {
		exists, err := manager.client.UserWithZeroId.
			Query().
			Where(userwithzeroid.ChainIDEQ(tgUser.ChainId)).
			Where(userwithzeroid.ChatOrChannelNameEQ(tgUser.ChatName)).
			Where(userwithzeroid.ChatOrChannelIDEQ(tgUser.ChatId)).
			Where(userwithzeroid.IsGroupEQ(tgUser.IsGroup)).
			Where(userwithzeroid.TypeEQ(userwithzeroid.TypeTelegram)).
			Exist(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while checking if user exists: %v", err)
		}
		if !exists {
			_, err := manager.client.UserWithZeroId.
				Create().
				SetChainID(tgUser.ChainId).
				SetChatOrChannelName(tgUser.ChatName).
				SetChatOrChannelID(tgUser.ChatId).
				SetChainID(tgUser.ChainId).
				SetIsGroup(tgUser.IsGroup).
				SetType(userwithzeroid.TypeTelegram).
				Save(manager.ctx)
			if err != nil {
				log.Sugar.Panic(err)
			}
		}
	}

	for _, discordUser := range discordUsers {
		exists, err := manager.client.UserWithZeroId.
			Query().
			Where(userwithzeroid.ChainIDEQ(discordUser.ChainId)).
			Where(userwithzeroid.ChatOrChannelNameEQ(discordUser.ChannelName)).
			Where(userwithzeroid.ChatOrChannelIDEQ(discordUser.ChannelId)).
			Where(userwithzeroid.IsGroupEQ(discordUser.IsGroup)).
			Where(userwithzeroid.TypeEQ(userwithzeroid.TypeDiscord)).
			Exist(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while checking if user exists: %v", err)
		}
		if !exists {
			_, err := manager.client.UserWithZeroId.
				Create().
				SetChainID(discordUser.ChainId).
				SetChatOrChannelName(discordUser.ChannelName).
				SetChatOrChannelID(discordUser.ChannelId).
				SetChainID(discordUser.ChainId).
				SetIsGroup(discordUser.IsGroup).
				SetType(userwithzeroid.TypeDiscord).
				Save(manager.ctx)
			if err != nil {
				log.Sugar.Panic(err)
			}
		}
	}
}

// TODO: remove after migration
func (manager *SubscriptionManager) ImportDbZeroIds() {
	var zeroIdTgUsers []TelegramUser
	var zeroIdDiscordUsers []DiscordUser
	readFromFile("telegram_users_zero_id.json", &zeroIdTgUsers)
	readFromFile("discord_users_zero_id.json", &zeroIdDiscordUsers)

	for _, tgUser := range zeroIdTgUsers {
		if tgUser.ChatId == 0 {
			// TODO: send notification to user
		}
	}
	for _, discordUser := range zeroIdDiscordUsers {
		if discordUser.ChannelId == 0 {
			// TODO: send notification to user
		}
	}
}

// TODO: remove after migration
func readFromFile(fileName string, target interface{}) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Sugar.Panicf("Could not open file %s: %v", fileName, err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Sugar.Panicf("Could not read file %s: %v", fileName, err)
	}
	err = json.Unmarshal(byteValue, &target)
	if err != nil {
		log.Sugar.Panicf("Could not unmarshal file %s: %v", fileName, err)
	}
}
