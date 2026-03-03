package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
)

type DepartmentService struct {
	departmentRepo *repository.DepartmentRepository
}

func NewDepartmentService(dr *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{departmentRepo: dr}
}

func (s *DepartmentService) List(ctx context.Context) ([]models.Department, error) {
	return s.departmentRepo.List(ctx)
}

func (s *DepartmentService) GetByID(ctx context.Context, id int) (*models.Department, error) {
	return s.departmentRepo.GetByID(ctx, id)
}
