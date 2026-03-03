package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
	"github.com/summit/summit-api/pkg/apperror"
	"github.com/summit/summit-api/pkg/pagination"
)

type OrderService struct {
	orderRepo    *repository.OrderRepository
	customerRepo *repository.CustomerRepository
	productRepo  *repository.ProductRepository
}

func NewOrderService(or *repository.OrderRepository, cr *repository.CustomerRepository, pr *repository.ProductRepository) *OrderService {
	return &OrderService{orderRepo: or, customerRepo: cr, productRepo: pr}
}

func (s *OrderService) GetByID(ctx context.Context, id int) (*models.Order, error) {
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	items, err := s.orderRepo.GetItems(ctx, id)
	if err != nil {
		return nil, err
	}
	order.Items = items
	return order, nil
}

func (s *OrderService) List(ctx context.Context, customerID *int, pg pagination.Params) ([]models.Order, int, error) {
	return s.orderRepo.List(ctx, customerID, pg)
}

// CreateOrder implements PRE-INSERT trigger logic:
// 1. Validate credit: if PaymentType == "CREDIT", check customer's credit_rating.
//    If not "GOOD" or "EXCELLENT", force PaymentType to "CASH".
// 2. Look up customer's sales rep and assign to order.
func (s *OrderService) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.Order, error) {
	// Credit validation (ported from Oracle Forms PL/SQL trigger)
	if req.PaymentType != nil && *req.PaymentType == "CREDIT" {
		customer, err := s.customerRepo.GetByID(ctx, req.CustomerID)
		if err != nil {
			return nil, err
		}

		if customer.CreditRating == nil || (*customer.CreditRating != "GOOD" && *customer.CreditRating != "EXCELLENT") {
			cash := "CASH"
			req.PaymentType = &cash
		}
	}

	// Auto-assign sales rep from customer if not provided
	if req.SalesRepID == nil {
		customer, err := s.customerRepo.GetByID(ctx, req.CustomerID)
		if err != nil {
			return nil, err
		}
		req.SalesRepID = customer.SalesRepID
	}

	return s.orderRepo.Create(ctx, req)
}

func (s *OrderService) UpdateOrder(ctx context.Context, id int, req models.UpdateOrderRequest) (*models.Order, error) {
	return s.orderRepo.Update(ctx, id, req)
}

// DeleteOrder implements ON-CHECK-DELETE-MASTER logic:
// Check if order has items. If it does, return an error.
func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	hasItems, err := s.orderRepo.HasItems(ctx, id)
	if err != nil {
		return err
	}
	if hasItems {
		return apperror.Conflict("cannot delete order with existing items; delete items first")
	}
	return s.orderRepo.Delete(ctx, id)
}

// Order Items

func (s *OrderService) GetItems(ctx context.Context, orderID int) ([]models.OrderItem, error) {
	return s.orderRepo.GetItems(ctx, orderID)
}

// AddItem looks up the product price and creates the order item.
func (s *OrderService) AddItem(ctx context.Context, orderID int, req models.CreateOrderItemRequest) (*models.OrderItem, error) {
	price, err := s.productRepo.GetPrice(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}
	return s.orderRepo.CreateItem(ctx, orderID, req, price)
}

func (s *OrderService) UpdateItem(ctx context.Context, orderID, itemID int, req models.UpdateOrderItemRequest) error {
	return s.orderRepo.UpdateItem(ctx, orderID, itemID, req)
}

func (s *OrderService) DeleteItem(ctx context.Context, orderID, itemID int) error {
	return s.orderRepo.DeleteItem(ctx, orderID, itemID)
}
