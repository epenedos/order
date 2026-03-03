package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
)

type InventoryService struct {
	inventoryRepo *repository.InventoryRepository
}

func NewInventoryService(ir *repository.InventoryRepository) *InventoryService {
	return &InventoryService{inventoryRepo: ir}
}

func (s *InventoryService) GetByProduct(ctx context.Context, productID int) ([]models.Inventory, error) {
	return s.inventoryRepo.GetByProduct(ctx, productID)
}

func (s *InventoryService) List(ctx context.Context) ([]models.Inventory, error) {
	return s.inventoryRepo.List(ctx)
}
