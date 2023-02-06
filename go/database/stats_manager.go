package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
)

type StatsManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewStatsManager(client *ent.Client, ctx context.Context) *StatsManager {
	return &StatsManager{client: client, ctx: ctx}
}

type Stats struct {
	Chains                int
	Contracts             int
	Users                 int
	TelegramUsers         int
	DiscordUsers          int
	TelegramChats         int
	DiscordChannels       int
	Subscriptions         int
	TelegramSubscriptions int
	DiscordSubscriptions  int
}

func (m *StatsManager) GetStats() (*Stats, error) {
	cntChains := m.client.Chain.
		Query().
		CountX(m.ctx)
	cntContracts := m.client.Contract.
		Query().
		CountX(m.ctx)
	cntUsers := m.client.User.
		Query().
		CountX(m.ctx)
	cntTelegramUsers := m.client.User.
		Query().
		Where(user.TypeEQ(user.TypeTelegram)).
		CountX(m.ctx)
	cntDiscordUsers := m.client.User.
		Query().
		Where(user.TypeEQ(user.TypeDiscord)).
		CountX(m.ctx)
	cntTelegramChats := m.client.TelegramChat.
		Query().
		CountX(m.ctx)
	cntDiscordChannels := m.client.DiscordChannel.
		Query().
		CountX(m.ctx)

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

	cntTelegramSubs := 0
	cntDiscordSubs := 0
	for _, s := range tgStats {
		cntTelegramSubs += s.Cnt
	}
	for _, s := range dStats {
		cntDiscordSubs += s.Cnt
	}
	cntTotalSubs := cntTelegramSubs + cntDiscordSubs

	return &Stats{
		Chains:                cntChains,
		Contracts:             cntContracts,
		Users:                 cntUsers,
		TelegramUsers:         cntTelegramUsers,
		DiscordUsers:          cntDiscordUsers,
		TelegramChats:         cntTelegramChats,
		DiscordChannels:       cntDiscordChannels,
		Subscriptions:         cntTotalSubs,
		TelegramSubscriptions: cntTelegramSubs,
		DiscordSubscriptions:  cntDiscordSubs,
	}, nil
}
