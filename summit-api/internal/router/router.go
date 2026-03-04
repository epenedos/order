package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/summit/summit-api/internal/config"
	"github.com/summit/summit-api/internal/handler"
	"github.com/summit/summit-api/internal/middleware"
)

func New(h *handler.Handlers, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimw.RequestID)
	r.Use(middleware.Logging)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(middleware.CORS(cfg.CORSOrigins)))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Static files (product images)
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Route("/api/v1", func(r chi.Router) {
		// Public routes
		r.Post("/auth/login", h.Auth.Login)
		r.Post("/auth/register", h.Auth.Register)

		// Protected routes (auth temporarily disabled)
		r.Group(func(r chi.Router) {
			// TODO: re-enable when JWT is working
			// r.Use(middleware.Auth(h.Auth.Service()))

			// Customers
			r.Route("/customers", func(r chi.Router) {
				r.Get("/", h.Customer.List)
				r.Post("/", h.Customer.Create)
				r.Get("/countries", h.Customer.GetCountries)
				r.Get("/tree", h.Customer.GetTree)
				r.Get("/by-country/{country}", h.Customer.GetByCountry)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.Customer.GetByID)
					r.Put("/", h.Customer.Update)
					r.Delete("/", h.Customer.Delete)
				})
			})

			// Employees
			r.Route("/employees", func(r chi.Router) {
				r.Get("/", h.Employee.List)
				r.Get("/sales-reps", h.Employee.ListSalesReps)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.Employee.GetByID)
					r.Get("/customers", h.Employee.GetCustomers)
				})
			})

			// Orders
			r.Route("/orders", func(r chi.Router) {
				r.Get("/", h.Order.List)
				r.Post("/", h.Order.Create)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.Order.GetByID)
					r.Put("/", h.Order.Update)
					r.Delete("/", h.Order.Delete)
					r.Get("/items", h.Order.GetItems)
					r.Post("/items", h.Order.CreateItem)
					r.Put("/items/{itemId}", h.Order.UpdateItem)
					r.Delete("/items/{itemId}", h.Order.DeleteItem)
				})
			})

			// Products
			r.Route("/products", func(r chi.Router) {
				r.Get("/", h.Product.List)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", h.Product.GetByID)
					r.Get("/stock", h.Inventory.GetByProduct)
				})
			})

			// Inventory
			r.Get("/inventory", h.Inventory.List)

			// Departments
			r.Route("/departments", func(r chi.Router) {
				r.Get("/", h.Department.List)
				r.Get("/{id}", h.Department.GetByID)
			})

			// Regions
			r.Get("/regions", h.Region.List)

			// Warehouses
			r.Get("/warehouses", h.Warehouse.List)
		})
	})

	return r
}
