"use client";

import { useQuery } from "@tanstack/react-query";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { AuthGuard } from "@/components/shared/AuthGuard";
import { api } from "@/lib/api";
import type { Inventory } from "@/lib/types";

export default function InventoryPage() {
  const { data: inventory, isLoading } = useQuery({
    queryKey: ["inventory"],
    queryFn: () => api.get<Inventory[]>("/api/v1/inventory"),
  });

  return (
    <AuthGuard>
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          <h1 className="text-2xl font-bold mb-4">Inventory</h1>
          {isLoading ? (
            <LoadingSpinner />
          ) : (
            <div className="border rounded-md bg-white overflow-auto">
              <table className="w-full text-sm">
                <thead>
                  <tr className="border-b bg-gray-50">
                    <th className="px-4 py-3 text-left font-medium">Product ID</th>
                    <th className="px-4 py-3 text-left font-medium">Warehouse</th>
                    <th className="px-4 py-3 text-right font-medium">In Stock</th>
                    <th className="px-4 py-3 text-right font-medium">Reorder Point</th>
                    <th className="px-4 py-3 text-right font-medium">Max Stock</th>
                  </tr>
                </thead>
                <tbody>
                  {inventory?.map((inv) => (
                    <tr key={`${inv.product_id}-${inv.warehouse_id}`} className="border-b">
                      <td className="px-4 py-3">{inv.product_id}</td>
                      <td className="px-4 py-3">{inv.warehouse_city || inv.warehouse_id}</td>
                      <td className="px-4 py-3 text-right">{inv.amount_in_stock}</td>
                      <td className="px-4 py-3 text-right">{inv.reorder_point ?? "-"}</td>
                      <td className="px-4 py-3 text-right">{inv.max_in_stock ?? "-"}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </main>
      </div>
    </div>
    </AuthGuard>
  );
}
