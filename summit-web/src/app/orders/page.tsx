"use client";

import { useRouter } from "next/navigation";
import { type ColumnDef } from "@tanstack/react-table";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { DataTable } from "@/components/shared/DataTable";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { AuthGuard } from "@/components/shared/AuthGuard";
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

  return (
    <AuthGuard>
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          <h1 className="text-2xl font-bold mb-4">Orders</h1>
          {isLoading ? (
            <LoadingSpinner />
          ) : (
            <DataTable
              columns={columns}
              data={data?.data || []}
              onRowClick={(row) => router.push(`/orders/${row.id}`)}
            />
          )}
        </main>
      </div>
    </div>
    </AuthGuard>
  );
}
