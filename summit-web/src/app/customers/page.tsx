"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { type ColumnDef } from "@tanstack/react-table";
import { useCustomers, useCustomerTree } from "@/hooks/useCustomers";
import { DataTable } from "@/components/shared/DataTable";
import { TreeView } from "@/components/shared/TreeView";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import type { Customer, TreeNode } from "@/lib/types";

const columns: ColumnDef<Customer, unknown>[] = [
  { accessorKey: "name", header: "Name" },
  { accessorKey: "phone", header: "Phone" },
  { accessorKey: "city", header: "City" },
  { accessorKey: "country", header: "Country" },
  { accessorKey: "credit_rating", header: "Credit Rating" },
  { accessorKey: "sales_rep_name", header: "Sales Rep" },
];

export default function CustomersPage() {
  const router = useRouter();
  const [treeMode, setTreeMode] = useState<"by_country" | "by_sales_rep">(
    "by_country"
  );
  const [selectedNode, setSelectedNode] = useState<string | undefined>();
  const [filterParams, setFilterParams] = useState<{
    country?: string;
    sales_rep_id?: number;
  }>({});

  const { data: tree, isLoading: treeLoading } = useCustomerTree(treeMode);
  const { data: customersData, isLoading } = useCustomers({
    ...filterParams,
    limit: 50,
  });

  const handleNodeSelect = (node: TreeNode) => {
    setSelectedNode(node.id);
    if (node.id.startsWith("cust-")) {
      const customerId = node.value as number;
      router.push(`/customers/${customerId}`);
    } else if (node.id.startsWith("country-")) {
      setFilterParams({ country: node.value as string });
    } else if (node.id.startsWith("rep-")) {
      setFilterParams({ sales_rep_id: node.value as number });
    }
  };

  return (
    <div className="flex gap-6 h-full">
      {/* Tree Navigator */}
      <div className="w-64 flex-shrink-0 border rounded-md bg-white overflow-auto">
        <div className="p-3 border-b">
          <h2 className="font-semibold text-sm">Navigator</h2>
          <div className="flex gap-2 mt-2">
            <button
              onClick={() => {
                setTreeMode("by_country");
                setFilterParams({});
              }}
              className={`text-xs px-2 py-1 rounded ${
                treeMode === "by_country"
                  ? "bg-primary text-primary-foreground"
                  : "bg-gray-100"
              }`}
            >
              By Country
            </button>
            <button
              onClick={() => {
                setTreeMode("by_sales_rep");
                setFilterParams({});
              }}
              className={`text-xs px-2 py-1 rounded ${
                treeMode === "by_sales_rep"
                  ? "bg-primary text-primary-foreground"
                  : "bg-gray-100"
              }`}
            >
              By Sales Rep
            </button>
          </div>
        </div>
        {treeLoading ? (
          <LoadingSpinner />
        ) : (
          <TreeView
            nodes={tree || []}
            onSelect={handleNodeSelect}
            selectedId={selectedNode}
          />
        )}
      </div>

      {/* Customer List */}
      <div className="flex-1">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-2xl font-bold">Customers</h1>
          <button
            onClick={() => setFilterParams({})}
            className="text-sm text-muted-foreground hover:text-foreground"
          >
            Clear Filters
          </button>
        </div>
        {isLoading ? (
          <LoadingSpinner />
        ) : (
          <DataTable
            columns={columns}
            data={customersData?.data || []}
            onRowClick={(row) => router.push(`/customers/${row.id}`)}
          />
        )}
      </div>
    </div>
  );
}
