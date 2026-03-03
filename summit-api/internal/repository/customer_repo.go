package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/pkg/pagination"
)

type CustomerRepository struct {
	pool *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{pool: pool}
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	query := `
		SELECT c.*, e.last_name AS sales_rep_name
		FROM customers c
		LEFT JOIN employees e ON c.sales_rep_id = e.id
		WHERE c.id = $1`

	row := r.pool.QueryRow(ctx, query, id)
	return r.scanCustomer(row)
}

func (r *CustomerRepository) List(ctx context.Context, filter models.CustomerFilter, pg pagination.Params) ([]models.Customer, int, error) {
	where := "WHERE 1=1"
	args := []interface{}{}
	argIdx := 1

	if filter.Country != nil {
		where += fmt.Sprintf(" AND c.country = $%d", argIdx)
		args = append(args, *filter.Country)
		argIdx++
	}
	if filter.SalesRepID != nil {
		where += fmt.Sprintf(" AND c.sales_rep_id = $%d", argIdx)
		args = append(args, *filter.SalesRepID)
		argIdx++
	}

	orderBy := "c.name"
	switch filter.SortBy {
	case "phone":
		orderBy = "c.phone"
	case "sales_rep_id":
		orderBy = "c.sales_rep_id"
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM customers c %s", where)
	var total int
	if err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count customers: %w", err)
	}

	query := fmt.Sprintf(`
		SELECT c.*, e.last_name AS sales_rep_name
		FROM customers c
		LEFT JOIN employees e ON c.sales_rep_id = e.id
		%s ORDER BY %s LIMIT $%d OFFSET $%d`,
		where, orderBy, argIdx, argIdx+1)
	args = append(args, pg.Limit, pg.Offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list customers: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		c, err := r.scanCustomerFromRows(rows)
		if err != nil {
			return nil, 0, err
		}
		customers = append(customers, *c)
	}
	return customers, total, nil
}

func (r *CustomerRepository) Create(ctx context.Context, req models.CreateCustomerRequest) (*models.Customer, error) {
	query := `
		INSERT INTO customers (name, phone, address, city, state, country, zip_code, credit_rating, sales_rep_id, region_id, comments)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`

	var id int
	err := r.pool.QueryRow(ctx, query,
		req.Name, req.Phone, req.Address, req.City, req.State,
		req.Country, req.ZipCode, req.CreditRating, req.SalesRepID, req.RegionID, req.Comments,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create customer: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *CustomerRepository) Update(ctx context.Context, id int, req models.UpdateCustomerRequest) (*models.Customer, error) {
	query := `
		UPDATE customers SET
			name = COALESCE($2, name),
			phone = COALESCE($3, phone),
			address = COALESCE($4, address),
			city = COALESCE($5, city),
			state = COALESCE($6, state),
			country = COALESCE($7, country),
			zip_code = COALESCE($8, zip_code),
			credit_rating = COALESCE($9, credit_rating),
			sales_rep_id = COALESCE($10, sales_rep_id),
			region_id = COALESCE($11, region_id),
			comments = COALESCE($12, comments)
		WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, id,
		req.Name, req.Phone, req.Address, req.City, req.State,
		req.Country, req.ZipCode, req.CreditRating, req.SalesRepID, req.RegionID, req.Comments,
	)
	if err != nil {
		return nil, fmt.Errorf("update customer: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *CustomerRepository) Delete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("delete customer: %w", err)
	}
	return nil
}

func (r *CustomerRepository) GetDistinctCountries(ctx context.Context) ([]string, error) {
	rows, err := r.pool.Query(ctx, "SELECT DISTINCT country FROM customers WHERE country IS NOT NULL ORDER BY country")
	if err != nil {
		return nil, fmt.Errorf("get countries: %w", err)
	}
	defer rows.Close()

	var countries []string
	for rows.Next() {
		var country string
		if err := rows.Scan(&country); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func (r *CustomerRepository) GetByCountry(ctx context.Context, country string) ([]models.Customer, error) {
	query := `SELECT c.*, e.last_name AS sales_rep_name
		FROM customers c
		LEFT JOIN employees e ON c.sales_rep_id = e.id
		WHERE c.country = $1 ORDER BY c.name`

	rows, err := r.pool.Query(ctx, query, country)
	if err != nil {
		return nil, fmt.Errorf("get customers by country: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		c, err := r.scanCustomerFromRows(rows)
		if err != nil {
			return nil, err
		}
		customers = append(customers, *c)
	}
	return customers, nil
}

func (r *CustomerRepository) scanCustomer(row pgx.Row) (*models.Customer, error) {
	var c models.Customer
	err := row.Scan(
		&c.ID, &c.Name, &c.Phone, &c.Address, &c.City, &c.State,
		&c.Country, &c.ZipCode, &c.CreditRating, &c.SalesRepID, &c.RegionID,
		&c.Comments, &c.CreatedAt, &c.UpdatedAt, &c.SalesRepName,
	)
	if err != nil {
		return nil, fmt.Errorf("scan customer: %w", err)
	}
	return &c, nil
}

func (r *CustomerRepository) scanCustomerFromRows(rows pgx.Rows) (*models.Customer, error) {
	var c models.Customer
	err := rows.Scan(
		&c.ID, &c.Name, &c.Phone, &c.Address, &c.City, &c.State,
		&c.Country, &c.ZipCode, &c.CreditRating, &c.SalesRepID, &c.RegionID,
		&c.Comments, &c.CreatedAt, &c.UpdatedAt, &c.SalesRepName,
	)
	if err != nil {
		return nil, fmt.Errorf("scan customer: %w", err)
	}
	return &c, nil
}
