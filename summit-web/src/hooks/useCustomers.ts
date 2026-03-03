import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type {
  Customer,
  PagedResponse,
  TreeNode,
  CreateCustomerRequest,
} from "@/lib/types";

export function useCustomers(params?: {
  page?: number;
  limit?: number;
  sort?: string;
  country?: string;
  sales_rep_id?: number;
}) {
  const searchParams = new URLSearchParams();
  if (params?.page) searchParams.set("page", String(params.page));
  if (params?.limit) searchParams.set("limit", String(params.limit));
  if (params?.sort) searchParams.set("sort", params.sort);
  if (params?.country) searchParams.set("country", params.country);
  if (params?.sales_rep_id)
    searchParams.set("sales_rep_id", String(params.sales_rep_id));

  const query = searchParams.toString();
  return useQuery({
    queryKey: ["customers", params],
    queryFn: () =>
      api.get<PagedResponse<Customer>>(
        `/api/v1/customers${query ? `?${query}` : ""}`
      ),
  });
}

export function useCustomer(id: number) {
  return useQuery({
    queryKey: ["customers", id],
    queryFn: () => api.get<Customer>(`/api/v1/customers/${id}`),
    enabled: !!id,
  });
}

export function useCustomerTree(mode: "by_country" | "by_sales_rep") {
  return useQuery({
    queryKey: ["customer-tree", mode],
    queryFn: () =>
      api.get<TreeNode[]>(`/api/v1/customers/tree?mode=${mode}`),
  });
}

export function useCreateCustomer() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: CreateCustomerRequest) =>
      api.post<Customer>("/api/v1/customers", data),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["customers"] }),
  });
}

export function useUpdateCustomer(id: number) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: Partial<CreateCustomerRequest>) =>
      api.put<Customer>(`/api/v1/customers/${id}`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["customers"] });
      queryClient.invalidateQueries({ queryKey: ["customers", id] });
    },
  });
}

export function useDeleteCustomer() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: number) => api.delete(`/api/v1/customers/${id}`),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["customers"] }),
  });
}
