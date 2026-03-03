package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
)

type WarehouseRepository struct {
	pool *pgxpool.Pool
}

func NewWarehouseRepository(pool *pgxpool.Pool) *WarehouseRepository {
	return &WarehouseRepository{pool: pool}
}

func (r *WarehouseRepository) List(ctx context.Context) ([]models.Warehouse, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM warehouses ORDER BY city")
	if err != nil {
		return nil, fmt.Errorf("list warehouses: %w", err)
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var w models.Warehouse
		if err := rows.Scan(
			&w.ID, &w.RegionID, &w.Address, &w.City, &w.State,
			&w.Country, &w.ZipCode, &w.Phone, &w.ManagerID,
			&w.CreatedAt, &w.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan warehouse: %w", err)
		}
		warehouses = append(warehouses, w)
	}
	return warehouses, nil
}

func (r *WarehouseRepository) GetByID(ctx context.Context, id int) (*models.Warehouse, error) {
	var w models.Warehouse
	err := r.pool.QueryRow(ctx, "SELECT * FROM warehouses WHERE id = $1", id).Scan(
		&w.ID, &w.RegionID, &w.Address, &w.City, &w.State,
		&w.Country, &w.ZipCode, &w.Phone, &w.ManagerID,
		&w.CreatedAt, &w.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get warehouse: %w", err)
	}
	return &w, nil
}
