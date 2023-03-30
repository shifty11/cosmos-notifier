package database

import (
	"context"
	"errors"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/ent/validator"
	"github.com/shifty11/cosmos-notifier/log"
)

type ValidatorManager struct {
	client                *ent.Client
	ctx                   context.Context
	addressTrackerManager *AddressTrackerManager
}

func NewValidatorManager(client *ent.Client, ctx context.Context, addressTrackerManager *AddressTrackerManager) *ValidatorManager {
	return &ValidatorManager{client: client, ctx: ctx, addressTrackerManager: addressTrackerManager}
}

func (manager *ValidatorManager) AddValidator(
	chainEnt *ent.Chain,
	address string,
	moniker string,
) (*ent.Validator, error) {
	return manager.client.Validator.
		Create().
		SetChain(chainEnt).
		SetAddress(address).
		SetMoniker(moniker).
		Save(manager.ctx)
}

func (manager *ValidatorManager) UpdateValidator(validatorEnt *ent.Validator, moniker string) error {
	return validatorEnt.
		Update().
		SetMoniker(moniker).
		Exec(manager.ctx)
}

func (manager *ValidatorManager) DeleteValidator(validatorEnt *ent.Validator) error {
	return manager.client.Validator.
		DeleteOne(validatorEnt).
		Exec(manager.ctx)
}

func (manager *ValidatorManager) GetActive() []*ent.Validator {
	// TODO: filter active
	return manager.client.Validator.
		Query().
		WithChain().
		AllX(manager.ctx)
}

func (manager *ValidatorManager) GetByMoniker(moniker string) []*ent.Validator {
	return manager.client.Validator.
		Query().
		Where(validator.Moniker(moniker)).
		WithChain().
		AllX(manager.ctx)
}

func (manager *ValidatorManager) GetForUser(userEnt *ent.User) ([]*ent.Validator, error) {
	if userEnt.Type == user.TypeTelegram {
		return userEnt.
			QueryTelegramChats().
			//Where(telegramchat.IDEQ(telegramChatId)).
			QueryValidators().
			WithAddressTrackers().
			All(manager.ctx)
	} else {
		return userEnt.
			QueryDiscordChannels().
			//Where(discordchannel.IDEQ(discordChannelId)).
			QueryValidators().
			WithAddressTrackers().
			All(manager.ctx)
	}
}

// TrackValidator tracks a validator for a user.
// It adds the validator to the user's list of tracked validators (via Discord/Telegram relation)
// and creates a new AddressTracker if it doesn't exist.
func (manager *ValidatorManager) TrackValidator(
	userEnt *ent.User,
	validatorEnt *ent.Validator,
	discordChannelId int,
	telegramChatId int,
	notificationInterval int64,
) (*ent.AddressTracker, error) {
	if telegramChatId == 0 && discordChannelId == 0 {
		return nil, errors.New("a telegram chat or a discord channel must be provided")
	}
	if telegramChatId != 0 && discordChannelId != 0 {
		return nil, errors.New("only a telegram chat or a discord channel can be provided")
	}
	updateQuery := manager.client.Validator.UpdateOne(validatorEnt)
	if telegramChatId != 0 {
		updateQuery = updateQuery.AddTelegramChatIDs(telegramChatId)
	} else {
		updateQuery = updateQuery.AddDiscordChannelIDs(discordChannelId)
	}
	if !manager.addressTrackerManager.Exists(discordChannelId, telegramChatId, validatorEnt.Address) {
		tracker, err := withTxGeneric(manager.client, manager.ctx, func(tx *ent.Tx) (*ent.AddressTracker, error) {
			tracker, err := manager.addressTrackerManager.AddTracker(
				userEnt,
				validatorEnt.Address,
				discordChannelId,
				telegramChatId,
				notificationInterval,
			)
			if err != nil {
				return nil, err
			}
			return tracker, updateQuery.
				AddAddressTrackers(tracker).
				Exec(manager.ctx)
		})
		return tracker, err
	} else {
		trackers, err := manager.addressTrackerManager.
			AllByChatRoomsAndAddress(discordChannelId, telegramChatId, validatorEnt.Address)
		if err != nil {
			return nil, err
		}
		if len(trackers) == 0 {
			log.Sugar.Errorf("no address tracker found for address %s", validatorEnt.Address) // should never happen
			return nil, errors.New("no address tracker found")
		}
		return trackers[0], updateQuery.
			AddAddressTrackers(trackers[0]).
			Exec(manager.ctx)
	}
}

func (manager *ValidatorManager) UntrackValidator(userEnt *ent.User, validatorEnt *ent.Validator) ([]int, error) {
	toBeDeletedIds, err := validatorEnt.
		QueryAddressTrackers().
		Where(
			addresstracker.Or(
				addresstracker.HasDiscordChannelWith(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))),
				addresstracker.HasTelegramChatWith(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))),
			),
		).
		IDs(manager.ctx)
	if err != nil {
		return nil, err
	}
	if len(toBeDeletedIds) != 0 {
		_, err := manager.client.AddressTracker.
			Delete().
			Where(addresstracker.IDIn(toBeDeletedIds...)).
			Exec(manager.ctx)
		if err != nil {
			return nil, err
		}
	}

	if userEnt.Type == user.TypeTelegram {
		ids, err := validatorEnt.
			QueryTelegramChats().
			Where(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))).
			IDs(manager.ctx)
		if err != nil {
			return nil, err
		}
		if len(ids) != 0 {
			err := manager.client.Validator.
				UpdateOne(validatorEnt).
				RemoveTelegramChatIDs(ids...).
				Exec(manager.ctx)
			if err != nil {
				return nil, err
			}
		}
	} else {
		ids, err := validatorEnt.
			QueryDiscordChannels().
			Where(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))).
			IDs(manager.ctx)
		if err != nil {
			return nil, err
		}
		if len(ids) != 0 {
			err := manager.client.Validator.
				UpdateOne(validatorEnt).
				RemoveDiscordChannelIDs(ids...).
				Exec(manager.ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return toBeDeletedIds, nil
}
