package bug

import (
	"context"
	"testing"
	"text/template"

	"ariga.io/atlas/sql/migrate"
	schema2 "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"github.com/lagzi/EntSymbolBug/ent"
)

func Test(t *testing.T) {
	ctx := context.Background()

	db, err := sql.Open(dialect.Postgres, "host=localhost port=5432 user=postgres dbname=test password=nopass sslmode=disable")
	require.NoError(t, err)
	client := ent.NewClient(ent.Driver(db))

	// Run the "old" migration engine.
	require.NoError(t, client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)))

	// Create 2 records for the second allocation (user).
	_, err = client.User.Create().SetName("user 1").Save(ctx)
	require.NoError(t, err)
	_, err = client.User.Create().SetName("user 1").Save(ctx)
	require.NoError(t, err)

	// Now create a migration dir and let atlas create a diff for it.
	// dir, err := migrate.NewLocalDir("migrations")
	// require.NoError(t, err)
	require.NoError(t, client.Schema.Create(ctx, schema.WithGlobalUniqueID(true), schema.WithAtlas(true), schema.WithDiffHook(func(differ schema.Differ) schema.Differ {
		return schema.DiffFunc(func(current, desired *schema2.Schema) ([]schema2.Change, error) {
			changes, err := differ.Diff(current, desired)
			return changes, err
		})
	})))

	// There should be one, containing the ALTER ... RESTART statement.
	// require.FileExists(t, "migrations/changes.sql")

	// Apply that file.
	// d, err := os.ReadFile("migrations/changes.sql")
	// require.NoError(t, err)
	// _, err = db.ExecContext(ctx, string(d))
	// require.NoError(t, err)

	// Now create new entities.
}

func simplerFormatter(t *testing.T) migrate.Formatter {
	f, err := migrate.NewTemplateFormatter(
		template.Must(template.New("").Parse("{{ .Name }}.sql")),
		template.Must(template.New("").Parse(
			`{{ range .Changes }}{{ printf "%s;\n" .Cmd }}{{ end }}`,
		)),
	)
	require.NoError(t, err)
	return f
}
