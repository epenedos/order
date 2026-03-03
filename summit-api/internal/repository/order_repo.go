package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/pkg/pagination"
)

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}

func (r *OrderRepository) GetByID(ctx context.Context, id int) (*models.Order, error) {
	query := `
		SELECT o.id, o.customer_id, o.date_ordered, o.date_shipped, o.sales_rep_id,
			o.total, o.payment_type, o.order_filled, o.created_at, o.updated_at,
			c.name AS customer_name, e.last_name AS sales_rep_name
		FROM orders o
		LEFT JOIN customers c ON o.customer_id = c.id
		LEFT JOIN employees e ON o.sales_rep_id = e.id
		WHERE o.id = $1`

	var o models.Order
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&o.ID, &o.CustomerID, &o.DateOrdered, &o.DateShipped, &o.SalesRepID,
		&o.Total, &o.PaymentType, &o.OrderFilled, &o.CreatedAt, &o.UpdatedAt,
		&o.CustomerName, &o.SalesRepName,
	)
	if err != nil {
		return nil, fmt.Errorf("get order: %w", err)
	}
	return &o, nil
}

func (r *OrderRepository) List(ctx context.Context, customerID *int, pg pagination.Params) ([]models.Order, int, error) {
	where := ""
	args := []interface{}{}
	argIdx := 1

	if customerID != nil {
		where = fmt.Sprintf("WHERE o.customer_id = $%d", argIdx)
		args = append(args, *customerID)
		argIdx++
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM orders o %s", where)
	var total int
	if err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count orders: %w", err)
	}

	query := fmt.Sprintf(`
		SELECT o.id, o.customer_id, o.date_ordered, o.date_shipped, o.sales_rep_id,
			o.total, o.payment_type, o.order_filled, o.created_at, o.updated_at,
			c.name AS customer_name, e.last_name AS sales_rep_name
		FROM orders o
		LEFT JOIN customers c ON o.customer_id = c.id
		LEFT JOIN employees e ON o.sales_rep_id = e.id
		%s ORDER BY o.date_ordered DESC LIMIT $%d OFFSET $%d`,
		where, argIdx, argIdx+1)
	args = append(args, pg.Limit, pg.Offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list orders: %w", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(
			&o.ID, &o.CustomerID, &o.DateOrdered, &o.DateShipped, &o.SalesRepID,
			&o.Total, &o.PaymentType, &o.OrderFilled, &o.CreatedAt, &o.UpdatedAt,
			&o.CustomerName, &o.SalesRepName,
		); err != nil {
			return nil, 0, fmt.Errorf("scan order: %w", err)
		}
		orders = append(orders, o)
	}
	return orders, total, nil
}

func (r *OrderRepository) Create(ctx context.Context, req models.CreateOrderRequest) (*models.Order, error) {
	query := `
		INSERT INTO orders (customer_id, date_ordered, sales_rep_id, payment_type)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	var id int
	err := r.pool.QueryRow(ctx, query,
		req.CustomerID, req.DateOrdered, req.SalesRepID, req.PaymentType,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}
	return r.GetByID(ctx, id)
}

func (r *OrderRepository) Update(ctx context.Context, id int, req models.UpdateOrderRequest) (*models.Order, error) {
	query := `
		UPDATE orders SET
			date_ordered = COALESCE($2, date_ordered),
			date_shipped = COALESCE($3, date_shipped),
			sales_rep_id = COALESCE($4, sales_rep_id),
			payment_type = COALESCE($5, payment_type),
			order_filled = COALESCE($6, order_filled)
		WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, id,
		req.DateOrdered, req.DateShipped, req.SalesRepID, req.PaymentType, req.OrderFilled,
	)
	if err != nil {
		return nil, fmt.Errorf("update order: %w", err)
	}
	return r.GetByID(ctx, id)
}

func (r *OrderRepository) Delete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("delete order: %w", err)
	}
	return nil
}

func (r *OrderRepository) UpdateTotal(ctx context.Context, orderID int) error {
	query := `
		UPDATE orders SET total = (
			SELECT COALESCE(SUM(price * quantity), 0)
			FROM order_items WHERE ord_id = $1
		) WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, orderID)
	if err != nil {
		return fmt.Errorf("update order total: %w", err)
	}
	return nil
}

func (r *OrderRepository) HasItems(ctx context.Context, orderID int) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM order_items WHERE ord_id = $1)", orderID).Scan(&exists)
	return exists, err
}

// Order Items

func (r *OrderRepository) GetItems(ctx context.Context, orderID int) ([]models.OrderItem, error) {
	query := `
		SELECT oi.ord_id, oi.item_id, oi.product_id, oi.price, oi.quantity,
			oi.quantity_shipped, oi.created_at, oi.updated_at,
			p.name AS product_name
		FROM order_items oi
		LEFT JOIN products p ON oi.product_id = p.id
		WHERE oi.ord_id = $1
		ORDER BY oi.item_id`

	rows, err := r.pool.Query(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("get order items: %w", err)
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(
			&item.OrdID, &item.ItemID, &item.ProductID, &item.Price,
			&item.Quantity, &item.QuantityShipped, &item.CreatedAt, &item.UpdatedAt,
			&item.ProductName,
		); err != nil {
			return nil, fmt.Errorf("scan order item: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *OrderRepository) CreateItem(ctx context.Context, orderID int, req models.CreateOrderItemRequest, price float64) (*models.OrderItem, error) {
	// Auto-assign next item_id within the order
	var nextItemID int
	err := r.pool.QueryRow(ctx,
		"SELECT COALESCE(MAX(item_id), 0) + 1 FROM order_items WHERE ord_id = $1",
		orderID,
	).Scan(&nextItemID)
	if err != nil {
		return nil, fmt.Errorf("get next item id: %w", err)
	}

	query := `
		INSERT INTO order_items (ord_id, item_id, product_id, price, quantity)
		VALUES ($1, $2, $3, $4, $5)`

	_, err = r.pool.Exec(ctx, query, orderID, nextItemID, req.ProductID, price, req.Quantity)
	if err != nil {
		return nil, fmt.Errorf("create order item: %w", err)
	}

	// Recalculate total
	if err := r.UpdateTotal(ctx, orderID); err != nil {
		return nil, err
	}

	item := &models.OrderItem{
		OrdID:     orderID,
		ItemID:    nextItemID,
		ProductID: req.ProductID,
		Price:     &price,
		Quantity:  req.Quantity,
	}
	return item, nil
}

func (r *OrderRepository) UpdateItem(ctx context.Context, orderID, itemID int, req models.UpdateOrderItemRequest) error {
	query := `
		UPDATE order_items SET
			quantity = COALESCE($3, quantity),
			quantity_shipped = COALESCE($4, quantity_shipped)
		WHERE ord_id = $1 AND item_id = $2`

	_, err := r.pool.Exec(ctx, query, orderID, itemID, req.Quantity, req.QuantityShipped)
	if err != nil {
		return fmt.Errorf("update order item: %w", err)
	}

	return r.UpdateTotal(ctx, orderID)
}

func (r *OrderRepository) DeleteItem(ctx context.Context, orderID, itemID int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM order_items WHERE ord_id = $1 AND item_id = $2", orderID, itemID)
	if err != nil {
		return fmt.Errorf("delete order item: %w", err)
	}
	return r.UpdateTotal(ctx, orderID)
}
