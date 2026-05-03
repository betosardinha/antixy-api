package testutil

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type TestDB struct {
	DB  *sql.DB
	DSN string
}

func SetupTestDB(t *testing.T) *TestDB {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	pg, err := postgres.Run(ctx,
		"postgres:16",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
	)
	if err != nil {
		t.Fatal(err)
	}

	dsn, err := pg.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Logf("db close error: %v", err)
		}

		if err := pg.Terminate(ctx); err != nil {
			t.Logf("container terminate error: %v", err)
		}
	})

	runMigrations(t, dsn)

	return &TestDB{
		DB:  db,
		DSN: dsn,
	}
}
