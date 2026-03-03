"use client";

import { LogOut, User } from "lucide-react";
import { useAuth } from "@/providers/AuthProvider";

export function Header() {
  const { user, logout } = useAuth();

  return (
    <header className="h-14 border-b border-gray-200 bg-white flex items-center justify-between px-6">
      <div />
      <div className="flex items-center gap-4">
        <span className="text-sm text-muted-foreground flex items-center gap-2">
          <User className="h-4 w-4" />
          {user?.email}
        </span>
        <button
          onClick={logout}
          className="text-sm text-muted-foreground hover:text-foreground flex items-center gap-1"
        >
          <LogOut className="h-4 w-4" />
          Sign out
        </button>
      </div>
    </header>
  );
}
