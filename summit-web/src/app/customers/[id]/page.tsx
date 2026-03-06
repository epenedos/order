"use client";

import { useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { Pencil, Trash2 } from "lucide-react";
import { useCustomer, useDeleteCustomer } from "@/hooks/useCustomers";
import { LoadingSpinner } from "@/components/shared/LoadingSpinner";
import { ConfirmDialog } from "@/components/shared/ConfirmDialog";
import { CustomerForm } from "@/components/customers/CustomerForm";
import { formatDate } from "@/lib/utils";

export default function CustomerDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = Number(params.id);
  const { data: customer, isLoading } = useCustomer(id);
  const deleteCustomer = useDeleteCustomer();
  const [showEditForm, setShowEditForm] = useState(false);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

  if (isLoading) return <LoadingSpinner />;
  if (!customer) return <p>Customer not found.</p>;

  return (
    <div className="max-w-4xl">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">{customer.name}</h1>
        <div className="flex gap-2">
          <button
            onClick={() => setShowEditForm(true)}
            className="flex items-center gap-1 px-3 py-2 text-sm border rounded-md hover:bg-gray-50"
          >
            <Pencil className="h-4 w-4" /> Edit
          </button>
          <button
            onClick={() => setShowDeleteConfirm(true)}
            className="flex items-center gap-1 px-3 py-2 text-sm text-red-600 border border-red-200 rounded-md hover:bg-red-50"
          >
            <Trash2 className="h-4 w-4" /> Delete
          </button>
          <button
            onClick={() => router.push(`/customers/${id}/orders`)}
            className="px-4 py-2 bg-primary text-primary-foreground rounded-md text-sm hover:bg-primary/90"
          >
            View Orders
          </button>
        </div>
      </div>

      {/* Tabs */}
      <div className="border rounded-lg bg-white">
        {/* General Tab */}
        <div className="p-6 space-y-4">
          <h2 className="text-lg font-semibold border-b pb-2">General Information</h2>
          <div className="grid grid-cols-2 gap-4">
            <Field label="Customer ID" value={String(customer.id)} />
            <Field label="Phone" value={customer.phone} />
            <Field label="Credit Rating" value={customer.credit_rating} />
            <Field label="Sales Rep" value={customer.sales_rep_name} />
          </div>

          <h2 className="text-lg font-semibold border-b pb-2 mt-6">Address</h2>
          <div className="grid grid-cols-2 gap-4">
            <Field label="Address" value={customer.address} />
            <Field label="City" value={customer.city} />
            <Field label="State" value={customer.state} />
            <Field label="Country" value={customer.country} />
            <Field label="Zip Code" value={customer.zip_code} />
          </div>

          {customer.comments && (
            <>
              <h2 className="text-lg font-semibold border-b pb-2 mt-6">Comments</h2>
              <p className="text-sm text-muted-foreground">{customer.comments}</p>
            </>
          )}
        </div>
      </div>

      <CustomerForm
        open={showEditForm}
        onClose={() => setShowEditForm(false)}
        customer={customer}
      />

      <ConfirmDialog
        open={showDeleteConfirm}
        title="Delete Customer"
        message={`Are you sure you want to delete "${customer.name}"? This action cannot be undone.`}
        variant="destructive"
        confirmLabel="Delete"
        onConfirm={async () => {
          await deleteCustomer.mutateAsync(id);
          router.push("/customers");
        }}
        onCancel={() => setShowDeleteConfirm(false)}
      />
    </div>
  );
}

function Field({ label, value }: { label: string; value: string | null | undefined }) {
  return (
    <div>
      <dt className="text-sm font-medium text-muted-foreground">{label}</dt>
      <dd className="mt-1 text-sm">{value || "-"}</dd>
    </div>
  );
}
