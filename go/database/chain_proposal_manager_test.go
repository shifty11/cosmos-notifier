package database

import (
	"context"
	"fmt"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/types"
	"testing"
	"time"
)

func newTestChainProposalManager(t *testing.T) *ChainProposalManager {
	manager := NewChainProposalManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func addChainProposals(m *ChainProposalManager, chains []*ent.Chain) {
	stati := []types.ChainProposalStatus{types.ChainProposalStatusPassed, types.ChainProposalStatusRejected, types.ChainProposalStatusFailed, types.ChainProposalStatusVotingPeriod}
	oneWeekAgo := time.Now().Add(-time.Hour * 24 * 7)
	twoWeeksAgo := time.Now().Add(-time.Hour * 24 * 14)
	threeWeeksAgo := time.Now().Add(-time.Hour * 24 * 21)
	oneWeekInFuture := time.Now().Add(time.Hour * 24 * 7)
	votingStartTimes := []time.Time{threeWeeksAgo, twoWeeksAgo, oneWeekAgo, oneWeekAgo}
	votingEndTimes := []time.Time{twoWeeksAgo, oneWeekAgo, time.Now(), oneWeekInFuture}
	for _, chainDto := range chains {
		for i := 1; i <= len(stati); i++ {
			m.CreateOrUpdate(chainDto, &types.ChainProposal{
				ProposalId: i,
				Content: types.ChainProposalContent{
					Title:       fmt.Sprintf("title %d", i),
					Description: fmt.Sprintf("description %d", i),
				},
				Status:          stati[i-1],
				VotingStartTime: votingStartTimes[i-1],
				VotingEndTime:   votingEndTimes[i-1],
			})
		}
	}
}

func TestChainProposalManager_CreateOrUpdate(t *testing.T) {
	m := newTestChainProposalManager(t)
	cm := newTestChainManager(t)

	data := &types.Chain{
		ChainId:     "chainId-1",
		Name:        "chain-1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image.com",
	}
	c := cm.Create(data, data.Image)

	atTime, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}

	propData := &types.ChainProposal{
		ProposalId: 1,
		Content: types.ChainProposalContent{
			Title:       "title",
			Description: "description",
		},
		Status:          types.ChainProposalStatusVotingPeriod,
		VotingStartTime: atTime,
		VotingEndTime:   atTime.Add(time.Hour * 24 * 7),
	}

	prop, status := m.CreateOrUpdate(c, propData)
	if status != ProposalCreated {
		t.Errorf("Expected status %s, got %s", ProposalCreated, status)
	}
	if prop.ProposalID != propData.ProposalId {
		t.Fatalf("Expected %d, got %d", propData.ProposalId, prop.ProposalID)
	}
	if prop.Title != propData.Content.Title {
		t.Fatalf("Expected %s, got %s", propData.Content.Title, prop.Title)
	}
	if prop.Description != propData.Content.Description {
		t.Fatalf("Expected %s, got %s", propData.Content.Description, prop.Description)
	}
	if prop.Status.String() != string(propData.Status) {
		t.Fatalf("Expected %s, got %s", string(propData.Status), prop.Status.String())
	}
	if prop.VotingStartTime != propData.VotingStartTime {
		t.Fatalf("Expected %s, got %s", propData.VotingStartTime, prop.VotingStartTime)
	}
	if prop.VotingEndTime != propData.VotingEndTime {
		t.Fatalf("Expected %s, got %s", propData.VotingEndTime, prop.VotingEndTime)
	}
	qc := prop.QueryChain().OnlyX(m.ctx)
	if qc.ChainID != c.ChainID {
		t.Fatalf("Expected %s, got %s", c.ChainID, qc.ChainID)
	}

	propData.Content.Title = "new title"
	prop, status = m.CreateOrUpdate(c, propData)
	if status != ProposalUpdated {
		t.Errorf("Expected status %s, got %s", ProposalUpdated, status)
	}
	if prop.Title != propData.Content.Title {
		t.Fatalf("Expected %s, got %s", propData.Content.Title, prop.Title)
	}

	propData.Content.Description = "new description"
	prop, status = m.CreateOrUpdate(c, propData)
	if status != ProposalUpdated {
		t.Errorf("Expected status %s, got %s", ProposalUpdated, status)
	}
	if prop.Description != propData.Content.Description {
		t.Fatalf("Expected %s, got %s", propData.Content.Description, prop.Description)
	}

	propData.Status = types.ChainProposalStatusPassed
	prop, status = m.CreateOrUpdate(c, propData)
	if status != ProposalStatusChangedFromOpen {
		t.Errorf("Expected status %s, got %s", ProposalStatusChangedFromOpen, status)
	}
	if prop.Status.String() != string(propData.Status) {
		t.Fatalf("Expected %s, got %s", string(propData.Status), prop.Status.String())
	}

	for _, statType := range []types.ChainProposalStatus{types.ChainProposalStatusRejected, types.ChainProposalStatusFailed, types.ChainProposalStatusPassed} {
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

func TestChainProposalManager_VotingPeriodExpired(t *testing.T) {
	m := newTestChainProposalManager(t)
	cm := newTestChainManager(t)

	data := &types.Chain{
		ChainId:     "chainId-1",
		Name:        "chain-1",
		PrettyName:  "Chain 1",
		NetworkType: "mainnet",
		Image:       "https://image.com",
	}
	c := cm.Create(data, data.Image)

	atTime, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}

	propData := &types.ChainProposal{
		ProposalId: 1,
		Content: types.ChainProposalContent{
			Title:       "title",
			Description: "description",
		},
		Status:          types.ChainProposalStatusVotingPeriod,
		VotingStartTime: atTime,
		VotingEndTime:   atTime.Add(time.Hour * 24 * 7),
	}

	prop1, status := m.CreateOrUpdate(c, propData)
	if status != ProposalCreated {
		t.Errorf("Expected status %s, got %s", ProposalCreated, status)
	}

	propData.ProposalId = 2
	propData.VotingEndTime = time.Now().Add(time.Hour)
	prop2, status := m.CreateOrUpdate(c, propData)
	if status != ProposalCreated {
		t.Errorf("Expected status %s, got %s", ProposalCreated, status)
	}

	props := m.VotingPeriodExpired(c)
	if len(props) != 1 {
		t.Errorf("Expected 1, got %d", len(props))
	}
	if props[0].ProposalID != prop1.ProposalID {
		t.Errorf("Expected %d, got %d", prop2.ProposalID, props[0].ProposalID)
	}

	m.client.ChainProposal.DeleteOneID(prop2.ID).ExecX(m.ctx)
	propData.VotingEndTime = time.Now().Add(-time.Hour)
	prop2, status = m.CreateOrUpdate(c, propData)
	if status != ProposalCreated {
		t.Errorf("Expected status %s, got %s", ProposalUpdated, status)
	}

	props = m.VotingPeriodExpired(c)
	if len(props) != 2 {
		t.Errorf("Expected 2, got %d", len(props))
	}
}
