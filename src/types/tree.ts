export type TreeNode = {
  id: number;
  title: string;
  parentId: number | null;
};

export type TreeDataType = TreeNode & {
  children: TreeDataType[];
};
