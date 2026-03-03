package handler

import "github.com/summit/summit-api/internal/service"

type Handlers struct {
	Customer  *CustomerHandler
	Order     *OrderHandler
	Product   *ProductHandler
	Employee  *EmployeeHandler
	Inventory *InventoryHandler
	Auth      *AuthHandler
	Department *DepartmentHandler
	Region    *RegionHandler
	Warehouse *WarehouseHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Customer:   NewCustomerHandler(services.Customer),
		Order:      NewOrderHandler(services.Order),
		Product:    NewProductHandler(services.Product),
		Employee:   NewEmployeeHandler(services.Employee),
		Inventory:  NewInventoryHandler(services.Inventory),
		Auth:       NewAuthHandler(services.Auth),
		Department: NewDepartmentHandler(services.Department),
		Region:     NewRegionHandler(services.Region),
		Warehouse:  NewWarehouseHandler(services.Warehouse),
	}
}
