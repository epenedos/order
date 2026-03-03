package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	err := r.pool.QueryRow(ctx,
		"SELECT id, email, password_hash, employee_id, role, is_active, created_at, updated_at FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.EmployeeID, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return &u, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User
	err := r.pool.QueryRow(ctx,
		"SELECT id, email, password_hash, employee_id, role, is_active, created_at, updated_at FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.EmployeeID, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return &u, nil
}

func (r *UserRepository) Create(ctx context.Context, email, passwordHash string, employeeID *int) (*models.User, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		"INSERT INTO users (email, password_hash, employee_id) VALUES ($1, $2, $3) RETURNING id",
		email, passwordHash, employeeID,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return r.GetByID(ctx, id)
}
