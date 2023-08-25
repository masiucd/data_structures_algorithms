"use client";
import {motion} from "framer-motion";
import Link from "next/link";
import {type PropsWithChildren, useState} from "react";

import {cn} from "@/app/lib/css";
import {Icons} from "@/components/icons";
import type {TreeDataType} from "@/types/tree";

function renderIcon(node: TreeDataType, on: boolean) {
  if (node.children.length > 0 && !on) {
    return <Icons.FilePlus />;
  }
  if (node.children.length === 0) {
    return <Icons.Code />;
  }
  return <Icons.FileMinus />;
}

export function Row({
  node,
  children,
}: PropsWithChildren<{
  node: TreeDataType;
}>) {
  let [on, setOn] = useState(false);

  return (
    <motion.li
      initial={{opacity: 0.5, height: 0}}
      animate={{
        opacity: 1,
        height: "auto",
        transition: {
          duration: 1.2,
          type: "spring",
        },
      }}
    >
      <div className="flex gap-1">
        {node.children.length > 0 ? (
          <button
            className={cn(
              "flex items-center disabled:opacity-60",
              node.children.length === 0 && "cursor-not-allowed text-blue-950"
            )}
            onClick={() => {
              setOn(!on);
            }}
          >
            <span>{renderIcon(node, on)}</span>
            <span>{node.title}</span>
          </button>
        ) : (
          <div>
            <Link
              href={node.href}
              className={cn(
                "flex items-center gap-1",
                node.children.length > 0 && "pointer-events-none"
              )}
            >
              <span>{renderIcon(node, on)}</span>
              <span>{node.title}</span>
            </Link>
          </div>
        )}
      </div>
      {on && children}
    </motion.li>
  );
}
