package database

import (
	"context"
	"errors"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
	"time"
)

type AddressTrackerManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewAddressTrackerManager(client *ent.Client, ctx context.Context) *AddressTrackerManager {
	return &AddressTrackerManager{client: client, ctx: ctx}
}

func (manager *AddressTrackerManager) QueryByUser(userEnt *ent.User) ([]*ent.AddressTracker, error) {
	return manager.client.AddressTracker.
		Query().
		Where(
			addresstracker.Or(
				addresstracker.HasDiscordChannelWith(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))),
				addresstracker.HasTelegramChatWith(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))),
			),
		).
		WithDiscordChannel().
		WithTelegramChat().
		WithValidator().
		All(manager.ctx)
}

func (manager *AddressTrackerManager) QueryByChatRoomsAndAddress(
	discordChannelId int,
	telegramChatId int,
	address string,
) ([]*ent.AddressTracker, error) {
	if discordChannelId == 0 && telegramChatId == 0 {
		return nil, errors.New("at least one of discordChannelId or telegramChatId must be non-zero")
	}
	if discordChannelId != 0 && telegramChatId != 0 {
		return nil, errors.New("only one of discordChannelId or telegramChatId must be non-zero")
	}
	if discordChannelId != 0 {
		return manager.client.AddressTracker.
			Query().
			Where(addresstracker.And(
				addresstracker.HasDiscordChannelWith(discordchannel.IDEQ(discordChannelId)),
				addresstracker.AddressEQ(address),
			)).
			WithChain().
			WithDiscordChannel().
			WithValidator().
			All(manager.ctx)
	}
	if telegramChatId != 0 {
		return manager.client.AddressTracker.
			Query().
			Where(addresstracker.And(
				addresstracker.HasTelegramChatWith(telegramchat.IDEQ(telegramChatId)),
				addresstracker.AddressEQ(address),
			)).
			WithChain().
			WithTelegramChat().
			WithValidator().
			All(manager.ctx)
	}
	return []*ent.AddressTracker{}, nil
}

func (manager *AddressTrackerManager) QueryIsValid(address string) (bool, *ent.Chain) {
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

func (manager *AddressTrackerManager) Create(
	ctx context.Context,
	userEnt *ent.User,
	address string,
	discordChannelId int,
	telegramChatId int,
	notificationInterval int64,
) (*ent.AddressTracker, error) {
	isValid, chainEnt := manager.QueryIsValid(address)
	if !isValid {
		return nil, errors.New("invalid address")
	}
	if discordChannelId == 0 && telegramChatId == 0 {
		return nil, errors.New("at least one of discordChannelId or telegramChatId must be non-zero")
	}
	if discordChannelId != 0 && telegramChatId != 0 {
		return nil, errors.New("only one of discordChannelId or telegramChatId must be non-zero")
	}
	if notificationInterval < 0 {
		return nil, errors.New("notification interval must be non-negative")
	}
	client := getClient(ctx, manager.client)

	createQuery := client.AddressTracker.
		Create().
		SetChain(chainEnt).
		SetAddress(address).
		SetNotificationInterval(notificationInterval)

	if discordChannelId != 0 {
		exist, err := userEnt.QueryDiscordChannels().
			Where(discordchannel.IDEQ(discordChannelId)).
			Exist(ctx)
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
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.New("telegram chat not found")
		}
		createQuery.SetTelegramChatID(telegramChatId)
	}

	created, err := createQuery.Save(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddressTracker.Query().
		Where(addresstracker.IDEQ(created.ID)).
		WithChain().
		WithDiscordChannel().
		WithTelegramChat().
		Only(ctx)
}

func (manager *AddressTrackerManager) Update(
	userEnt *ent.User,
	addressTrackerId int,
	discordChannelId int,
	telegramChatId int,
	notificationInterval int64,
) (*ent.AddressTracker, error) {
	if discordChannelId == 0 && telegramChatId == 0 {
		return nil, errors.New("at least one of discordChannelId or telegramChatId must be non-zero")
	}
	if discordChannelId != 0 && telegramChatId != 0 {
		return nil, errors.New("only one of discordChannelId or telegramChatId must be non-zero")
	}
	if notificationInterval < 0 {
		return nil, errors.New("notification interval must be non-negative")
	}

	addressTracker, err := manager.client.AddressTracker.
		Query().
		Where(addresstracker.And(
			addresstracker.IDEQ(addressTrackerId),
			addresstracker.Or(
				addresstracker.HasDiscordChannelWith(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))),
				addresstracker.HasTelegramChatWith(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))),
			),
		)).
		Only(manager.ctx)
	if err != nil {
		return nil, err
	}
	updateQuery := addressTracker.
		Update().
		SetNotificationInterval(notificationInterval)

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
		updateQuery.SetDiscordChannelID(discordChannelId)
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
		updateQuery.SetTelegramChatID(telegramChatId)
	}

	updated, err := updateQuery.Save(manager.ctx)
	if err != nil {
		return nil, err
	}
	return manager.client.AddressTracker.Query().
		Where(addresstracker.IDEQ(updated.ID)).
		WithChain().
		WithDiscordChannel().
		WithTelegramChat().
		Only(manager.ctx)
}

