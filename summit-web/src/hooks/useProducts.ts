import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type { Product, Inventory, PagedResponse } from "@/lib/types";

export function useProducts(params?: { page?: number; search?: string }) {
  const searchParams = new URLSearchParams();
  if (params?.page) searchParams.set("page", String(params.page));
  if (params?.search) searchParams.set("search", params.search);

  const query = searchParams.toString();
  return useQuery({
    queryKey: ["products", params],
    queryFn: () =>
      api.get<PagedResponse<Product>>(
        `/api/v1/products${query ? `?${query}` : ""}`
      ),
  });
}

export function useProduct(id: number) {
  return useQuery({
    queryKey: ["products", id],
    queryFn: () => api.get<Product>(`/api/v1/products/${id}`),
    enabled: !!id,
  });
}

export function useProductStock(productId: number) {
  return useQuery({
    queryKey: ["products", productId, "stock"],
    queryFn: () =>
      api.get<Inventory[]>(`/api/v1/products/${productId}/stock`),
    enabled: !!productId,
  });
}
