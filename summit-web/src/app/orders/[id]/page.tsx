"use client";

import { useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { Trash2, Plus } from "lucide-react";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { ConfirmDialog } from "@/components/shared/ConfirmDialog";
import { useOrder, useDeleteOrder, useDeleteOrderItem } from "@/hooks/useOrders";
import { formatDate, formatCurrency } from "@/lib/utils";

export default function OrderDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = Number(params.id);
  const { data: order, isLoading } = useOrder(id);
  const deleteOrder = useDeleteOrder();
  const deleteItem = useDeleteOrderItem(id);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

  if (isLoading) {
    return (
      <div className="flex min-h-screen">
        <Sidebar />
        <div className="flex-1"><LoadingSpinner /></div>
      </div>
    );
  }

  if (!order) return <p>Order not found.</p>;

  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          <div className="flex justify-between items-center mb-6">
            <h1 className="text-2xl font-bold">Order #{order.id}</h1>
            <div className="flex gap-2">
              <button
                onClick={() => setShowDeleteConfirm(true)}
                className="flex items-center gap-1 px-3 py-2 text-sm text-red-600 border border-red-200 rounded-md hover:bg-red-50"
              >
                <Trash2 className="h-4 w-4" /> Delete Order
              </button>
            </div>
          </div>

          {/* Order Header */}
          <div className="border rounded-lg bg-white p-6 mb-6">
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
              <Field label="Customer" value={order.customer_name} />
              <Field label="Date Ordered" value={formatDate(order.date_ordered)} />
              <Field label="Date Shipped" value={formatDate(order.date_shipped)} />
              <Field label="Sales Rep" value={order.sales_rep_name} />
              <Field label="Payment Type" value={order.payment_type} />
              <Field label="Total" value={formatCurrency(order.total)} />
              <Field label="Order Filled" value={order.order_filled ? "Yes" : "No"} />
            </div>
          </div>

          {/* Order Items */}
          <div className="border rounded-lg bg-white">
            <div className="p-4 border-b flex justify-between items-center">
              <h2 className="font-semibold">Line Items</h2>
            </div>
            <table className="w-full text-sm">
              <thead>
                <tr className="border-b bg-gray-50">
                  <th className="px-4 py-3 text-left font-medium">#</th>
                  <th className="px-4 py-3 text-left font-medium">Product</th>
                  <th className="px-4 py-3 text-right font-medium">Price</th>
                  <th className="px-4 py-3 text-right font-medium">Qty</th>
                  <th className="px-4 py-3 text-right font-medium">Shipped</th>
                  <th className="px-4 py-3 text-right font-medium">Subtotal</th>
                  <th className="px-4 py-3 w-10"></th>
                </tr>
              </thead>
              <tbody>
                {order.items?.map((item) => (
                  <tr key={item.item_id} className="border-b">
                    <td className="px-4 py-3">{item.item_id}</td>
                    <td className="px-4 py-3">{item.product_name || `Product #${item.product_id}`}</td>
                    <td className="px-4 py-3 text-right">{formatCurrency(item.price)}</td>
                    <td className="px-4 py-3 text-right">{item.quantity}</td>
                    <td className="px-4 py-3 text-right">{item.quantity_shipped ?? "-"}</td>
                    <td className="px-4 py-3 text-right">
                      {formatCurrency(
                        item.price && item.quantity ? item.price * item.quantity : null
                      )}
                    </td>
                    <td className="px-4 py-3">
                      <button
                        onClick={() => deleteItem.mutate(item.item_id)}
                        className="text-red-500 hover:text-red-700"
                        title="Delete item"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    </td>
                  </tr>
                ))}
                {(!order.items || order.items.length === 0) && (
                  <tr>
                    <td colSpan={7} className="px-4 py-8 text-center text-muted-foreground">
                      No items in this order.
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>

          <ConfirmDialog
            open={showDeleteConfirm}
            title="Delete Order"
            message="Are you sure you want to delete this order? This action cannot be undone."
            variant="destructive"
            confirmLabel="Delete"
            onConfirm={async () => {
              await deleteOrder.mutateAsync(id);
              router.push("/orders");
            }}
            onCancel={() => setShowDeleteConfirm(false)}
          />
        </main>
      </div>
    </div>
  );
}

function Field({ label, value }: { label: string; value: string | null | undefined }) {
  return (
    <div>
      <dt className="text-sm font-medium text-muted-foreground">{label}</dt>
      <dd className="mt-1 text-sm">{value || "-"}</dd>
    </div>
  );
}