func (manager *AddressTrackerManager) Delete(
	userEnt *ent.User,
	addressTrackerId int,
) error {
	addressTracker, err := manager.client.AddressTracker.
		Query().
		Where(addresstracker.And(
			addresstracker.IDEQ(addressTrackerId),
			addresstracker.Or(
				addresstracker.HasDiscordChannelWith(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))),
				addresstracker.HasTelegramChatWith(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))),
			),
		)).
		Only(manager.ctx)
	if err != nil {
		return err
	}
	return manager.client.AddressTracker.DeleteOne(addressTracker).Exec(manager.ctx)
}

type AddressTrackerWithChainProposal struct {
	AddressTracker *ent.AddressTracker
	ChainProposal  *ent.ChainProposal
}

func (manager *AddressTrackerManager) QueryUnnotifiedTrackers() []AddressTrackerWithChainProposal {
	proposals, err := manager.client.AddressTracker.
		Query().
		QueryChain().
		QueryChainProposals().
		Where(chainproposal.StatusEQ(chainproposal.StatusPROPOSAL_STATUS_VOTING_PERIOD)).
		WithChain().
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("error while getting all proposals: %v", err)
	}
	var result []AddressTrackerWithChainProposal
	for _, proposal := range proposals {
		timeUntilVotingEnd := proposal.VotingEndTime.Sub(time.Now())
		addressTrackers, err := manager.client.AddressTracker.
			Query().
			Where(
				addresstracker.And(
					addresstracker.HasChainWith(chain.IDEQ(proposal.Edges.Chain.ID)),
					addresstracker.Not(addresstracker.HasChainProposalsWith(chainproposal.IDEQ(proposal.ID))),
					addresstracker.NotificationIntervalGTE(int64(timeUntilVotingEnd.Seconds())),
				),
			).
			WithTelegramChat().
			WithDiscordChannel().
			All(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("error while getting address trackers: %v", err)
		}
		for _, addressTracker := range addressTrackers {
			result = append(result, AddressTrackerWithChainProposal{
				AddressTracker: addressTracker,
				ChainProposal:  proposal,
			})
		}
	}
	return result
}

func (manager *AddressTrackerManager) UpdateSetNotified(data AddressTrackerWithChainProposal) {
	err := data.AddressTracker.
		Update().
		AddChainProposals(data.ChainProposal).
		Exec(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("error while setting notified: %v", err)
	}
}

func (manager *AddressTrackerManager) QueryChatRooms(userEnt *ent.User) ([]*ent.DiscordChannel, []*ent.TelegramChat, error) {
	discordChannels, err := userEnt.
		QueryDiscordChannels().
		WithValidators().
		All(manager.ctx)
	if err != nil {
		return nil, nil, err
	}
	telegramChats, err := userEnt.
		QueryTelegramChats().
		WithValidators().
		All(manager.ctx)
	if err != nil {
		return nil, nil, err
	}
	if len(discordChannels) == 0 && len(telegramChats) == 0 {
		return nil, nil, errors.New("no chat rooms found")
	}
	if len(discordChannels) != 0 && len(telegramChats) != 0 {
		return nil, nil, errors.New("only one of discord channels or telegram chats must be non-zero")
	}
	return discordChannels, telegramChats, nil
}

func (manager *AddressTrackerManager) QueryDoesExist(discordChannelId int, telegramChatId int, validatorAddress string) bool {
	if discordChannelId != 0 {
		return manager.client.AddressTracker.
			Query().
			Where(addresstracker.And(
				addresstracker.HasDiscordChannelWith(discordchannel.IDEQ(discordChannelId)),
				addresstracker.AddressEQ(validatorAddress),
			)).
			ExistX(manager.ctx)
	}
	if telegramChatId != 0 {
		return manager.client.AddressTracker.
			Query().
			Where(addresstracker.And(
				addresstracker.HasTelegramChatWith(telegramchat.IDEQ(telegramChatId)),
				addresstracker.AddressEQ(validatorAddress),
			)).
			ExistX(manager.ctx)
	}
	return false
}
