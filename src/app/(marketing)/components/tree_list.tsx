import {cn} from "@/app/lib/css";
import {TreeDataType} from "@/types/tree";

import {Row} from "./row";

type Props = {
  tree: TreeDataType[];
  parentId?: number | null;
  level?: number;
};

const marginLeft = Object.freeze(
  new Map<number, string>()
    .set(0, "ml-0")
    .set(1, "ml-[12px]")
    .set(2, "ml-[22px]")
    .set(3, "ml-[32px]")
    .set(4, "ml-[42px]")
    .set(5, "ml-[52px]")
    .set(6, "ml-[62px]")
);

export function TreeList({tree, parentId = null, level = 0}: Props) {
  let nodes = tree.filter((n) => n.parentId === parentId);
  return (
    <ul className={cn("flex flex-col gap-2", marginLeft.get(level))}>
      {nodes.map((node) => (
        <Row node={node} key={node.id}>
          <TreeList tree={node.children} parentId={node.id} level={level + 1} />
        </Row>
      ))}
    </ul>
  );
}
