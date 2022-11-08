package database

import (
	"context"
	"errors"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/chain"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/ent/discordchannel"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
)

type IDiscordChannelManager interface {
	AddOrRemoveChain(tgChatId int64, chainId int) (hasContract bool, err error)
	AddOrRemoveContract(dChannelId int64, contractId int) (hasContract bool, err error)
	CreateOrUpdateChannel(userId int64, userName string, channelId int64, name string, isGroup bool) (dc *ent.DiscordChannel, created bool)
	Delete(userId int64, channelId int64) error
	GetChannelUsers(channelId int64) []*ent.User
	CountSubscriptions(channelId int64) int
	GetSubscribedIds(query *ent.DiscordChannelQuery) []types.DiscordChannelQueryResult
	DeleteMultiple(channelIds []int64)
	GetAllIds() []types.DiscordChannelQueryResult
}

type DiscordChannelManager struct {
	client          *ent.Client
	ctx             context.Context
	chainManager    *ChainManager
	contractManager IContractManager
	userManager     *UserManager
}

func NewDiscordChannelManager(
	client *ent.Client,
	ctx context.Context,
	chainManager *ChainManager,
	contractManager IContractManager,
	userManager *UserManager,
) *DiscordChannelManager {
	return &DiscordChannelManager{
		client:          client,
		ctx:             ctx,
		chainManager:    chainManager,
		contractManager: contractManager,
		userManager:     userManager,
	}
}

// AddOrRemoveChain adds or removes a chain from a discord channel
// Returns true if the contract is now added
func (m *DiscordChannelManager) AddOrRemoveChain(dChannelId int64, chainId int) (hasContract bool, err error) {
	log.Sugar.Debugf("Adding or removing chain %d from discord channel %d", chainId, dChannelId)
	dChannel, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(dChannelId)).
		First(m.ctx)
	if err != nil {
		return false, err
	}

	entChain, err := m.chainManager.Get(chainId)
	if err != nil {
		return false, err
	}

	exists, err := dChannel.
		QueryChains().
		Where(chain.IDEQ(entChain.ID)).
		Exist(m.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := dChannel.
			Update().
			RemoveChainIDs(entChain.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := dChannel.
			Update().
			AddChainIDs(entChain.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	}
	return !exists, nil
}

// AddOrRemoveContract adds or removes a contract from a discord channel
// Returns true if the contract is now added
func (m *DiscordChannelManager) AddOrRemoveContract(dChannelId int64, contractId int) (hasContract bool, err error) {
	log.Sugar.Debugf("Adding or removing contract %d from discord channel %d", contractId, dChannelId)
	dChannel, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(dChannelId)).
		First(m.ctx)
	if err != nil {
		return false, err
	}

	dbContract, err := m.contractManager.Get(contractId)
	if err != nil {
		return false, err
	}

	exists, err := dChannel.
		QueryContracts().
		Where(contract.IDEQ(dbContract.ID)).
		Exist(m.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := dChannel.
			Update().
			RemoveContractIDs(dbContract.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := dChannel.
			Update().
			AddContractIDs(dbContract.ID).
			Save(m.ctx)
		if err != nil {
			return false, err
		}
	}
	return !exists, nil
}

// CreateOrUpdateChannel Add adds a new discord channel to the database or updates an existing one
// Returns the channel and a boolean indicating if the channel was created
func (m *DiscordChannelManager) CreateOrUpdateChannel(userId int64, userName string, channelId int64, name string, isGroup bool) (dc *ent.DiscordChannel, created bool) {
	log.Sugar.Debugf("Create or update Discord channel %s (%d) for user %s (%d)", name, channelId, userName, userId)
	entUser := m.userManager.createOrUpdateUser(userId, userName, user.TypeDiscord)
	discordChannel, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(channelId)).
		WithUsers().
		Only(m.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			discordChannel, err = m.client.DiscordChannel.
				Create().
				AddUsers(entUser).
				SetChannelID(channelId).
				SetName(name).
				SetIsGroup(isGroup).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not create discord channel: %v", err)
			}
			return discordChannel, true
		} else {
			log.Sugar.Panicf("Could not find discord channel: %v", err)
		}
	} else {
		hasUser := false
		for _, u := range discordChannel.Edges.Users {
			if u.ID == entUser.ID {
				hasUser = true
				break
			}
		}
		if !hasUser {
			discordChannel, err = discordChannel.
				Update().
				AddUsers(entUser).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not update discord channel: %v", err)
			}
		}
		if discordChannel.Name != name {
			discordChannel, err = discordChannel.
				Update().
				SetName(name).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Could not update discord channel: %v", err)
			}
		}
	}
	return discordChannel, false
}

