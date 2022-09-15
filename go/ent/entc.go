//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/hedwigz/entviz"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureVersionedMigration},
		},
		entc.Extensions(entviz.Extension{}))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
