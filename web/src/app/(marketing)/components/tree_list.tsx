import {cn} from "@/app/lib/css";
import {TreeDataType} from "@/types/tree";

import {Row} from "./row";

type Props = {
  tree: TreeDataType[];
  parentId?: number | null;
  level?: number;
};

const MARGIN_VALUES = Object.freeze(
  new Map([
    [0, "ml-0"],
    [1, "ml-[12px]"],
    [2, "ml-[22px]"],
    [3, "ml-[32px]"],
    [4, "ml-[42px]"],
    [5, "ml-[52px]"],
    [6, "ml-[62px]"],
    [7, "ml-[72px]"],
  ])
);

export function TreeList({tree, parentId = null, level = 0}: Props) {
  let nodes = tree.filter((n) => n.parentId === parentId);
  return (
    <ul className={cn("flex flex-col gap-2", MARGIN_VALUES.get(level))}>
      {nodes.map((node) => (
        <Row node={node} key={node.id}>
          <TreeList tree={node.children} parentId={node.id} level={level + 1} />
        </Row>
      ))}
    </ul>
  );
}