// Delete deletes a discord channel for a user
// If the user doesn't have any more channels, the user is deleted
func (m *DiscordChannelManager) Delete(userId int64, channelId int64) error {
	log.Sugar.Debugf("Deleting discord channel %d for user %d", channelId, userId)
	discordChannel, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(channelId)).
		WithUsers().
		Only(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not find discord channel: %v", err)
		return err
	}
	var entUser *ent.User
	for _, u := range discordChannel.Edges.Users {
		if u.UserID == userId {
			entUser = u
			break
		}
	}
	if entUser == nil {
		log.Sugar.Errorf("Could not find user: %v", err)
		return errors.New("could not find user")
	}
	if len(discordChannel.Edges.Users) == 1 {
		err := m.client.DiscordChannel.
			DeleteOne(discordChannel).
			Exec(m.ctx)
		if err != nil {
			log.Sugar.Errorf("Could not delete discord channel: %v", err)
		}
	} else {
		_, err = m.client.DiscordChannel.
			UpdateOne(discordChannel).
			RemoveUsers(entUser).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Errorf("Could not remove user from discord channel: %v", err)
		}
	}
	m.userManager.deleteIfUnused(entUser)
	return err
}

func (m *DiscordChannelManager) DeleteMultiple(channelIds []int64) {
	log.Sugar.Debugf("Delete %v discord channels", len(channelIds))

	for _, channelId := range channelIds {
		discordChannels, err := m.client.DiscordChannel.
			Query().
			Where(discordchannel.ChannelID(channelId)).
			WithUsers().
			All(m.ctx)
		if err != nil {
			log.Sugar.Errorf("Error while querying discord channels: %v", err)
		}
		for _, channel := range discordChannels {
			for _, u := range channel.Edges.Users {
				err := m.Delete(u.UserID, channelId)
				if err != nil {
					log.Sugar.Errorf("Error while deleting discord channels: %v", err)
				}
			}
		}
	}
}

func (m *DiscordChannelManager) GetChannelUsers(channelId int64) []*ent.User {
	users, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(channelId)).
		QueryUsers().
		All(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not get users for discord channel: %v", err)
	}
	return users
}

func (m *DiscordChannelManager) CountSubscriptions(channelId int64) int {
	count, err := m.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelIDEQ(channelId)).
		QueryContracts().
		Count(m.ctx)
	if err != nil {
		log.Sugar.Errorf("Could not count subscriptions for discord channel: %v", err)
	}
	return count
}

func (m *DiscordChannelManager) GetSubscribedIds(query *ent.DiscordChannelQuery) []types.DiscordChannelQueryResult {
	var v []types.DiscordChannelQueryResult
	err := query.
		Select(discordchannel.FieldChannelID, discordchannel.FieldName).
		Scan(m.ctx, &v)
	if err != nil {
		log.Sugar.Panicf("Could not get discord channels: %v", err)
	}
	return v
}

func (m *DiscordChannelManager) GetAllIds() []types.DiscordChannelQueryResult {
	return m.GetSubscribedIds(m.client.DiscordChannel.Query())
}
