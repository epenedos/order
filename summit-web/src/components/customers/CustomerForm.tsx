"use client";

import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { customerSchema, type CustomerFormData } from "@/lib/validators";
import { useCreateCustomer, useUpdateCustomer } from "@/hooks/useCustomers";
import { useSalesReps } from "@/hooks/useEmployees";
import type { Customer } from "@/lib/types";

interface CustomerFormProps {
  open: boolean;
  onClose: () => void;
  customer?: Customer;
}

export function CustomerForm({ open, onClose, customer }: CustomerFormProps) {
  const isEdit = !!customer;
  const createCustomer = useCreateCustomer();
  const updateCustomer = useUpdateCustomer(customer?.id ?? 0);
  const { data: salesReps } = useSalesReps();

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors, isSubmitting },
  } = useForm<CustomerFormData>({
    resolver: zodResolver(customerSchema),
    defaultValues: customer
      ? {
          name: customer.name,
          phone: customer.phone,
          address: customer.address,
          city: customer.city,
          state: customer.state,
          country: customer.country,
          zip_code: customer.zip_code,
          credit_rating: customer.credit_rating,
          sales_rep_id: customer.sales_rep_id,
          comments: customer.comments,
        }
      : { name: "" },
  });

  useEffect(() => {
    if (open && customer) {
      reset({
        name: customer.name,
        phone: customer.phone,
        address: customer.address,
        city: customer.city,
        state: customer.state,
        country: customer.country,
        zip_code: customer.zip_code,
        credit_rating: customer.credit_rating,
        sales_rep_id: customer.sales_rep_id,
        comments: customer.comments,
      });
    } else if (open && !customer) {
      reset({ name: "" });
    }
  }, [open, customer, reset]);

  if (!open) return null;

  const onSubmit = async (data: CustomerFormData) => {
    // Convert nullable fields to undefined for the API
    const payload = {
      name: data.name,
      phone: data.phone ?? undefined,
      address: data.address ?? undefined,
      city: data.city ?? undefined,
      state: data.state ?? undefined,
      country: data.country ?? undefined,
      zip_code: data.zip_code ?? undefined,
      credit_rating: data.credit_rating ?? undefined,
      sales_rep_id: data.sales_rep_id ?? undefined,
      region_id: data.region_id ?? undefined,
      comments: data.comments ?? undefined,
    };
    if (isEdit) {
      await updateCustomer.mutateAsync(payload);
    } else {
      await createCustomer.mutateAsync(payload);
    }
    onClose();
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="fixed inset-0 bg-black/50" onClick={onClose} />
      <div className="relative bg-white rounded-lg shadow-lg p-6 max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <h2 className="text-lg font-semibold mb-4">
          {isEdit ? "Edit Customer" : "New Customer"}
        </h2>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="col-span-2">
              <label className="block text-sm font-medium mb-1">
                Name <span className="text-red-500">*</span>
              </label>
              <input
                {...register("name")}
                className="w-full border rounded-md px-3 py-2 text-sm"
                placeholder="Customer name"
              />
              {errors.name && (
                <p className="text-red-500 text-xs mt-1">{errors.name.message}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">Phone</label>
              <input
                {...register("phone")}
                className="w-full border rounded-md px-3 py-2 text-sm"
                placeholder="Phone number"
              />
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">Credit Rating</label>
              <select
                {...register("credit_rating")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              >
                <option value="">-- Select --</option>
                <option value="EXCELLENT">Excellent</option>
                <option value="GOOD">Good</option>
                <option value="POOR">Poor</option>
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">Sales Rep</label>
              <select
                {...register("sales_rep_id", { valueAsNumber: true })}
                className="w-full border rounded-md px-3 py-2 text-sm"
              >
                <option value="">-- Select --</option>
                {salesReps?.map((rep) => (
                  <option key={rep.id} value={rep.id}>
                    {rep.full_name}
                  </option>
                ))}
              </select>
            </div>

            <div className="col-span-2">
              <label className="block text-sm font-medium mb-1">Address</label>
              <input
                {...register("address")}
                className="w-full border rounded-md px-3 py-2 text-sm"
                placeholder="Street address"
              />
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">City</label>
              <input
                {...register("city")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              />
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">State</label>
              <input
                {...register("state")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              />
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">Country</label>
              <input
                {...register("country")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              />
            </div>

            <div>
              <label className="block text-sm font-medium mb-1">Zip Code</label>
              <input
                {...register("zip_code")}
                className="w-full border rounded-md px-3 py-2 text-sm"
              />
            </div>

            <div className="col-span-2">
              <label className="block text-sm font-medium mb-1">Comments</label>
              <textarea
                {...register("comments")}
                className="w-full border rounded-md px-3 py-2 text-sm"
                rows={3}
              />
            </div>
          </div>

          <div className="flex justify-end gap-2 pt-2">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 text-sm border rounded-md hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={isSubmitting}
              className="px-4 py-2 text-sm rounded-md text-white bg-primary hover:bg-primary/90 disabled:opacity-50"
            >
              {isSubmitting ? "Saving..." : isEdit ? "Update" : "Create"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
