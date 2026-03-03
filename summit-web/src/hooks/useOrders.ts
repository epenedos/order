import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type {
  Order,
  OrderItem,
  PagedResponse,
  CreateOrderRequest,
  CreateOrderItemRequest,
} from "@/lib/types";

export function useOrders(params?: {
  page?: number;
  limit?: number;
  customer_id?: number;
}) {
  const searchParams = new URLSearchParams();
  if (params?.page) searchParams.set("page", String(params.page));
  if (params?.limit) searchParams.set("limit", String(params.limit));
  if (params?.customer_id)
    searchParams.set("customer_id", String(params.customer_id));

  const query = searchParams.toString();
  return useQuery({
    queryKey: ["orders", params],
    queryFn: () =>
      api.get<PagedResponse<Order>>(
        `/api/v1/orders${query ? `?${query}` : ""}`
      ),
  });
}

export function useOrder(id: number) {
  return useQuery({
    queryKey: ["orders", id],
    queryFn: () => api.get<Order>(`/api/v1/orders/${id}`),
    enabled: !!id,
  });
}

export function useCreateOrder() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: CreateOrderRequest) =>
      api.post<Order>("/api/v1/orders", data),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["orders"] }),
  });
}

export function useDeleteOrder() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: number) => api.delete(`/api/v1/orders/${id}`),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["orders"] }),
  });
}

export function useOrderItems(orderId: number) {
  return useQuery({
    queryKey: ["orders", orderId, "items"],
    queryFn: () => api.get<OrderItem[]>(`/api/v1/orders/${orderId}/items`),
    enabled: !!orderId,
  });
}

export function useAddOrderItem(orderId: number) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: CreateOrderItemRequest) =>
      api.post<OrderItem>(`/api/v1/orders/${orderId}/items`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["orders", orderId, "items"],
      });
      queryClient.invalidateQueries({ queryKey: ["orders", orderId] });
    },
  });
}

export function useDeleteOrderItem(orderId: number) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (itemId: number) =>
      api.delete(`/api/v1/orders/${orderId}/items/${itemId}`),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["orders", orderId, "items"],
      });
      queryClient.invalidateQueries({ queryKey: ["orders", orderId] });
    },
  });
}
