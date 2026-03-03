// TypeScript interfaces matching Go backend models

export interface Region {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface Department {
  id: number;
  name: string;
  region_id: number | null;
  created_at: string;
  updated_at: string;
}

export interface Employee {
  id: number;
  last_name: string;
  first_name: string | null;
  userid: string | null;
  start_date: string | null;
  comments: string | null;
  manager_id: number | null;
  title: string | null;
  dept_id: number | null;
  salary: number | null;
  commission_pct: number | null;
  created_at: string;
  updated_at: string;
}

export interface SalesRep {
  id: number;
  full_name: string;
  title: string;
}

export interface Customer {
  id: number;
  name: string;
  phone: string | null;
  address: string | null;
  city: string | null;
  state: string | null;
  country: string | null;
  zip_code: string | null;
  credit_rating: "EXCELLENT" | "GOOD" | "POOR" | null;
  sales_rep_id: number | null;
  region_id: number | null;
  comments: string | null;
  created_at: string;
  updated_at: string;
  sales_rep_name?: string;
}

export interface Order {
  id: number;
  customer_id: number;
  date_ordered: string | null;
  date_shipped: string | null;
  sales_rep_id: number | null;
  total: number | null;
  payment_type: "CASH" | "CREDIT" | null;
  order_filled: boolean;
  created_at: string;
  updated_at: string;
  customer_name?: string;
  sales_rep_name?: string;
  items?: OrderItem[];
}

export interface OrderItem {
  ord_id: number;
  item_id: number;
  product_id: number;
  price: number | null;
  quantity: number | null;
  quantity_shipped: number | null;
  created_at: string;
  updated_at: string;
  product_name?: string;
  product_image_url?: string;
}

export interface Product {
  id: number;
  name: string;
  short_desc: string | null;
  longtext_id: number | null;
  image_id: number | null;
  suggested_whlsl_price: number | null;
  whlsl_units: string | null;
  created_at: string;
  updated_at: string;
  image_url?: string;
  description?: string;
}

export interface Inventory {
  product_id: number;
  warehouse_id: number;
  amount_in_stock: number;
  reorder_point: number | null;
  max_in_stock: number | null;
  out_of_stock_explanation: string | null;
  restock_date: string | null;
  created_at: string;
  updated_at: string;
  warehouse_city?: string;
  product_name?: string;
}

export interface Warehouse {
  id: number;
  region_id: number;
  address: string | null;
  city: string | null;
  state: string | null;
  country: string | null;
  zip_code: string | null;
  phone: string | null;
  manager_id: number | null;
  created_at: string;
  updated_at: string;
}

export interface TreeNode {
  id: string;
  label: string;
  value: string | number;
  children?: TreeNode[];
}

export interface User {
  id: number;
  email: string;
  employee_id: number | null;
  role: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}

export interface PagedResponse<T> {
  data: T[];
  page: number;
  limit: number;
  total_count: number;
  total_pages: number;
}

// Request types
export interface CreateCustomerRequest {
  name: string;
  phone?: string;
  address?: string;
  city?: string;
  state?: string;
  country?: string;
  zip_code?: string;
  credit_rating?: "EXCELLENT" | "GOOD" | "POOR";
  sales_rep_id?: number;
  region_id?: number;
  comments?: string;
}

export interface CreateOrderRequest {
  customer_id: number;
  date_ordered?: string;
  sales_rep_id?: number;
  payment_type?: "CASH" | "CREDIT";
}

export interface CreateOrderItemRequest {
  product_id: number;
  quantity: number;
}
