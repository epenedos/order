# Summit Sporting Goods - Migration Project

Order management system migrated from Oracle Forms 11g to a modern stack:
Go backend, Next.js (TypeScript) frontend, PostgreSQL database.

## Repository Structure

```
/                           # Root: legacy Oracle Forms files + new project
├── summit-api/             # Go backend (REST API)
│   ├── cmd/server/         # Entry point (main.go)
│   ├── internal/
│   │   ├── config/         # Environment-based configuration
│   │   ├── database/       # pgx connection pool + migrations/
│   │   ├── handler/        # HTTP handlers (one per domain)
│   │   ├── middleware/     # Auth (JWT), CORS, logging, panic recovery
│   │   ├── models/         # Domain structs + request/response types
│   │   ├── repository/     # Data access layer (SQL queries via pgx)
│   │   ├── router/         # Chi router with route definitions
│   │   └── service/        # Business logic layer
│   ├── pkg/                # Shared packages (apperror, pagination, validator)
│   ├── static/images/      # Converted product images
│   ├── Dockerfile          # Multi-stage production build
│   ├── Makefile            # Build, test, lint, migrate targets
│   └── go.mod              # Go 1.22, module: github.com/summit/summit-api
├── summit-web/             # TypeScript frontend (Next.js 14)
│   ├── src/
│   │   ├── app/            # App Router pages (customers, orders, products, etc.)
│   │   ├── components/     # React components (layout/, shared/, customers/, orders/, products/)
│   │   ├── hooks/          # TanStack Query hooks (useCustomers, useOrders, etc.)
│   │   ├── lib/            # API client, TypeScript types, Zod validators, utils
│   │   ├── providers/      # AuthProvider, QueryProvider
│   │   └── styles/         # Tailwind globals
│   ├── Dockerfile          # Multi-stage production build
│   ├── package.json        # Next.js 14, React 18, TanStack Query v5
│   └── tsconfig.json       # Strict mode, path alias @/*
├── docker-compose.yml      # PostgreSQL 16 + API + Web
├── *.fmb                   # Legacy Oracle Forms binaries (read-only reference)
├── *.pll                   # Legacy PL/SQL libraries (read-only reference)
└── summit.dmp              # Oracle database dump (source for data migration)
```

## Tech Stack

### Backend (`summit-api/`)
- **Language:** Go 1.22
- **Router:** chi v5 (`github.com/go-chi/chi/v5`)
- **Database:** PostgreSQL 16 via pgx v5 (`github.com/jackc/pgx/v5`) with connection pooling
- **Auth:** JWT (`github.com/golang-jwt/jwt/v5`) + bcrypt (`golang.org/x/crypto/bcrypt`)
- **Validation:** go-playground/validator v10 (struct tag validation)
- **Logging:** stdlib `log/slog` (structured JSON)
- **Migrations:** SQL files in `internal/database/migrations/`, applied via golang-migrate

### Frontend (`summit-web/`)
- **Framework:** Next.js 14 (App Router, `"use client"` components)
- **Language:** TypeScript 5.5 (strict mode)
- **State:** TanStack Query v5 (server state), React Hook Form + Zod (forms)
- **UI:** Tailwind CSS 3.4 + Radix UI primitives (shadcn/ui pattern)
- **Tables:** TanStack Table v8
- **Icons:** Lucide React
- **Dates:** date-fns

### Infrastructure
- **Database:** PostgreSQL 16 (Alpine)
- **Containers:** Docker Compose with health checks
- **Ports:** DB=5432, API=8080, Web=3000

## Quick Start

```bash
# Start everything
docker compose up --build

# Or run services individually for development:

# Terminal 1: Database
docker compose up db

# Terminal 2: Go API
cd summit-api
DATABASE_URL="postgres://summit:summit_dev@localhost:5432/summit?sslmode=disable" \
JWT_SECRET="dev-secret" \
go run ./cmd/server

# Terminal 3: Frontend
cd summit-web
npm install
npm run dev
```

## Common Commands

### Backend (run from `summit-api/`)
```bash
make build          # Compile to bin/server
make run            # go run ./cmd/server
make test           # go test ./...
make lint           # golangci-lint run ./...
make migrate-up     # Apply migrations (requires DATABASE_URL env)
make migrate-down   # Rollback one migration
```

### Frontend (run from `summit-web/`)
```bash
npm run dev         # Next.js dev server (port 3000)
npm run build       # Production build
npm run lint        # ESLint
npm run type-check  # tsc --noEmit
```

