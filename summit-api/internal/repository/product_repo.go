package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/pkg/pagination"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}

func (r *ProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT p.id, p.name, p.short_desc, p.longtext_id, p.image_id,
			p.suggested_whlsl_price, p.whlsl_units, p.created_at, p.updated_at,
			i.filename AS image_url, lt.text_content AS description
		FROM products p
		LEFT JOIN images i ON p.image_id = i.id
		LEFT JOIN long_texts lt ON p.longtext_id = lt.id
		WHERE p.id = $1`

	var p models.Product
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.Name, &p.ShortDesc, &p.LongtextID, &p.ImageID,
		&p.SuggestedWhlslPrice, &p.WhlslUnits, &p.CreatedAt, &p.UpdatedAt,
		&p.ImageURL, &p.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("get product: %w", err)
	}
	return &p, nil
}

func (r *ProductRepository) List(ctx context.Context, search string, pg pagination.Params) ([]models.Product, int, error) {
	where := ""
	args := []interface{}{}
	argIdx := 1

	if search != "" {
		where = fmt.Sprintf("WHERE p.name ILIKE $%d", argIdx)
		args = append(args, "%"+search+"%")
		argIdx++
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products p %s", where)
	var total int
	if err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count products: %w", err)
	}

	query := fmt.Sprintf(`
		SELECT p.id, p.name, p.short_desc, p.longtext_id, p.image_id,
			p.suggested_whlsl_price, p.whlsl_units, p.created_at, p.updated_at,
			i.filename AS image_url, lt.text_content AS description
		FROM products p
		LEFT JOIN images i ON p.image_id = i.id
		LEFT JOIN long_texts lt ON p.longtext_id = lt.id
		%s ORDER BY p.name LIMIT $%d OFFSET $%d`,
		where, argIdx, argIdx+1)
	args = append(args, pg.Limit, pg.Offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID, &p.Name, &p.ShortDesc, &p.LongtextID, &p.ImageID,
			&p.SuggestedWhlslPrice, &p.WhlslUnits, &p.CreatedAt, &p.UpdatedAt,
			&p.ImageURL, &p.Description,
		); err != nil {
			return nil, 0, fmt.Errorf("scan product: %w", err)
		}
		products = append(products, p)
	}
	return products, total, nil
}

func (r *ProductRepository) GetPrice(ctx context.Context, productID int) (float64, error) {
	var price float64
	err := r.pool.QueryRow(ctx,
		"SELECT COALESCE(suggested_whlsl_price, 0) FROM products WHERE id = $1",
		productID,
	).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("get product price: %w", err)
	}
	return price, nil
}
