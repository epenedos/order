"use client";

import { useParams, useRouter } from "next/navigation";
import { type ColumnDef } from "@tanstack/react-table";
import { useOrders } from "@/hooks/useOrders";
import { useCustomer } from "@/hooks/useCustomers";
import { DataTable } from "@/components/shared/DataTable";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { formatDate, formatCurrency } from "@/lib/utils";
import type { Order } from "@/lib/types";

const columns: ColumnDef<Order, unknown>[] = [
  { accessorKey: "id", header: "Order ID" },
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

export default function CustomerOrdersPage() {
  const params = useParams();
  const router = useRouter();
  const customerId = Number(params.id);

  const { data: customer } = useCustomer(customerId);
  const { data: ordersData, isLoading } = useOrders({ customer_id: customerId, limit: 50 });

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <div>
          <h1 className="text-2xl font-bold">
            Orders for {customer?.name || "..."}
          </h1>
          <button
            onClick={() => router.push(`/customers/${customerId}`)}
            className="text-sm text-muted-foreground hover:text-foreground"
          >
            Back to customer
          </button>
        </div>
      </div>

      {isLoading ? (
        <LoadingSpinner />
      ) : (
        <DataTable
          columns={columns}
          data={ordersData?.data || []}
          onRowClick={(row) => router.push(`/orders/${row.id}`)}
        />
      )}
    </div>
  );
}
