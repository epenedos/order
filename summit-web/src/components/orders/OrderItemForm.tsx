"use client";

import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  orderItemSchema,
  updateOrderItemSchema,
  type OrderItemFormData,
  type UpdateOrderItemFormData,
} from "@/lib/validators";
import { useAddOrderItem, useUpdateOrderItem } from "@/hooks/useOrders";
import { useProducts } from "@/hooks/useProducts";
import type { OrderItem } from "@/lib/types";

interface OrderItemFormProps {
  open: boolean;
  onClose: () => void;
  orderId: number;
  item?: OrderItem;
}

export function OrderItemForm({ open, onClose, orderId, item }: OrderItemFormProps) {
  const isEdit = !!item;
  const addItem = useAddOrderItem(orderId);
  const updateItem = useUpdateOrderItem(orderId);
  const { data: productsData } = useProducts({ search: "" });

  const addForm = useForm<OrderItemFormData>({
    resolver: zodResolver(orderItemSchema),
    defaultValues: { product_id: undefined, quantity: 1 },
  });

  const editForm = useForm<UpdateOrderItemFormData>({
    resolver: zodResolver(updateOrderItemSchema),
    defaultValues: item
      ? { quantity: item.quantity ?? 1, quantity_shipped: item.quantity_shipped ?? undefined }
      : {},
  });

  useEffect(() => {
    if (open && item) {
      editForm.reset({
        quantity: item.quantity ?? 1,
        quantity_shipped: item.quantity_shipped ?? undefined,
      });
    } else if (open && !item) {
      addForm.reset({ product_id: undefined, quantity: 1 });
    }
  }, [open, item, addForm, editForm]);

  if (!open) return null;

  const onAddSubmit = async (data: OrderItemFormData) => {
    await addItem.mutateAsync(data);
    onClose();
  };

  const onEditSubmit = async (data: UpdateOrderItemFormData) => {
    await updateItem.mutateAsync({ itemId: item!.item_id, data });
    onClose();
  };

  if (isEdit) {
    return (
      <div className="fixed inset-0 z-50 flex items-center justify-center">
        <div className="fixed inset-0 bg-black/50" onClick={onClose} />
        <div className="relative bg-white rounded-lg shadow-lg p-6 max-w-md w-full mx-4">
          <h2 className="text-lg font-semibold mb-4">
            Edit Item #{item.item_id} — {item.product_name || `Product #${item.product_id}`}
          </h2>
          <form onSubmit={editForm.handleSubmit(onEditSubmit)} className="space-y-4">
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium mb-1">Quantity</label>
                <input
                  type="number"
                  min={1}
                  {...editForm.register("quantity", { valueAsNumber: true })}
                  className="w-full border rounded-md px-3 py-2 text-sm"
                />
                {editForm.formState.errors.quantity && (
                  <p className="text-red-500 text-xs mt-1">
                    {editForm.formState.errors.quantity.message}
                  </p>
                )}
              </div>
              <div>
                <label className="block text-sm font-medium mb-1">Qty Shipped</label>
                <input
                  type="number"
                  min={0}
                  {...editForm.register("quantity_shipped", { valueAsNumber: true })}
                  className="w-full border rounded-md px-3 py-2 text-sm"
                />
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
      <div className="relative bg-white rounded-lg shadow-lg p-6 max-w-md w-full mx-4">
        <h2 className="text-lg font-semibold mb-4">Add Item to Order</h2>
        <form onSubmit={addForm.handleSubmit(onAddSubmit)} className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-1">
              Product <span className="text-red-500">*</span>
            </label>
            <select
              {...addForm.register("product_id", { valueAsNumber: true })}
              className="w-full border rounded-md px-3 py-2 text-sm"
            >
              <option value="">-- Select Product --</option>
              {productsData?.data?.map((p) => (
                <option key={p.id} value={p.id}>
                  {p.name} — {p.suggested_whlsl_price != null ? `$${p.suggested_whlsl_price.toFixed(2)}` : "N/A"}
                </option>
              ))}
            </select>
            {addForm.formState.errors.product_id && (
              <p className="text-red-500 text-xs mt-1">
                {addForm.formState.errors.product_id.message}
              </p>
            )}
          </div>
          <div>
            <label className="block text-sm font-medium mb-1">
              Quantity <span className="text-red-500">*</span>
            </label>
            <input
              type="number"
              min={1}
              {...addForm.register("quantity", { valueAsNumber: true })}
              className="w-full border rounded-md px-3 py-2 text-sm"
            />
            {addForm.formState.errors.quantity && (
              <p className="text-red-500 text-xs mt-1">
                {addForm.formState.errors.quantity.message}
              </p>
            )}
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
              disabled={addForm.formState.isSubmitting}
              className="px-4 py-2 text-sm rounded-md text-white bg-primary hover:bg-primary/90 disabled:opacity-50"
            >
              {addForm.formState.isSubmitting ? "Adding..." : "Add Item"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
