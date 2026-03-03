package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
)

type WarehouseService struct {
	warehouseRepo *repository.WarehouseRepository
}

func NewWarehouseService(wr *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{warehouseRepo: wr}
}

func (s *WarehouseService) List(ctx context.Context) ([]models.Warehouse, error) {
	return s.warehouseRepo.List(ctx)
}
