"use client";

import { useState } from "react";
import { ChevronRight, ChevronDown } from "lucide-react";
import { cn } from "@/lib/utils";
import type { TreeNode } from "@/lib/types";

interface TreeViewProps {
  nodes: TreeNode[];
  onSelect: (node: TreeNode) => void;
  selectedId?: string;
}

export function TreeView({ nodes, onSelect, selectedId }: TreeViewProps) {
  return (
    <div className="text-sm">
      {nodes.map((node) => (
        <TreeNodeItem
          key={node.id}
          node={node}
          onSelect={onSelect}
          selectedId={selectedId}
          level={0}
        />
      ))}
    </div>
  );
}

function TreeNodeItem({
  node,
  onSelect,
  selectedId,
  level,
}: {
  node: TreeNode;
  onSelect: (node: TreeNode) => void;
  selectedId?: string;
  level: number;
}) {
  const [expanded, setExpanded] = useState(false);
  const hasChildren = node.children && node.children.length > 0;
  const isSelected = node.id === selectedId;

  return (
    <div>
      <div
        className={cn(
          "flex items-center gap-1 py-1 px-2 rounded cursor-pointer transition-colors",
          isSelected
            ? "bg-primary text-primary-foreground"
            : "hover:bg-accent"
        )}
        style={{ paddingLeft: `${level * 16 + 8}px` }}
        onClick={() => {
          if (hasChildren) setExpanded(!expanded);
          onSelect(node);
        }}
        onDoubleClick={() => onSelect(node)}
      >
        {hasChildren ? (
          expanded ? (
            <ChevronDown className="h-3 w-3 flex-shrink-0" />
          ) : (
            <ChevronRight className="h-3 w-3 flex-shrink-0" />
          )
        ) : (
          <span className="w-3" />
        )}
        <span className="truncate">{node.label}</span>
      </div>
      {expanded && hasChildren && (
        <div>
          {node.children!.map((child) => (
            <TreeNodeItem
              key={child.id}
              node={child}
              onSelect={onSelect}
              selectedId={selectedId}
              level={level + 1}
            />
          ))}
        </div>
      )}
    </div>
  );
}
