import {TreeNode} from "@/types/tree";

export function buildAbsolutePath(
  treeData: readonly TreeNode[],
  id: number | null = null
): string {
  let node = treeData.find((item) => item.id === id);
  if (!node) {
    return "";
  }
  let parentAbsolutePath = buildAbsolutePath(treeData, node.parentId);
  return `${parentAbsolutePath}/${generateSlug(node.title)}`;
}

function generateSlug(title: string) {
  return title.toLowerCase().replace(/ /g, "-");
}
