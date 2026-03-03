import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type { SalesRep, Employee } from "@/lib/types";

export function useSalesReps() {
  return useQuery({
    queryKey: ["sales-reps"],
    queryFn: () => api.get<SalesRep[]>("/api/v1/employees/sales-reps"),
  });
}

export function useEmployees() {
  return useQuery({
    queryKey: ["employees"],
    queryFn: () => api.get<Employee[]>("/api/v1/employees"),
  });
}
