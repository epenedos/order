package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
)

type RegionRepository struct {
	pool *pgxpool.Pool
}

func NewRegionRepository(pool *pgxpool.Pool) *RegionRepository {
	return &RegionRepository{pool: pool}
}

func (r *RegionRepository) List(ctx context.Context) ([]models.Region, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM regions ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("list regions: %w", err)
	}
	defer rows.Close()

	var regions []models.Region
	for rows.Next() {
		var reg models.Region
		if err := rows.Scan(&reg.ID, &reg.Name, &reg.CreatedAt, &reg.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan region: %w", err)
		}
		regions = append(regions, reg)
	}
	return regions, nil
}
