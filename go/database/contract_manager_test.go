package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent/contractproposal"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/types"
	"testing"
	"time"
)

func newTestContractManager(t *testing.T) *ContractManager {
	manager := NewContractManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
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
}

func TestContractManager_Update(t *testing.T) {
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

	contract, err := m.Create(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	updatedTime := contract.UpdateTime
	contract = m.Update(contract, data)
	if contract.UpdateTime.Unix() != updatedTime.Unix() {
		t.Errorf("Expected update time %v, got %v", updatedTime.Unix(), contract.UpdateTime.Unix())
	}

	data.Name = "updated"
	contract = m.Update(contract, data)
	if contract.Name != "updated" {
		t.Errorf("Expected name %s, got %s", "updated", contract.Name)
	}

	data.Description = "updated"
	contract = m.Update(contract, data)
	if contract.Description != "updated" {
		t.Errorf("Expected description %s, got %s", "updated", contract.Description)
	}

	data.ImageUrl = "https://updated.com"
	contract = m.Update(contract, data)
	if contract.ImageURL != "https://updated.com" {
		t.Errorf("Expected image url %s, got %s", "https://updated.com", contract.ImageURL)
	}
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
}

func TestContractManager_ByAddress(t *testing.T) {
	m := newTestContractManager(t)

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

	c, err := m.ByAddress(address)
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
}

func TestContractManager_ByAddress_NotFound(t *testing.T) {
	m := newTestContractManager(t)

	_, err := m.ByAddress("0x123")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestContractManager_Delete(t *testing.T) {
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
	m.client.ContractProposal.
		Create().
		SetContract(contract).
		SetDescription("test").
		SetTitle("test").
		SetProposalID(1).
		SetExpiresAt(time.Now()).
		SetStatus(contractproposal.StatusOpen).
		SaveX(m.ctx)
	u := m.client.User.
		Create().
		SetUserID(1).
		SetName("test").
		SetType(user.TypeDiscord).
		SetRole(user.RoleUser).
		SaveX(m.ctx)
	m.client.DiscordChannel.
		Create().
		AddContracts(contract).
		AddUsers(u).
		SetName("test").
		SetChannelID(1).
		SetIsGroup(false).
		SaveX(m.ctx)

	dc := m.client.DiscordChannel.Query().WithContracts().FirstX(m.ctx)
	if len(dc.Edges.Contracts) != 1 {
		t.Errorf("Expected 1 contract, got %d", len(dc.Edges.Contracts))
	}

	err := m.Delete(contract.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	dc = m.client.DiscordChannel.Query().WithContracts().FirstX(m.ctx)
	if len(dc.Edges.Contracts) != 0 {
		t.Errorf("Expected no contracts, got %d", len(dc.Edges.Contracts))
	}

	cnt := m.client.ContractProposal.Query().CountX(m.ctx)
	if cnt != 0 {
		t.Errorf("Expected no proposals, got %d", cnt)
	}
}

func TestContractManager_Delete_NotFound(t *testing.T) {
	m := newTestContractManager(t)
	err := m.Delete(1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestContractManager_Create(t *testing.T) {
	m := newTestContractManager(t)
	data := &types.ContractData{
		Name:        "test",
		Description: "description",
		Address:     "0x123",
		ImageUrl:    "https://image.com",
	}
	contract, err := m.Create(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if contract.Name != data.Name {
		t.Errorf("Expected name %s, got %s", data.Name, contract.Name)
	}
	if contract.Description != data.Description {
		t.Errorf("Expected description %s, got %s", data.Description, contract.Description)
	}
	if contract.Address != data.Address {
		t.Errorf("Expected address %s, got %s", data.Address, contract.Address)
	}
	if contract.ImageURL != data.ImageUrl {
		t.Errorf("Expected image url %s, got %s", data.ImageUrl, contract.ImageURL)
	}

	_, err = m.Create(data)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
