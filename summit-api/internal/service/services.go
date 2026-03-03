package service

import (
	"github.com/summit/summit-api/internal/config"
	"github.com/summit/summit-api/internal/repository"
)

type Services struct {
	Customer  *CustomerService
	Order     *OrderService
	Product   *ProductService
	Employee  *EmployeeService
	Inventory *InventoryService
	Auth      *AuthService
	Region    *RegionService
	Department *DepartmentService
	Warehouse *WarehouseService
}

func NewServices(repos *repository.Repositories, cfg *config.Config) *Services {
	return &Services{
		Customer:   NewCustomerService(repos.Customer, repos.Employee),
		Order:      NewOrderService(repos.Order, repos.Customer, repos.Product),
		Product:    NewProductService(repos.Product),
		Employee:   NewEmployeeService(repos.Employee),
		Inventory:  NewInventoryService(repos.Inventory),
		Auth:       NewAuthService(repos.User, cfg.JWTSecret),
		Region:     NewRegionService(repos.Region),
		Department: NewDepartmentService(repos.Department),
		Warehouse:  NewWarehouseService(repos.Warehouse),
	}
}
