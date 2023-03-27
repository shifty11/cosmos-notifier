package database

import (
	"context"
	"errors"
	"github.com/shifty11/cosmos-notifier/ent"
)

type ValidatorManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewValidatorManager(client *ent.Client, ctx context.Context) *ValidatorManager {
	return &ValidatorManager{client: client, ctx: ctx}
}

func (manager *ValidatorManager) AddValidator(
	chainEnt *ent.Chain,
	address string,
	moniker string,
) (*ent.Validator, error) {
	if address == "" {
		return nil, errors.New("address is empty")
	}
	if moniker == "" {
		return nil, errors.New("moniker is empty")
	}

	return manager.client.Validator.
		Create().
		SetChain(chainEnt).
		SetAddress(address).
		SetMoniker(moniker).
		Save(manager.ctx)
}

func (manager *ValidatorManager) UpdateValidator(validatorEnt *ent.Validator, moniker string) error {
	if moniker == "" {
		return nil
	}

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
