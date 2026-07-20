package store

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"taskman/internal/types"
	db "taskman/pkg/db/sqlite"
)

func TestProjectStoreModifyProject(t *testing.T) {
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open database: %v", err)
	}
	defer database.Close()

	if _, err := database.Exec(`CREATE TABLE projects (name TEXT PRIMARY KEY, rate REAL)`); err != nil {
		t.Fatalf("create projects table: %v", err)
	}

	ctx := context.Background()
	store := NewProjectStore(ctx, database)

	projectName := "alpha"
	projectRate := 12.5
	if _, err := store.Add(types.Project{Name: &projectName, Rate: &projectRate}); err != nil {
		t.Fatalf("add project: %v", err)
	}

	updatedName := "beta"
	updatedRate := 20.0
	updated, err := store.Modify(&projectName, types.Project{Name: &updatedName, Rate: &updatedRate})
	if err != nil {
		t.Fatalf("modify project: %v", err)
	}
	if updated == nil || updated.Name == nil || *updated.Name != updatedName {
		t.Fatalf("expected renamed project, got %+v", updated)
	}

	got, err := db.New(database).GetProject(ctx, &updatedName)
	if err != nil {
		t.Fatalf("get updated project: %v", err)
	}
	if got.Name == nil || *got.Name != updatedName {
		t.Fatalf("expected stored project name %q, got %+v", updatedName, got)
	}
}