### Docker
```bash
docker compose up --build    # Build and start all services
docker compose up db         # Database only
docker compose down          # Stop all
docker compose down -v       # Stop all and delete data volume
```

## Environment Variables

| Variable | Service | Default | Description |
|---|---|---|---|
| `DATABASE_URL` | API | `postgres://summit:summit_dev@localhost:5432/summit?sslmode=disable` | PostgreSQL connection string |
| `JWT_SECRET` | API | *(required)* | Secret key for JWT signing |
| `PORT` | API | `8080` | API server port |
| `NEXT_PUBLIC_API_URL` | Web | `http://localhost:8080` | Backend API base URL |

## Architecture Patterns

### Backend: 4-Layer Architecture

```
HTTP Request → Handler → Service → Repository → PostgreSQL
```

- **Handler** (`internal/handler/`): Parse HTTP requests, validate input via `pkg/validator`, call service, write JSON response via `writeJSON`/`writeError` helpers in `response.go`
- **Service** (`internal/service/`): Business logic. Contains ported Oracle Forms trigger logic (credit validation, delete guards, tree hierarchy building, auto-total calculation)
- **Repository** (`internal/repository/`): Raw SQL queries via pgx. Each repo owns its table's queries. Returns domain models
- **Models** (`internal/models/`): Domain structs with `json`, `db`, and `validate` struct tags. Request/response types colocated with their entity

### Backend: Dependency Wiring

All dependencies are wired in `cmd/server/main.go`:
```
Config → Pool → Repositories → Services → Handlers → Router → http.Server
```

Aggregates: `repository.Repositories`, `service.Services`, `handler.Handlers` (each in their own `*s.go` file) hold all instances.

### Backend: Error Handling

Use `pkg/apperror` to return typed HTTP errors:
```go
return apperror.NotFound("customer not found")
return apperror.BadRequest(err.Error())
return apperror.Conflict("cannot delete order with existing items")
```

The `writeError` helper in `handler/response.go` maps `*apperror.AppError` to the correct HTTP status code. Untyped errors become 500.

### Backend: Request Validation

Use `pkg/validator.DecodeAndValidate(r, &req)` which combines JSON decoding + struct tag validation in one call. Validation rules use `validate:"required,max=50"` style tags on request structs.

### Frontend: Data Fetching Pattern

All API calls flow through TanStack Query hooks in `src/hooks/`:
```
Page Component → useCustomers() hook → api.get() → fetch() → Go API
```

- **Queries:** `useQuery` with structured query keys like `["customers", params]`
- **Mutations:** `useMutation` with automatic cache invalidation via `queryClient.invalidateQueries`
- **API Client:** `src/lib/api.ts` - singleton `ApiClient` class with auto-attached JWT from localStorage

### Frontend: Component Organization

- `src/components/shared/` — Reusable: `DataTable`, `TreeView`, `ConfirmDialog`, `LoadingSpinner`
- `src/components/layout/` — Shell: `Sidebar`, `Header`, `Toolbar`
- `src/components/customers/`, `orders/`, `products/` — Domain-specific (to be expanded)
- `src/app/` pages use `"use client"` directive and compose shared components

### Frontend: Type Safety

TypeScript interfaces in `src/lib/types.ts` mirror Go models exactly. Zod schemas in `src/lib/validators.ts` mirror Go validation tags. Keep these in sync when modifying models.

## Database Schema

13 tables in PostgreSQL, migrated from Oracle `S_*` tables. Schema in `summit-api/internal/database/migrations/001_initial_schema.sql`.

Key tables and relationships:
```
regions ← departments ← employees (self-ref: manager_id)
                              ↑
customers ←───────── sales_rep_id
    ↑
orders ←── customer_id
    ↑
order_items ←── product_id → products → images, long_texts
                                ↑
                           inventory ← warehouse_id → warehouses

users → employee_id (auth layer, new table)
```

Custom types: `credit_rating_type` (EXCELLENT/GOOD/POOR), `payment_type` (CASH/CREDIT), `user_role` (admin/sales_rep/viewer).

All tables have `created_at`/`updated_at` with auto-update triggers.

## Critical Business Rules (Ported from Oracle Forms)

These rules were extracted from PL/SQL triggers in the `.fmb` files and must be preserved:

