package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
)

type InventoryRepository struct {
	pool *pgxpool.Pool
}

func NewInventoryRepository(pool *pgxpool.Pool) *InventoryRepository {
	return &InventoryRepository{pool: pool}
}

func (r *InventoryRepository) GetByProduct(ctx context.Context, productID int) ([]models.Inventory, error) {
	query := `
		SELECT inv.product_id, inv.warehouse_id, inv.amount_in_stock, inv.reorder_point,
			inv.max_in_stock, inv.out_of_stock_explanation, inv.restock_date,
			inv.created_at, inv.updated_at, w.city AS warehouse_city
		FROM inventory inv
		LEFT JOIN warehouses w ON inv.warehouse_id = w.id
		WHERE inv.product_id = $1
		ORDER BY w.city`

	rows, err := r.pool.Query(ctx, query, productID)
	if err != nil {
		return nil, fmt.Errorf("get inventory by product: %w", err)
	}
	defer rows.Close()

	var items []models.Inventory
	for rows.Next() {
		var inv models.Inventory
		if err := rows.Scan(
			&inv.ProductID, &inv.WarehouseID, &inv.AmountInStock, &inv.ReorderPoint,
			&inv.MaxInStock, &inv.OutOfStockExplanation, &inv.RestockDate,
			&inv.CreatedAt, &inv.UpdatedAt, &inv.WarehouseCity,
		); err != nil {
			return nil, fmt.Errorf("scan inventory: %w", err)
		}
		items = append(items, inv)
	}
	return items, nil
}

func (r *InventoryRepository) List(ctx context.Context) ([]models.Inventory, error) {
	query := `
		SELECT inv.product_id, inv.warehouse_id, inv.amount_in_stock, inv.reorder_point,
			inv.max_in_stock, inv.out_of_stock_explanation, inv.restock_date,
			inv.created_at, inv.updated_at, w.city AS warehouse_city
		FROM inventory inv
		LEFT JOIN warehouses w ON inv.warehouse_id = w.id
		ORDER BY inv.product_id, w.city`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list inventory: %w", err)
	}
	defer rows.Close()

	var items []models.Inventory
	for rows.Next() {
		var inv models.Inventory
		if err := rows.Scan(
			&inv.ProductID, &inv.WarehouseID, &inv.AmountInStock, &inv.ReorderPoint,
			&inv.MaxInStock, &inv.OutOfStockExplanation, &inv.RestockDate,
			&inv.CreatedAt, &inv.UpdatedAt, &inv.WarehouseCity,
		); err != nil {
			return nil, fmt.Errorf("scan inventory: %w", err)
		}
		items = append(items, inv)
	}
	return items, nil
}
