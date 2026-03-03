package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
)

type EmployeeService struct {
	employeeRepo *repository.EmployeeRepository
}

func NewEmployeeService(er *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepo: er}
}

func (s *EmployeeService) GetByID(ctx context.Context, id int) (*models.Employee, error) {
	return s.employeeRepo.GetByID(ctx, id)
}

func (s *EmployeeService) List(ctx context.Context) ([]models.Employee, error) {
	return s.employeeRepo.List(ctx)
}

func (s *EmployeeService) ListSalesReps(ctx context.Context) ([]models.SalesRep, error) {
	return s.employeeRepo.ListSalesReps(ctx)
}

func (s *EmployeeService) GetCustomersBySalesRep(ctx context.Context, repID int) ([]models.Customer, error) {
	return s.employeeRepo.GetCustomersBySalesRep(ctx, repID)
}
