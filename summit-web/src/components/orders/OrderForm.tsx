"use client";

import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  orderSchema,
  updateOrderSchema,
  type OrderFormData,
  type UpdateOrderFormData,
} from "@/lib/validators";
import { useCreateOrder, useUpdateOrder } from "@/hooks/useOrders";
import { useCustomers } from "@/hooks/useCustomers";
import { useSalesReps } from "@/hooks/useEmployees";
import type { Order, CreateOrderRequest, UpdateOrderRequest } from "@/lib/types";

interface OrderFormProps {
  open: boolean;
  onClose: () => void;
  order?: Order;
  onCreated?: (order: Order) => void;
}

export function OrderForm({ open, onClose, order, onCreated }: OrderFormProps) {
  const isEdit = !!order;
  const createOrder = useCreateOrder();
  const updateOrder = useUpdateOrder(order?.id ?? 0);
  const { data: customersData } = useCustomers({ limit: 200 });
  const { data: salesReps } = useSalesReps();

  const createForm = useForm<OrderFormData>({
    resolver: zodResolver(orderSchema),
    defaultValues: { customer_id: undefined, payment_type: undefined, date_ordered: "" },
  });

  const editForm = useForm<UpdateOrderFormData>({
    resolver: zodResolver(updateOrderSchema),
    defaultValues: order
      ? {
          date_ordered: order.date_ordered?.split("T")[0] ?? "",
          date_shipped: order.date_shipped?.split("T")[0] ?? "",
          payment_type: order.payment_type ?? undefined,
          order_filled: order.order_filled,
          sales_rep_id: order.sales_rep_id ?? undefined,
        }
      : {},
  });

  useEffect(() => {
    if (open && order) {
      editForm.reset({
        date_ordered: order.date_ordered?.split("T")[0] ?? "",
        date_shipped: order.date_shipped?.split("T")[0] ?? "",
        payment_type: order.payment_type ?? undefined,
        order_filled: order.order_filled,
        sales_rep_id: order.sales_rep_id ?? undefined,
      });
    } else if (open && !order) {
      createForm.reset({ customer_id: undefined, payment_type: undefined, date_ordered: "" });
    }
  }, [open, order, createForm, editForm]);

  if (!open) return null;

  const onCreateSubmit = async (data: OrderFormData) => {
    const payload: CreateOrderRequest = {
      customer_id: data.customer_id,
      ...(data.date_ordered
        ? { date_ordered: data.date_ordered + "T00:00:00Z" }
        : {}),
      ...(data.payment_type ? { payment_type: data.payment_type } : {}),
    };
    const result = await createOrder.mutateAsync(payload);
    onCreated?.(result);
    onClose();
  };

  const onEditSubmit = async (data: UpdateOrderFormData) => {
    const payload: UpdateOrderRequest = {
      ...(data.date_ordered
        ? { date_ordered: data.date_ordered + "T00:00:00Z" }
        : {}),
      ...(data.date_shipped
        ? { date_shipped: data.date_shipped + "T00:00:00Z" }
        : {}),
      ...(data.payment_type ? { payment_type: data.payment_type } : {}),
      ...(data.sales_rep_id ? { sales_rep_id: data.sales_rep_id } : {}),
      ...(data.order_filled !== undefined
        ? { order_filled: data.order_filled }
        : {}),
    };
    await updateOrder.mutateAsync(payload);
    onClose();
  };

  if (isEdit) {
    return (
      <div className="fixed inset-0 z-50 flex items-center justify-center">
        <div className="fixed inset-0 bg-black/50" onClick={onClose} />
        <div className="relative bg-white rounded-lg shadow-lg p-6 max-w-lg w-full mx-4">
          <h2 className="text-lg font-semibold mb-4">Edit Order #{order.id}</h2>
          <form onSubmit={editForm.handleSubmit(onEditSubmit)} className="space-y-4">
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium mb-1">Date Ordered</label>
                <input
                  type="date"
                  {...editForm.register("date_ordered")}
                  className="w-full border rounded-md px-3 py-2 text-sm"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-1">Date Shipped</label>
                <input
                  type="date"
                  {...editForm.register("date_shipped")}
                  className="w-full border rounded-md px-3 py-2 text-sm"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-1">Payment Type</label>
                <select
                  {...editForm.register("payment_type")}
                  className="w-full border rounded-md px-3 py-2 text-sm"
                >
                  <option value="">-- Select --</option>
                  <option value="CASH">Cash</option>
                  <option value="CREDIT">Credit</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium mb-1">Sales Rep</label>
                <select
                  {...editForm.register("sales_rep_id", { valueAsNumber: true })}
                  className="w-full border rounded-md px-3 py-2 text-sm"
                >
                  <option value="">-- Select --</option>
                  {salesReps?.map((rep) => (
                    <option key={rep.id} value={rep.id}>
                      {rep.full_name}
                    </option>
                  ))}
                </select>
              </div>
              <div className="col-span-2">
                <label className="flex items-center gap-2 text-sm font-medium">
                  <input
                    type="checkbox"
                    {...editForm.register("order_filled")}
                    className="rounded"
                  />
                  Order Filled
                </label>
              </div>
            </div>
            <div className="flex justify-end gap-2 pt-2">
              <button
                type="button"
                onClick={onClose}
                className="px-4 py-2 text-sm border rounded-md hover:bg-gray-50"
              >
                Cancel
              </button>
              <button
                type="submit"
                disabled={editForm.formState.isSubmitting}
                className="px-4 py-2 text-sm rounded-md text-white bg-primary hover:bg-primary/90 disabled:opacity-50"
              >
                {editForm.formState.isSubmitting ? "Saving..." : "Update"}
              </button>
            </div>
          </form>
        </div>
      </div>
    );
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="fixed inset-0 bg-black/50" onClick={onClose} />
      <div className="relative bg-white rounded-lg shadow-lg p-6 max-w-lg w-full mx-4">
        <h2 className="text-lg font-semibold mb-4">New Order</h2>
        <form onSubmit={createForm.handleSubmit(onCreateSubmit)} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="col-span-2">
              <label className="block text-sm font-medium mb-1">
                Customer <span className="text-red-500">*</span>
              </label>
              <select
                {...createForm.register("customer_id", { valueAsNumber: true })}
                className="w-full border rounded-md px-3 py-2 text-sm"
              >
                <option value="">-- Select Customer --</option>
                {customersData?.data?.map((c) => (
                  <option key={c.id} value={c.id}>
                    {c.name} (#{c.id})
                  </option>
                ))}
              </select>
              {createForm.formState.errors.customer_id && (
                <p className="text-red-500 text-xs mt-1">
                  {createForm.formState.errors.customer_id.message}
                </p>
              )}
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Date Ordered</label>
              <input
                type="date"
                {...createForm.register("date_ordered")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Payment Type</label>
              <select
                {...createForm.register("payment_type")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              >
                <option value="">-- Select --</option>
                <option value="CASH">Cash</option>
                <option value="CREDIT">Credit</option>
              </select>
            </div>
          </div>
          <div className="flex justify-end gap-2 pt-2">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 text-sm border rounded-md hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={createForm.formState.isSubmitting}
              className="px-4 py-2 text-sm rounded-md text-white bg-primary hover:bg-primary/90 disabled:opacity-50"
            >
              {createForm.formState.isSubmitting ? "Creating..." : "Create Order"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
