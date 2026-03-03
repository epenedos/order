import { z } from "zod";

export const loginSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(8, "Password must be at least 8 characters"),
});

export const registerSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(8, "Password must be at least 8 characters"),
  employee_id: z.number().optional(),
});

export const customerSchema = z.object({
  name: z.string().min(1, "Name is required").max(50),
  phone: z.string().max(25).optional().nullable(),
  address: z.string().max(400).optional().nullable(),
  city: z.string().max(30).optional().nullable(),
  state: z.string().max(20).optional().nullable(),
  country: z.string().max(30).optional().nullable(),
  zip_code: z.string().max(75).optional().nullable(),
  credit_rating: z.enum(["EXCELLENT", "GOOD", "POOR"]).optional().nullable(),
  sales_rep_id: z.number().optional().nullable(),
  region_id: z.number().optional().nullable(),
  comments: z.string().optional().nullable(),
});

export const orderSchema = z.object({
  customer_id: z.number({ required_error: "Customer is required" }),
  date_ordered: z.string().optional(),
  sales_rep_id: z.number().optional(),
  payment_type: z.enum(["CASH", "CREDIT"]).optional(),
});

export const orderItemSchema = z.object({
  product_id: z.number({ required_error: "Product is required" }),
  quantity: z.number().min(1, "Quantity must be at least 1"),
});

export type LoginFormData = z.infer<typeof loginSchema>;
export type RegisterFormData = z.infer<typeof registerSchema>;
export type CustomerFormData = z.infer<typeof customerSchema>;
export type OrderFormData = z.infer<typeof orderSchema>;
export type OrderItemFormData = z.infer<typeof orderItemSchema>;
