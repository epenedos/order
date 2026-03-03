package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
	"github.com/summit/summit-api/pkg/pagination"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(pr *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: pr}
}

func (s *ProductService) GetByID(ctx context.Context, id int) (*models.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}

func (s *ProductService) List(ctx context.Context, search string, pg pagination.Params) ([]models.Product, int, error) {
	return s.productRepo.List(ctx, search, pg)
}
