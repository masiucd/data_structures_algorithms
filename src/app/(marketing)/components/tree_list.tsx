import {PropsWithChildren} from "react";

import {TreeDataType} from "@/types/tree";

import {Row} from "./row";

type Props = {
  tree: TreeDataType[];
  parentId?: number | null;
  level?: number;
};

export function TreeList({
  tree,
  // children,
  parentId = null,
  level = 0,
}: PropsWithChildren<Props>) {
  let nodes = tree.filter((n) => n.parentId === parentId);
  return (
    <ul className="flex flex-col gap-1">
      {nodes.map((node) => {
        return (
          <Row node={node} level={level} key={node.id}>
            <TreeList
              tree={node.children}
              parentId={node.id}
              level={level + 1}
            />
          </Row>
        );
      })}
    </ul>
  );
}
