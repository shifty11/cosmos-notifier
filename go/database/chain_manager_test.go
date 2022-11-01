package database

import (
	"context"
	"testing"
)

func newTestChainManager(t *testing.T) *ChainManager {
	manager := NewChainManager(testClient(t), context.Background())
	t.Cleanup(func() { closeTestClient(manager.client) })
	return manager
}
