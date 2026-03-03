"use client";

import { useQuery } from "@tanstack/react-query";
import { Sidebar } from "@/components/layout/Sidebar";
import { Header } from "@/components/layout/Header";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { api } from "@/lib/api";
import type { Department } from "@/lib/types";

export default function DepartmentsPage() {
  const { data: departments, isLoading } = useQuery({
    queryKey: ["departments"],
    queryFn: () => api.get<Department[]>("/api/v1/departments"),
  });

  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 p-6">
          <h1 className="text-2xl font-bold mb-4">Departments</h1>
          {isLoading ? (
            <LoadingSpinner />
          ) : (
            <div className="border rounded-md bg-white">
              <table className="w-full text-sm">
                <thead>
                  <tr className="border-b bg-gray-50">
                    <th className="px-4 py-3 text-left font-medium">ID</th>
                    <th className="px-4 py-3 text-left font-medium">Name</th>
                    <th className="px-4 py-3 text-left font-medium">Region ID</th>
                  </tr>
                </thead>
                <tbody>
                  {departments?.map((dept) => (
                    <tr key={dept.id} className="border-b">
                      <td className="px-4 py-3">{dept.id}</td>
                      <td className="px-4 py-3">{dept.name}</td>
                      <td className="px-4 py-3">{dept.region_id ?? "-"}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </main>
      </div>
    </div>
  );
}
