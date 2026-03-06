"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Plus } from "lucide-react";
import { type ColumnDef } from "@tanstack/react-table";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { DataTable } from "@/components/shared/DataTable";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { AuthGuard } from "@/components/shared/AuthGuard";
import { OrderForm } from "@/components/orders/OrderForm";
import { useOrders } from "@/hooks/useOrders";
import { formatDate, formatCurrency } from "@/lib/utils";
import type { Order } from "@/lib/types";

const columns: ColumnDef<Order, unknown>[] = [
  { accessorKey: "id", header: "Order ID" },
  { accessorKey: "customer_name", header: "Customer" },
  {
    accessorKey: "date_ordered",
    header: "Date Ordered",
    cell: ({ getValue }) => formatDate(getValue() as string | null),
  },
  {
    accessorKey: "total",
    header: "Total",
    cell: ({ getValue }) => formatCurrency(getValue() as number | null),
  },
  { accessorKey: "payment_type", header: "Payment" },
  {
    accessorKey: "order_filled",
    header: "Filled",
    cell: ({ getValue }) => (getValue() ? "Yes" : "No"),
  },
  { accessorKey: "sales_rep_name", header: "Sales Rep" },
];

export default function OrdersPage() {
  const router = useRouter();
  const { data, isLoading } = useOrders({ limit: 50 });
  const [showCreateForm, setShowCreateForm] = useState(false);

  return (
    <AuthGuard>
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          <div className="flex justify-between items-center mb-4">
            <h1 className="text-2xl font-bold">Orders</h1>
            <button
              onClick={() => setShowCreateForm(true)}
              className="flex items-center gap-1 px-3 py-2 text-sm bg-primary text-primary-foreground rounded-md hover:bg-primary/90"
            >
              <Plus className="h-4 w-4" /> New Order
            </button>
          </div>
          {isLoading ? (
            <LoadingSpinner />
          ) : (
            <DataTable
              columns={columns}
              data={data?.data || []}
              onRowClick={(row) => router.push(`/orders/${row.id}`)}
            />
          )}

          <OrderForm
            open={showCreateForm}
            onClose={() => setShowCreateForm(false)}
            onCreated={(order) => router.push(`/orders/${order.id}`)}
          />
        </main>
      </div>
    </div>
    </AuthGuard>
  );
}
