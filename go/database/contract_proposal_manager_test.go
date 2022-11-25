package database

import (
	"context"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
	"time"
)

func newTestProposalManager(t *testing.T) *ContractProposalManager {
	manager := NewContractProposalManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func TestProposalManager_CreateOrUpdate(t *testing.T) {
	m := newTestProposalManager(t)
	cm := newTestContractManager(t)

	c, _ := cm.Create(&types.ContractData{
		Address:     "0x123",
		Name:        "name",
		Description: "description",
		ImageUrl:    "https://image.com",
	})

	atTime, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}

	propData := &types.Proposal{
		Id:       1,
		Proposer: "0x123",
		Expires: types.Expires{
			AtTime: atTime,
		},
		Title:       "title",
		Description: "description",
		Status:      types.StatusOpen,
	}

	prop, status := m.CreateOrUpdate(c, propData)
	if status != ProposalCreated {
		t.Errorf("Expected status %s, got %s", ProposalCreated, status)
	}
	if prop.ProposalID != propData.Id {
		t.Fatalf("Expected %d, got %d", propData.Id, prop.ProposalID)
	}
	if prop.Title != propData.Title {
		t.Fatalf("Expected %s, got %s", propData.Title, prop.Title)
	}
	if prop.Description != propData.Description {
		t.Fatalf("Expected %s, got %s", propData.Description, prop.Description)
	}
	if prop.Status.String() != string(propData.Status) {
		t.Fatalf("Expected %s, got %s", string(propData.Status), prop.Status.String())
	}
	if prop.ExpiresAt != propData.Expires.AtTime {
		t.Fatalf("Expected %s, got %s", propData.Expires.AtTime, prop.ExpiresAt)
	}
	qc := prop.QueryContract().OnlyX(m.ctx)
	if qc.Address != c.Address {
		t.Fatalf("Expected %s, got %s", c.Address, qc.Address)
	}

	propData.Title = "new title"
	prop, status = m.CreateOrUpdate(c, propData)
	if status != ProposalUpdated {
		t.Errorf("Expected status %s, got %s", ProposalUpdated, status)
	}
	if prop.Title != propData.Title {
		t.Fatalf("Expected %s, got %s", propData.Title, prop.Title)
	}

	propData.Description = "new description"
	prop, status = m.CreateOrUpdate(c, propData)
	if status != ProposalUpdated {
		t.Errorf("Expected status %s, got %s", ProposalUpdated, status)
	}
	if prop.Description != propData.Description {
		t.Fatalf("Expected %s, got %s", propData.Description, prop.Description)
	}

	propData.Status = types.StatusPassed
	prop, status = m.CreateOrUpdate(c, propData)
	if status != ProposalStatusChangedFromOpen {
		t.Errorf("Expected status %s, got %s", ProposalStatusChangedFromOpen, status)
	}
	if prop.Status.String() != string(propData.Status) {
		t.Fatalf("Expected %s, got %s", string(propData.Status), prop.Status.String())
	}

	for _, statType := range []types.ProposalStatus{types.StatusRejected, types.StatusExecuted, types.StatusPassed, types.StatusClosed, types.StatusExecutionFailed} {
		propData.Status = statType
		prop, status = m.CreateOrUpdate(c, propData)
		if status != ProposalUpdated {
			t.Errorf("Expected status %s, got %s", ProposalUpdated, status)
		}
		if prop.Status.String() != string(propData.Status) {
			t.Fatalf("Expected %s, got %s", string(propData.Status), prop.Status.String())
		}
	}

}
