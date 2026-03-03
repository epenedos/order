package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
	"github.com/summit/summit-api/pkg/pagination"
)

type CustomerService struct {
	customerRepo *repository.CustomerRepository
	employeeRepo *repository.EmployeeRepository
}

func NewCustomerService(cr *repository.CustomerRepository, er *repository.EmployeeRepository) *CustomerService {
	return &CustomerService{customerRepo: cr, employeeRepo: er}
}

func (s *CustomerService) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	return s.customerRepo.GetByID(ctx, id)
}

func (s *CustomerService) List(ctx context.Context, filter models.CustomerFilter, pg pagination.Params) ([]models.Customer, int, error) {
	return s.customerRepo.List(ctx, filter, pg)
}

func (s *CustomerService) Create(ctx context.Context, req models.CreateCustomerRequest) (*models.Customer, error) {
	return s.customerRepo.Create(ctx, req)
}

func (s *CustomerService) Update(ctx context.Context, id int, req models.UpdateCustomerRequest) (*models.Customer, error) {
	return s.customerRepo.Update(ctx, id, req)
}

func (s *CustomerService) Delete(ctx context.Context, id int) error {
	return s.customerRepo.Delete(ctx, id)
}

func (s *CustomerService) GetDistinctCountries(ctx context.Context) ([]string, error) {
	return s.customerRepo.GetDistinctCountries(ctx)
}

func (s *CustomerService) GetByCountry(ctx context.Context, country string) ([]models.Customer, error) {
	return s.customerRepo.GetByCountry(ctx, country)
}

// GetCustomerTree builds the hierarchical tree data for the customer navigator.
// mode: "by_country" groups customers under country nodes.
// mode: "by_sales_rep" groups customers under sales rep nodes.
func (s *CustomerService) GetCustomerTree(ctx context.Context, mode string) ([]models.TreeNode, error) {
	switch mode {
	case "by_country":
		return s.buildCountryTree(ctx)
	case "by_sales_rep":
		return s.buildSalesRepTree(ctx)
	default:
		return nil, fmt.Errorf("invalid tree mode: %s", mode)
	}
}

func (s *CustomerService) buildCountryTree(ctx context.Context) ([]models.TreeNode, error) {
	countries, err := s.customerRepo.GetDistinctCountries(ctx)
	if err != nil {
		return nil, err
	}

	var tree []models.TreeNode
	for _, country := range countries {
		customers, err := s.customerRepo.GetByCountry(ctx, country)
		if err != nil {
			return nil, err
		}

		var children []models.TreeNode
		for _, c := range customers {
			children = append(children, models.TreeNode{
				ID:    fmt.Sprintf("cust-%d", c.ID),
				Label: c.Name,
				Value: c.ID,
			})
		}

		tree = append(tree, models.TreeNode{
			ID:       fmt.Sprintf("country-%s", country),
			Label:    country,
			Value:    country,
			Children: children,
		})
	}
	return tree, nil
}

func (s *CustomerService) buildSalesRepTree(ctx context.Context) ([]models.TreeNode, error) {
	reps, err := s.employeeRepo.ListSalesReps(ctx)
	if err != nil {
		return nil, err
	}

	var tree []models.TreeNode
	for _, rep := range reps {
		customers, err := s.employeeRepo.GetCustomersBySalesRep(ctx, rep.ID)
		if err != nil {
			return nil, err
		}

		var children []models.TreeNode
		for _, c := range customers {
			children = append(children, models.TreeNode{
				ID:    fmt.Sprintf("cust-%d", c.ID),
				Label: c.Name,
				Value: c.ID,
			})
		}

		tree = append(tree, models.TreeNode{
			ID:       fmt.Sprintf("rep-%s", strconv.Itoa(rep.ID)),
			Label:    rep.FullName,
			Value:    rep.ID,
			Children: children,
		})
	}
	return tree, nil
}
