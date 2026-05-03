package testutil

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func runMigrations(t *testing.T, dsn string) {
	t.Helper()

	cmd := exec.Command("goose", "up")

	cmd.Env = append(os.Environ(),
		"DATABASE_URL="+dsn,
	)

	cmd.Dir = filepath.Join("internal", "adapters", "database", "migrations")

	done := make(chan error, 1)

	go func() {
		_, err := cmd.CombinedOutput()
		done <- err
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("migrations failed: %v", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("migrations timeout")
	}
}
