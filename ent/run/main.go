package main

import (
	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"context"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	golangMigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	"github.com/lagzi/EntSymbolBug/ent"
	_ "github.com/lib/pq"
	"log"
)

func createDiff() {
	db, err := gorm.Open("postgres", "dbname=test host=localhost port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Failed opening DB: %v", err)
	}

	driver := entsql.OpenDB(db.Dialect().GetName(), db.DB())

	client := ent.NewClient(ent.Driver(driver))
	if client == nil {
		log.Fatalf("Error initiating ent client")
	}
	defer client.Close()

	ctx := context.Background()
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	err = client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithGlobalUniqueID(true), schema.WithFormatter(sqltool.GolangMigrateFormatter))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func main() {
	// Create 'test' DB before running

	createDiff()

	m, err := golangMigrate.New(
		"file://migrations",
		"postgres://postgres@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed golangMigrate.New: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("Failed running migrations: %v", err)
	}
}
