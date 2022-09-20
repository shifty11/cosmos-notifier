package database

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/enttest"
	"testing"
)

func testClient(t *testing.T) *ent.Client {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	filename := fmt.Sprintf("file:ent%v?mode=memory&cache=shared&_fk=1", t.Name())
	client := enttest.Open(t, "sqlite3", filename, opts...)
	return client
}

func closeTestClient(client *ent.Client) {
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()
}
