package store

import (
	"context"
	"database/sql"

	db "taskman/pkg/db/sqlite"

	"taskman/internal/types"
)

type ProjectStore struct {
	ctx      context.Context
	database *sql.DB
	queries  *db.Queries
}

func ProjectToSQLite(project types.Project) db.Project {
	return db.Project{
		Name: project.Name,
		Rate: project.Rate,
	}
}

func ProjectFromSQLite(project db.Project) (*types.Project, error) {
	return &types.Project{
		Name: project.Name,
		Rate: project.Rate,
	}, nil
}

func NewProjectStore(ctx context.Context, database *sql.DB) *ProjectStore {
	queries := db.New(database)
	return &ProjectStore{
		ctx:      ctx,
		database: database,
		queries:  queries,
	}
}

func (s *ProjectStore) GetProject(projectId string) (*types.Project, error) {
	project, err := s.queries.GetProject(s.ctx, &projectId)
	if err != nil {
		return nil, err
	}
	result, err := ProjectFromSQLite(project)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProjectStore) Add(project types.Project) (*types.Project, error) {
	data, err := s.queries.AddProject(s.ctx, project.Name)
	if err != nil {
		return nil, err
	}
	result, err := ProjectFromSQLite(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProjectStore) Modify(projectId *string, project types.Project) (*types.Project, error) {
	var params db.ModifyProjectParams
	if project.Name == nil || *projectId == *project.Name {
		params = db.ModifyProjectParams{
			Name:    projectId,
			Rate:    project.Rate,
			NewName: projectId,
		}
	} else {
		params = db.ModifyProjectParams{
			Name:    projectId,
			NewName: project.Name,
			Rate:    project.Rate,
		}
	}

	data, err := s.queries.ModifyProject(s.ctx, params)
	if err != nil {
		return nil, err
	}
	result, err := ProjectFromSQLite(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProjectStore) Delete(projectId string) error {
	err := s.queries.DeleteProject(s.ctx, &projectId)
	if err != nil {
		return err
	}
	return nil
}
