"use client";
import {type PropsWithChildren, useState} from "react";

import {cn} from "@/app/lib/css";
import type {TreeDataType} from "@/types/tree";

let marginLeft = new Map<number, string>()
  .set(0, "ml-0")
  .set(1, "ml-[20px]")
  .set(2, "ml-[40px]")
  .set(3, "ml-[60px]")
  .set(4, "ml-[80px]")
  .set(5, "ml-[100px]")
  .set(6, "ml-[120px]");

export function Row({
  node,
  level = 0,
  children,
}: PropsWithChildren<{
  node: TreeDataType;
  level?: number;
}>) {
  let [on, setOn] = useState(false);
  return (
    <li className={cn("flex", marginLeft.get(level))}>
      <button
        onClick={() => {
          setOn(!on);
        }}
      >
        {node.title}
      </button>
      {on && children}
    </li>
  );
}
