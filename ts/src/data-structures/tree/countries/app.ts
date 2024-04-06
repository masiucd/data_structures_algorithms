let data = [
  {
    id: 1,
    title: "Europe",
    parentId: null,
  },
  {
    id: 2,
    title: "Asia",
    parentId: null,
  },
  {
    id: 3,
    title: "Africa",
    parentId: null,
  },
  {
    id: 4,
    title: "Germany",
    parentId: 1,
  },
  {
    id: 5,
    title: "France",
    parentId: 1,
  },
  {
    id: 6,
    title: "China",
    parentId: 2,
  },
  {
    id: 7,
    title: "Japan",
    parentId: 2,
  },
  {
    id: 8,
    title: "Egypt",
    parentId: 3,
  },
  {
    id: 9,
    title: "Nigeria",
    parentId: 3,
  },
  {
    id: 10,
    title: "Berlin",
    parentId: 4,
  },
  {
    id: 11,
    title: "Munich",
    parentId: 4,
  },
  {
    id: 12,
    title: "Paris",
    parentId: 5,
  },
  {
    id: 13,
    title: "Lyon",
    parentId: 5,
  },
  {
    id: 14,
    title: "Beijing",
    parentId: 6,
  },
  {
    id: 15,
    title: "Shanghai",
    parentId: 6,
  },
  {
    id: 16,
    title: "Tokyo",
    parentId: 7,
  },
  {
    id: 17,
    title: "Osaka",
    parentId: 7,
  },
  {
    id: 18,
    title: "Cairo",
    parentId: 8,
  },
  {
    id: 19,
    title: "Alexandria",
    parentId: 8,
  },
  {
    id: 20,
    title: "Lagos",
    parentId: 9,
  },
  {
    id: 21,
    title: "Ibadan",
    parentId: 9,
  },
  {
    id: 22,
    title: "Thailand",
    parentId: 2,
  },
  {
    id: 23,
    title: "Poland",
    parentId: 1,
  },
  {
    id: 24,
    title: "Warszawa",
    parentId: 23,
  },
  {
    id: 25,
    title: "Krakow",
    parentId: 23,
  },
  {
    id: 26,
    title: "Mokotow",
    parentId: 24,
  },
];

type TreeNode = {
  id: number;
  title: string;
  parentId: number;
  children: TreeNode[];
};

function buildTree(treeData: typeof data, parentId: number | null): TreeNode[] {
  let children = treeData.filter((node) => node.parentId === parentId);
  return children.map((node) => ({
    ...node,
    open: false,
    children: buildTree(treeData, node.id),
  }));
}

function searchThroughTreeBFS(
  tree: TreeNode[],
  searchTerm: string
): TreeNode[] {
  let results: TreeNode[] = [];
  let queue: TreeNode[] = [...tree];
  while (queue.length > 0) {
    let node = queue.shift();
    if (node.title.toLowerCase().includes(searchTerm.toLowerCase())) {
      let n = {...node, open: true};
      results.push(n);
    }
    queue.push(...node.children);
  }
  return results;
}

let xs = buildTree(data, null);
console.log(searchThroughTreeBFS(xs, "Poland"));
