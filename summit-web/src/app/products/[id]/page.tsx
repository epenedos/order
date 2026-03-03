"use client";

import { useParams } from "next/navigation";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { useProduct, useProductStock } from "@/hooks/useProducts";
import { formatCurrency } from "@/lib/utils";

export default function ProductDetailPage() {
  const params = useParams();
  const id = Number(params.id);
  const { data: product, isLoading } = useProduct(id);
  const { data: stock } = useProductStock(id);

  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          {isLoading ? (
            <LoadingSpinner />
          ) : !product ? (
            <p>Product not found.</p>
          ) : (
            <div className="max-w-4xl">
              <h1 className="text-2xl font-bold mb-6">{product.name}</h1>

              <div className="border rounded-lg bg-white p-6 mb-6">
                <div className="grid grid-cols-2 gap-4">
                  <div>
                    <dt className="text-sm font-medium text-muted-foreground">Price</dt>
                    <dd className="mt-1">{formatCurrency(product.suggested_whlsl_price)}</dd>
                  </div>
                  <div>
                    <dt className="text-sm font-medium text-muted-foreground">Units</dt>
                    <dd className="mt-1">{product.whlsl_units || "-"}</dd>
                  </div>
                  <div className="col-span-2">
                    <dt className="text-sm font-medium text-muted-foreground">Description</dt>
                    <dd className="mt-1">{product.short_desc || product.description || "-"}</dd>
                  </div>
                </div>
              </div>

              {/* Stock Levels */}
              <div className="border rounded-lg bg-white">
                <div className="p-4 border-b">
                  <h2 className="font-semibold">Inventory / Stock Levels</h2>
                </div>
                <table className="w-full text-sm">
                  <thead>
                    <tr className="border-b bg-gray-50">
                      <th className="px-4 py-3 text-left font-medium">Warehouse</th>
                      <th className="px-4 py-3 text-right font-medium">In Stock</th>
                      <th className="px-4 py-3 text-right font-medium">Reorder Point</th>
                      <th className="px-4 py-3 text-right font-medium">Max Stock</th>
                    </tr>
                  </thead>
                  <tbody>
                    {stock?.map((inv) => (
                      <tr key={inv.warehouse_id} className="border-b">
                        <td className="px-4 py-3">{inv.warehouse_city || `Warehouse #${inv.warehouse_id}`}</td>
                        <td className="px-4 py-3 text-right">{inv.amount_in_stock}</td>
                        <td className="px-4 py-3 text-right">{inv.reorder_point ?? "-"}</td>
                        <td className="px-4 py-3 text-right">{inv.max_in_stock ?? "-"}</td>
                      </tr>
                    ))}
                    {(!stock || stock.length === 0) && (
                      <tr>
                        <td colSpan={4} className="px-4 py-8 text-center text-muted-foreground">
                          No inventory data.
                        </td>
                      </tr>
                    )}
                  </tbody>
                </table>
              </div>
            </div>
          )}
        </main>
      </div>
    </div>
  );
}
