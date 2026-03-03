"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Search } from "lucide-react";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { useProducts } from "@/hooks/useProducts";
import { formatCurrency } from "@/lib/utils";

export default function ProductsPage() {
  const router = useRouter();
  const [search, setSearch] = useState("");
  const { data, isLoading } = useProducts({ search });

  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          <h1 className="text-2xl font-bold mb-4">Products</h1>

          <div className="relative mb-4 max-w-sm">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <input
              type="text"
              placeholder="Search products..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="w-full pl-9 pr-4 py-2 border rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-primary"
            />
          </div>

          {isLoading ? (
            <LoadingSpinner />
          ) : (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
              {data?.data?.map((product) => (
                <div
                  key={product.id}
                  onClick={() => router.push(`/products/${product.id}`)}
                  className="border rounded-lg bg-white p-4 hover:shadow-md transition-shadow cursor-pointer"
                >
                  <h3 className="font-semibold text-sm">{product.name}</h3>
                  <p className="text-xs text-muted-foreground mt-1 line-clamp-2">
                    {product.short_desc || "No description"}
                  </p>
                  <p className="text-sm font-medium mt-2">
                    {formatCurrency(product.suggested_whlsl_price)}
                  </p>
                </div>
              ))}
              {data?.data?.length === 0 && (
                <p className="col-span-full text-center text-muted-foreground py-8">
                  No products found.
                </p>
              )}
            </div>
          )}
        </main>
      </div>
    </div>
  );
}
