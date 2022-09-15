package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/types"
	"testing"
)

func newTestContractManager(t *testing.T) *ContractManager {
	return NewContractManager(testClient(t), context.Background())
}

func TestContractManager_All(t *testing.T) {
	m := newTestContractManager(t)
	contracts := m.All()
	if len(contracts) != 0 {
		t.Errorf("Expected 0 contracts, got %d", len(contracts))
	}

	name := "test"
	description := "description"
	address := "0x123"
	imageUrl := "https://image.com"
	thumbnailUrl := "https://test.com"

	m.client.Contract.
		Create().
		SetName(name).
		SetDescription(description).
		SetAddress(address).
		SetImageURL(imageUrl).
		SetThumbnailURL(thumbnailUrl).
		SaveX(m.ctx)

	contracts = m.All()
	if len(contracts) != 1 {
		t.Errorf("Expected 1 contract, got %d", len(contracts))
	}
	if contracts[0].Name != name {
		t.Errorf("Expected name %s, got %s", name, contracts[0].Name)
	}
	if contracts[0].Description != description {
		t.Errorf("Expected description %s, got %s", description, contracts[0].Description)
	}
	if contracts[0].Address != address {
		t.Errorf("Expected address %s, got %s", address, contracts[0].Address)
	}
	if contracts[0].ImageURL != imageUrl {
		t.Errorf("Expected image url %s, got %s", imageUrl, contracts[0].ImageURL)
	}
	if contracts[0].ThumbnailURL != thumbnailUrl {
		t.Errorf("Expected thumbnail url %s, got %s", thumbnailUrl, contracts[0].ThumbnailURL)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer m.client.Close()
}

func TestContractManager_Get(t *testing.T) {
	m := newTestContractManager(t)
	_, err := m.Get(1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	name := "test"
	description := "description"
	address := "0x123"
	imageUrl := "https://image.com"
	thumbnailUrl := "https://test.com"

	contract := m.client.Contract.
		Create().
		SetName(name).
		SetDescription(description).
		SetAddress(address).
		SetImageURL(imageUrl).
		SetThumbnailURL(thumbnailUrl).
		SaveX(m.ctx)

	c, err := m.Get(contract.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if c.Name != name {
		t.Errorf("Expected name %s, got %s", name, c.Name)
	}
	if c.Description != description {
		t.Errorf("Expected description %s, got %s", description, c.Description)
	}
	if c.Address != address {
		t.Errorf("Expected address %s, got %s", address, c.Address)
	}
	if c.ImageURL != imageUrl {
		t.Errorf("Expected image url %s, got %s", imageUrl, c.ImageURL)
	}
	if c.ThumbnailURL != thumbnailUrl {
		t.Errorf("Expected thumbnail url %s, got %s", thumbnailUrl, c.ThumbnailURL)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer m.client.Close()
}

func TestContractManager_CreateOrUpdate(t *testing.T) {
	m := newTestContractManager(t)
	name := "test"
	description := "description"
	address := "0x123"
	imageUrl := "https://image.com"

	data := &types.ContractData{
		Address:     address,
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
	}

	contract, status := m.CreateOrUpdate(data)
	if status != ContractCreated {
		t.Errorf("Expected status %s, got %s", ContractCreated, status)
	}
	if contract.Name != name {
		t.Errorf("Expected name %s, got %s", name, contract.Name)
	}
	if contract.Description != description {
		t.Errorf("Expected description %s, got %s", description, contract.Description)
	}
	if contract.Address != address {
		t.Errorf("Expected address %s, got %s", address, contract.Address)
	}
	if contract.ImageURL != imageUrl {
		t.Errorf("Expected image url %s, got %s", imageUrl, contract.ImageURL)
	}
	if contract.ThumbnailURL != "" {
		t.Errorf("Expected thumbnail url %s, got %s", "", contract.ThumbnailURL)
	}

	updatedTime := contract.UpdateTime
	contract, status = m.CreateOrUpdate(data)
	if status != ContractNoChanges {
		t.Errorf("Expected status %s, got %s", ContractNoChanges, status)
	}
	if contract.UpdateTime.Unix() != updatedTime.Unix() {
		t.Errorf("Expected update time %v, got %v", updatedTime.Unix(), contract.UpdateTime.Unix())
	}

	data.Name = "updated"
	contract, status = m.CreateOrUpdate(data)
	if status != ContractUpdated {
		t.Errorf("Expected status %s, got %s", ContractUpdated, status)
	}
	if contract.Name != "updated" {
		t.Errorf("Expected name %s, got %s", "updated", contract.Name)
	}

	data.Description = "updated"
	contract, status = m.CreateOrUpdate(data)
	if status != ContractUpdated {
		t.Errorf("Expected status %s, got %s", ContractUpdated, status)
	}
	if contract.Description != "updated" {
		t.Errorf("Expected description %s, got %s", "updated", contract.Description)
	}

	data.ImageUrl = "https://updated.com"
	contract, status = m.CreateOrUpdate(data)
	if status != ContractImageChanged {
		t.Errorf("Expected status %s, got %s", ContractImageChanged, status)
	}
	if contract.ImageURL != "https://updated.com" {
		t.Errorf("Expected image url %s, got %s", "https://updated.com", contract.ImageURL)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer m.client.Close()
}

func TestContractManager_SaveThumbnailUrl(t *testing.T) {
	m := newTestContractManager(t)
	name := "test"
	description := "description"
	address := "0x123"
	imageUrl := "https://image.com"
	thumbnailUrl := "https://test.com"

	contract := m.client.Contract.
		Create().
		SetName(name).
		SetDescription(description).
		SetAddress(address).
		SetImageURL(imageUrl).
		SetThumbnailURL(thumbnailUrl).
		SaveX(m.ctx)

	contract = m.SaveThumbnailUrl(contract, "https://updated.com")
	if contract.ThumbnailURL != "https://updated.com" {
		t.Errorf("Expected thumbnail url %s, got %s", "https://updated.com", contract.ThumbnailURL)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer m.client.Close()
}
