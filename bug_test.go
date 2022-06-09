package bug

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/lagzi/EntSymbolBug/ent"
	"github.com/lagzi/EntSymbolBug/ent/enttest"
)

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"12": 5432} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(
				t,
				dialect.Postgres,
				fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port),
				enttest.WithMigrateOptions(
					schema.WithAtlas(true),
					schema.WithGlobalUniqueID(true),
				),
			)
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()

	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	fc := countFiles(t, "migrations")

	err = client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if fc != countFiles(t, "migrations") {
		t.Error("there is a drift")
	}
}

func countFiles(t *testing.T, path string) int {
	files, err := os.ReadDir(path)
	if err != nil {
		t.Error(err)
	}
	return len(files)
}
