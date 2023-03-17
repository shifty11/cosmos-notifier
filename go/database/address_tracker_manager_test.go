package database

import (
	"context"
	"testing"
)

func newTestAddressTrackerManager(t *testing.T) *AddressTrackerManager {
	manager := NewAddressTrackerManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}

func TestAddressTrackerManager_IsValid(t *testing.T) {
	manager := newTestAddressTrackerManager(t)

	chainManager := newTestChainManager(t)
	addChains(chainManager)

	if manager.IsValid("") {
		t.Error("Empty address is valid")
	}
	if manager.IsValid("juno1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02") {
		t.Error("Address of unknown chain is valid")
	}
	if manager.IsValid("cosmos1") {
		t.Error("Invalid address is valid")
	}
	for _, address := range []string{"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02", "osmo166y8reslaeuedyc6gd83m8r5p0pmdnvq0dggsq", "comdex1cx82d7pm4dgffy7a93rl6ul5g84vjgxkqfyp2m"} {
		if !manager.IsValid(address) {
			t.Error("Valid address is invalid")
		}
	}
}
