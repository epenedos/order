package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
)

type DepartmentRepository struct {
	pool *pgxpool.Pool
}

func NewDepartmentRepository(pool *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{pool: pool}
}

func (r *DepartmentRepository) List(ctx context.Context) ([]models.Department, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM departments ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("list departments: %w", err)
	}
	defer rows.Close()

	var depts []models.Department
	for rows.Next() {
		var d models.Department
		if err := rows.Scan(&d.ID, &d.Name, &d.RegionID, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan department: %w", err)
		}
		depts = append(depts, d)
	}
	return depts, nil
}

func (r *DepartmentRepository) GetByID(ctx context.Context, id int) (*models.Department, error) {
	var d models.Department
	err := r.pool.QueryRow(ctx, "SELECT * FROM departments WHERE id = $1", id).Scan(
		&d.ID, &d.Name, &d.RegionID, &d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get department: %w", err)
	}
	return &d, nil
}
