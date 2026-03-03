package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
)

type EmployeeRepository struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{pool: pool}
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id int) (*models.Employee, error) {
	query := `SELECT * FROM employees WHERE id = $1`
	var e models.Employee
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&e.ID, &e.LastName, &e.FirstName, &e.UserID, &e.StartDate,
		&e.Comments, &e.ManagerID, &e.Title, &e.DeptID, &e.Salary,
		&e.CommissionPct, &e.CreatedAt, &e.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get employee: %w", err)
	}
	return &e, nil
}

func (r *EmployeeRepository) ListSalesReps(ctx context.Context) ([]models.SalesRep, error) {
	query := `
		SELECT id, last_name || ' ' || COALESCE(first_name, '') AS full_name, COALESCE(title, '') AS title
		FROM employees
		WHERE title LIKE 'Sales%'
		ORDER BY title, last_name`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list sales reps: %w", err)
	}
	defer rows.Close()

	var reps []models.SalesRep
	for rows.Next() {
		var rep models.SalesRep
		if err := rows.Scan(&rep.ID, &rep.FullName, &rep.Title); err != nil {
			return nil, fmt.Errorf("scan sales rep: %w", err)
		}
		reps = append(reps, rep)
	}
	return reps, nil
}

func (r *EmployeeRepository) GetCustomersBySalesRep(ctx context.Context, repID int) ([]models.Customer, error) {
	query := `
		SELECT c.id, c.name, c.phone, c.address, c.city, c.state, c.country,
			c.zip_code, c.credit_rating, c.sales_rep_id, c.region_id, c.comments,
			c.created_at, c.updated_at, e.last_name AS sales_rep_name
		FROM customers c
		LEFT JOIN employees e ON c.sales_rep_id = e.id
		WHERE c.sales_rep_id = $1
		ORDER BY c.name`

	rows, err := r.pool.Query(ctx, query, repID)
	if err != nil {
		return nil, fmt.Errorf("get customers by sales rep: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(
			&c.ID, &c.Name, &c.Phone, &c.Address, &c.City, &c.State, &c.Country,
			&c.ZipCode, &c.CreditRating, &c.SalesRepID, &c.RegionID, &c.Comments,
			&c.CreatedAt, &c.UpdatedAt, &c.SalesRepName,
		); err != nil {
			return nil, fmt.Errorf("scan customer: %w", err)
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (r *EmployeeRepository) List(ctx context.Context) ([]models.Employee, error) {
	query := `SELECT * FROM employees ORDER BY last_name`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list employees: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		if err := rows.Scan(
			&e.ID, &e.LastName, &e.FirstName, &e.UserID, &e.StartDate,
			&e.Comments, &e.ManagerID, &e.Title, &e.DeptID, &e.Salary,
			&e.CommissionPct, &e.CreatedAt, &e.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan employee: %w", err)
		}
		employees = append(employees, e)
	}
	return employees, nil
}
