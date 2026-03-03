package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	Customer   *CustomerRepository
	Order      *OrderRepository
	Product    *ProductRepository
	Employee   *EmployeeRepository
	Inventory  *InventoryRepository
	Warehouse  *WarehouseRepository
	Department *DepartmentRepository
	Region     *RegionRepository
	User       *UserRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		Customer:   NewCustomerRepository(pool),
		Order:      NewOrderRepository(pool),
		Product:    NewProductRepository(pool),
		Employee:   NewEmployeeRepository(pool),
		Inventory:  NewInventoryRepository(pool),
		Warehouse:  NewWarehouseRepository(pool),
		Department: NewDepartmentRepository(pool),
		Region:     NewRegionRepository(pool),
		User:       NewUserRepository(pool),
	}
}
