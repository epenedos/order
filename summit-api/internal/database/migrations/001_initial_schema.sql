-- Summit Sporting Goods: PostgreSQL Schema
-- Migrated from Oracle Forms S_* tables

-- Enums for constrained values
CREATE TYPE credit_rating_type AS ENUM ('EXCELLENT', 'GOOD', 'POOR');
CREATE TYPE payment_type AS ENUM ('CASH', 'CREDIT');
CREATE TYPE user_role AS ENUM ('admin', 'sales_rep', 'viewer');

-- Regions
CREATE TABLE regions (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Job Titles (lookup)
CREATE TABLE titles (
    title       VARCHAR(25) PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Departments
CREATE TABLE departments (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(25) NOT NULL,
    region_id   INTEGER REFERENCES regions(id),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(name, region_id)
);

-- Employees (includes sales reps, managers)
CREATE TABLE employees (
    id              SERIAL PRIMARY KEY,
    last_name       VARCHAR(25) NOT NULL,
    first_name      VARCHAR(25),
    userid          VARCHAR(8) UNIQUE,
    start_date      TIMESTAMPTZ,
    comments        TEXT,
    manager_id      INTEGER REFERENCES employees(id),
    title           VARCHAR(25) REFERENCES titles(title),
    dept_id         INTEGER REFERENCES departments(id),
    salary          NUMERIC(11,2),
    commission_pct  NUMERIC(4,2),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_employees_manager ON employees(manager_id);
CREATE INDEX idx_employees_title ON employees(title);
CREATE INDEX idx_employees_dept ON employees(dept_id);

-- Customers
CREATE TABLE customers (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(50) NOT NULL,
    phone           VARCHAR(25),
    address         VARCHAR(400),
    city            VARCHAR(30),
    state           VARCHAR(20),
    country         VARCHAR(30),
    zip_code        VARCHAR(75),
    credit_rating   credit_rating_type,
    sales_rep_id    INTEGER REFERENCES employees(id),
    region_id       INTEGER REFERENCES regions(id),
    comments        TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_customers_sales_rep ON customers(sales_rep_id);
CREATE INDEX idx_customers_country ON customers(country);
CREATE INDEX idx_customers_name ON customers(name);
CREATE INDEX idx_customers_region ON customers(region_id);

-- Long text descriptions (for products)
CREATE TABLE long_texts (
    id              SERIAL PRIMARY KEY,
    use_filename    BOOLEAN DEFAULT FALSE,
    filename        VARCHAR(255),
    text_content    TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Product images
CREATE TABLE images (
    id              SERIAL PRIMARY KEY,
    format          VARCHAR(25),
    use_filename    BOOLEAN DEFAULT FALSE,
    filename        VARCHAR(255),
    image_data      BYTEA,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Products
CREATE TABLE products (
    id                      SERIAL PRIMARY KEY,
    name                    VARCHAR(50) NOT NULL UNIQUE,
    short_desc              VARCHAR(255),
    longtext_id             INTEGER REFERENCES long_texts(id),
    image_id                INTEGER REFERENCES images(id),
    suggested_whlsl_price   NUMERIC(11,2),
    whlsl_units             VARCHAR(25),
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Warehouses
CREATE TABLE warehouses (
    id          SERIAL PRIMARY KEY,
    region_id   INTEGER NOT NULL REFERENCES regions(id),
    address     TEXT,
    city        VARCHAR(30),
    state       VARCHAR(20),
    country     VARCHAR(30),
    zip_code    VARCHAR(75),
    phone       VARCHAR(25),
    manager_id  INTEGER REFERENCES employees(id),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_warehouses_region ON warehouses(region_id);

-- Inventory (composite PK)
CREATE TABLE inventory (
    product_id                  INTEGER NOT NULL REFERENCES products(id),
    warehouse_id                INTEGER NOT NULL REFERENCES warehouses(id),
    amount_in_stock             INTEGER DEFAULT 0,
    reorder_point               INTEGER,
    max_in_stock                INTEGER,
    out_of_stock_explanation    TEXT,
    restock_date                TIMESTAMPTZ,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (product_id, warehouse_id)
);

-- Orders
CREATE TABLE orders (
    id              SERIAL PRIMARY KEY,
    customer_id     INTEGER NOT NULL REFERENCES customers(id),
    date_ordered    TIMESTAMPTZ,
    date_shipped    TIMESTAMPTZ,
    sales_rep_id    INTEGER REFERENCES employees(id),
    total           NUMERIC(11,2),
    payment_type    payment_type,
    order_filled    BOOLEAN DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_orders_customer ON orders(customer_id);
CREATE INDEX idx_orders_sales_rep ON orders(sales_rep_id);
CREATE INDEX idx_orders_date_ordered ON orders(date_ordered);

-- Order items (composite PK)
CREATE TABLE order_items (
    ord_id              INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    item_id             INTEGER NOT NULL,
    product_id          INTEGER NOT NULL REFERENCES products(id),
    price               NUMERIC(11,2),
    quantity            INTEGER,
    quantity_shipped    INTEGER,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (ord_id, item_id)
);

CREATE INDEX idx_order_items_product ON order_items(product_id);
CREATE INDEX idx_order_items_ord ON order_items(ord_id);

-- Users table (NEW - for authentication)
CREATE TABLE users (
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    employee_id     INTEGER REFERENCES employees(id),
    role            user_role NOT NULL DEFAULT 'viewer',
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_employee ON users(employee_id);

-- Auto-update updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply trigger to all tables with updated_at
CREATE TRIGGER trg_regions_updated_at BEFORE UPDATE ON regions FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_departments_updated_at BEFORE UPDATE ON departments FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_employees_updated_at BEFORE UPDATE ON employees FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_customers_updated_at BEFORE UPDATE ON customers FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_products_updated_at BEFORE UPDATE ON products FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_warehouses_updated_at BEFORE UPDATE ON warehouses FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_inventory_updated_at BEFORE UPDATE ON inventory FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_orders_updated_at BEFORE UPDATE ON orders FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_order_items_updated_at BEFORE UPDATE ON order_items FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at();
