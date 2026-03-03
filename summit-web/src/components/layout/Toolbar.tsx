"use client";

import { Save, Plus, Trash2, Search, Undo2, Filter } from "lucide-react";
import { cn } from "@/lib/utils";

interface ToolbarAction {
  icon: React.ElementType;
  label: string;
  onClick: () => void;
  variant?: "default" | "destructive";
  disabled?: boolean;
}

interface ToolbarProps {
  actions: ToolbarAction[];
}

export function Toolbar({ actions }: ToolbarProps) {
  return (
    <div className="flex items-center gap-1 border-b border-gray-200 bg-gray-50 px-4 py-2">
      {actions.map((action) => (
        <button
          key={action.label}
          onClick={action.onClick}
          disabled={action.disabled}
          title={action.label}
          className={cn(
            "flex items-center gap-1 rounded px-2 py-1 text-xs font-medium transition-colors",
            action.variant === "destructive"
              ? "text-red-600 hover:bg-red-50"
              : "text-gray-700 hover:bg-gray-200",
            action.disabled && "opacity-50 cursor-not-allowed"
          )}
        >
          <action.icon className="h-4 w-4" />
          {action.label}
        </button>
      ))}
    </div>
  );
}
