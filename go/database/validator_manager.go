package database

import (
	"context"
	"errors"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/ent/validator"
	"github.com/shifty11/cosmos-notifier/log"
	"time"
)

const timeUntilConsideredInactive = 24 * time.Hour // TODO: set this to 1 month or so

type ValidatorManager struct {
	client                *ent.Client
	ctx                   context.Context
	addressTrackerManager *AddressTrackerManager
}

func NewValidatorManager(client *ent.Client, ctx context.Context, addressTrackerManager *AddressTrackerManager) *ValidatorManager {
	return &ValidatorManager{client: client, ctx: ctx, addressTrackerManager: addressTrackerManager}
}

func (manager *ValidatorManager) StartTx(ctx context.Context) (context.Context, error) {
	return startTx(ctx, manager.client)
}

func (manager *ValidatorManager) getAccountAddress(operatorAddress string, chainEnt *ent.Chain) (string, error) {
	_, valAddr, err := bech32.DecodeAndConvert(operatorAddress)
	if err != nil {
		return "", err
	}
	accAddr, err := cosmossdk.Bech32ifyAddressBytes(chainEnt.Bech32Prefix, valAddr)
	if err != nil {
		return "", err
	}
	return accAddr, nil
}

func (manager *ValidatorManager) getNillableFirstInactiveTime(isActive bool) *time.Time {
	var firstInactiveTime *time.Time
	if !isActive {
		var now = time.Now()
		firstInactiveTime = &now
	}
	return firstInactiveTime
}

func (manager *ValidatorManager) Create(
	chainEnt *ent.Chain,
	operatorAddress string,
	moniker string,
	isActive bool,
) (*ent.Validator, error) {
	accountAddress, err := manager.getAccountAddress(operatorAddress, chainEnt)
	if err != nil {
		log.Sugar.Errorf("Error while getting account address for validator %v: %v", operatorAddress, err)
		return nil, err
	}
	return manager.client.Validator.
		Create().
		SetChain(chainEnt).
		SetOperatorAddress(operatorAddress).
		SetAddress(accountAddress).
		SetMoniker(moniker).
		SetNillableFirstInactiveTime(manager.getNillableFirstInactiveTime(isActive)).
		Save(manager.ctx)
}

func (manager *ValidatorManager) Update(validatorEnt *ent.Validator, moniker string, isActive bool) error {
	return validatorEnt.
		Update().
		SetMoniker(moniker).
		ClearFirstInactiveTime(). // Clear field and set it on the next line if isActive is false
		SetNillableFirstInactiveTime(manager.getNillableFirstInactiveTime(isActive)).
		Exec(manager.ctx)
}

func (manager *ValidatorManager) Delete(validatorEnt *ent.Validator) error {
	return manager.client.Validator.
		DeleteOne(validatorEnt).
		Exec(manager.ctx)
}

func (manager *ValidatorManager) QueryActive() []*ent.Validator {
	return manager.client.Validator.
		Query().
		Where(validator.Or(
			validator.FirstInactiveTimeIsNil(),
			validator.FirstInactiveTimeGT(time.Now().Add(-timeUntilConsideredInactive)),
		)).
		WithChain().
		AllX(manager.ctx)
}

func (manager *ValidatorManager) QueryByMoniker(moniker string) []*ent.Validator {
	return manager.client.Validator.
		Query().
		Where(validator.Moniker(moniker)).
		WithChain().
		AllX(manager.ctx)
}

func (manager *ValidatorManager) QueryByUser(userEnt *ent.User) ([]*ent.Validator, error) {
	if userEnt.Type == user.TypeTelegram {
		return userEnt.
			QueryTelegramChats().
			QueryValidators().
			WithAddressTrackers().
			All(manager.ctx)
	} else {
		return userEnt.
			QueryDiscordChannels().
			QueryValidators().
			WithAddressTrackers().
			All(manager.ctx)
	}
}

// UpdateTrackValidator tracks a validator for a user.
// It adds the validator to the user's list of tracked validators (via Discord/Telegram relation)
// and creates a new AddressTracker if it doesn't exist.
func (manager *ValidatorManager) UpdateTrackValidator(
	ctx context.Context,
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
	client := getClient(ctx, manager.client)
	updateQuery := client.Validator.UpdateOne(validatorEnt)
	if telegramChatId != 0 {
		updateQuery = updateQuery.AddTelegramChatIDs(telegramChatId)
	} else {
		updateQuery = updateQuery.AddDiscordChannelIDs(discordChannelId)
	}
	if !manager.addressTrackerManager.QueryDoesExist(discordChannelId, telegramChatId, validatorEnt.Address) {
		tracker, err := manager.addressTrackerManager.Create(
			ctx,
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
			Exec(ctx)
	} else {
		trackers, err := manager.addressTrackerManager.
			QueryByChatRoomsAndAddress(discordChannelId, telegramChatId, validatorEnt.Address)
		if err != nil {
			return nil, err
		}
		if len(trackers) == 0 {
			log.Sugar.Errorf("no address tracker found for address %s", validatorEnt.Address) // should never happen
			return nil, errors.New("no address tracker found")
		}
		return trackers[0], updateQuery.
			AddAddressTrackers(trackers[0]).
			Exec(ctx)
	}
}

func (manager *ValidatorManager) UpdateUntrackValidator(ctx context.Context, userEnt *ent.User, validatorEnt *ent.Validator) ([]int, error) {
	client := getClient(ctx, manager.client)
	toBeDeletedIds, err := validatorEnt.
		QueryAddressTrackers().
		Where(
			addresstracker.Or(
				addresstracker.HasDiscordChannelWith(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))),
				addresstracker.HasTelegramChatWith(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))),
			),
		).
		IDs(ctx)
	if err != nil {
		return nil, err
	}
	if len(toBeDeletedIds) != 0 {
		_, err := client.AddressTracker.
			Delete().
			Where(addresstracker.IDIn(toBeDeletedIds...)).
			Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	if userEnt.Type == user.TypeTelegram {
		ids, err := validatorEnt.
			QueryTelegramChats().
			Where(telegramchat.HasUsersWith(user.IDEQ(userEnt.ID))).
			IDs(ctx)
		if err != nil {
			return nil, err
		}
		if len(ids) != 0 {
			err := client.Validator.
				UpdateOne(validatorEnt).
				RemoveTelegramChatIDs(ids...).
				Exec(ctx)
			if err != nil {
				return nil, err
			}
		}
	} else {
		ids, err := validatorEnt.
			QueryDiscordChannels().
			Where(discordchannel.HasUsersWith(user.IDEQ(userEnt.ID))).
			IDs(ctx)
		if err != nil {
			return nil, err
		}
		if len(ids) != 0 {
			err := client.Validator.
				UpdateOne(validatorEnt).
				RemoveDiscordChannelIDs(ids...).
				Exec(ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return toBeDeletedIds, nil
}