1. **Credit validation** (`order_service.go:CreateOrder`): When `payment_type=CREDIT`, look up customer's `credit_rating`. If not GOOD or EXCELLENT, force `payment_type` to CASH
2. **Order delete guard** (`order_service.go:DeleteOrder`): Cannot delete an order that has line items. Must delete items first
3. **Auto-assign sales rep** (`order_service.go:CreateOrder`): If no sales_rep_id provided, inherit from customer's assigned sales rep
4. **Auto-calculate total** (`order_repo.go:UpdateTotal`): Order total = SUM(price * quantity) across all order_items. Recalculated on item add/update/delete
5. **Auto-assign item ID** (`order_repo.go:CreateItem`): item_id is MAX(item_id)+1 within the order, not a global sequence
6. **Product price lookup** (`order_service.go:AddItem`): When adding an order item, price is auto-populated from `products.suggested_whlsl_price`
7. **Customer tree modes** (`customer_service.go:GetCustomerTree`): Two hierarchy modes — `by_country` groups customers under country nodes, `by_sales_rep` groups under sales rep nodes

## API Routes

All routes prefixed with `/api/v1/`. Auth routes are public; all others require `Authorization: Bearer <jwt>`.

```
Auth:       POST /auth/login, /auth/register
Customers:  GET/POST /customers, GET/PUT/DELETE /customers/:id
            GET /customers/countries, /customers/tree?mode=, /customers/by-country/:country
Employees:  GET /employees, /employees/sales-reps, /employees/:id, /employees/:id/customers
Orders:     GET/POST /orders, GET/PUT/DELETE /orders/:id
            GET/POST /orders/:id/items, PUT/DELETE /orders/:id/items/:itemId
Products:   GET /products?search=, GET /products/:id, GET /products/:id/stock
Inventory:  GET /inventory
Departments: GET /departments, GET /departments/:id
Regions:    GET /regions
Warehouses: GET /warehouses
Health:     GET /health
```

## Adding a New Entity

### Backend
1. Add model struct + request types in `internal/models/<entity>.go`
2. Add repository in `internal/repository/<entity>_repo.go`, wire into `repositories.go`
3. Add service in `internal/service/<entity>_service.go`, wire into `services.go`
4. Add handler in `internal/handler/<entity>_handler.go`, wire into `handlers.go`
5. Add routes in `internal/router/router.go`
6. Add migration SQL in `internal/database/migrations/`

### Frontend
1. Add TypeScript interface in `src/lib/types.ts`
2. Add Zod schema in `src/lib/validators.ts` (if forms are needed)
3. Add TanStack Query hooks in `src/hooks/use<Entity>.ts`
4. Add page in `src/app/<entity>/page.tsx`
5. Add components in `src/components/<entity>/`

## Conventions

### Go
- Use `*Type` (pointer) for optional/nullable fields in models
- Use `COALESCE` in UPDATE queries for partial updates
- Repository methods return `(*Model, error)` for single, `([]Model, int, error)` for lists (int = total count)
- Service methods contain business logic; repos are pure data access
- Handler methods always call `writeJSON` or `writeError` — never write responses manually
- URL path params via `chi.URLParam(r, "id")`, query params via `r.URL.Query().Get("key")`

### TypeScript
- All page components use `"use client"` directive
- Nullable API fields use `type | null`, optional joined fields use `type?`
- TanStack Query keys are structured arrays: `["entity", id, "sub-resource"]`
- Use `cn()` from `lib/utils.ts` for conditional Tailwind classes
- Date formatting via `formatDate()`, currency via `formatCurrency()` from `lib/utils.ts`

### SQL
- Table names: lowercase plural (`customers`, `order_items`)
- Column names: snake_case
- Foreign keys: `<entity>_id` (e.g., `customer_id`, `sales_rep_id`)
- All tables have `created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()` and `updated_at` with trigger
- Indexes on all foreign key columns and common filter columns

## Legacy Files (Read-Only Reference)

The root directory contains the original Oracle Forms application. These files are kept for reference during migration but should not be modified:

- `*.fmb` — Oracle Forms binary modules (customers, orders, pick, reg, dept)
- `*.pll` — PL/SQL libraries (calendar, picklist, wizard, d2kdlstr)
- `summit.dmp` — Oracle database dump (13 MB, exp format V10.02.01)
- `Readme.html` — Original workload test scenarios (data entry + query flows)
- `*.tif` — Product images (23 files, need conversion to WebP/PNG)
- `icons/` — Legacy toolbar icons (.ico format)
- `web/` — Legacy web assets (.gif, .jpg)
- `RoundedButton.jar` — Custom Java bean for Oracle Forms
