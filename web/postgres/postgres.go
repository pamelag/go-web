package postgres

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/pamelag/go-web/web/rde"
)

type projectRepository struct {
	pool *pgx.ConnPool
}

// NewProjectRepository creates a new instance of projectRepository
func NewProjectRepository(pool *pgx.ConnPool) *projectRepository {
	r := &projectRepository{
		pool: pool,
	}

	return r
}

func (r *projectRepository) Store(project *rde.Project) error {
	pool := r.pool
	tx, err := pool.BeginEx(context.Background(), &pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	sql := "insert into project (project_uuid, project_name, project_status) values ($1, $2, $3)"
	stmt, err := tx.Prepare("insert_project", sql)
	if err != nil {
		return err
	}

	_, err = tx.Exec(stmt.Name, string(project.ProjectID), project.Name, string(project.ProjectStatus))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *projectRepository) StoreFeature(feature *rde.Feature) error {
	pool := r.pool
	tx, err := pool.BeginEx(context.Background(), &pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	sql := "insert into feature (feature_uuid, project_uuid, feature_title) values ($1, $2, $3)"
	stmt, err := tx.Prepare("insert_feature", sql)
	if err != nil {
		return err
	}

	_, err = tx.Exec(stmt.Name, string(feature.FeatureID), string(feature.ProjectID), feature.Title)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *projectRepository) StoreWireframe(wireframe *rde.Wireframe) error {
	pool := r.pool
	tx, err := pool.BeginEx(context.Background(), &pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	sql := "insert into wireframe (wireframe_uuid, feature_uuid, project_uuid, wireframe_title) values ($1, $2, $3, $4)"
	stmt, err := tx.Prepare("insert_wireframe", sql)
	if err != nil {
		return err
	}

	_, err = tx.Exec(stmt.Name, string(wireframe.WireframeID), string(wireframe.FeatureID), string(wireframe.ProjectID), wireframe.Title)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *projectRepository) UpdateWireframe(wireframe *rde.Wireframe) error {
	pool := r.pool
	tx, err := pool.BeginEx(context.Background(), &pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	sql := "update wireframe set wireframe_title = $1 where wireframe_uuid = $2 and feature_uuid = $3 and project_uuid = $4"
	stmt, err := tx.Prepare("update_wireframe", sql)
	if err != nil {
		return err
	}

	_, err = tx.Exec(stmt.Name, wireframe.Title, string(wireframe.WireframeID), string(wireframe.FeatureID), string(wireframe.ProjectID))
	if err != nil {
		return err
	}

	return tx.Commit()
}
