//go:build go1.16
// +build go1.16

package iofs_test

import (
	"embed"
	"log"

	"github.com/KSerik/migrate"
	_ "github.com/KSerik/migrate/database/postgres"
	"github.com/KSerik/migrate/source/iofs"
)

//go:embed testdata/migrations/*.sql
var fs embed.FS

func Example() {
	d, err := iofs.New(fs, "testdata/migrations")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, "postgres://postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		// ...
	}
	// ...
}
