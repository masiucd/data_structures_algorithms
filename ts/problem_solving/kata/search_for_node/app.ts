type Node = {
  id: number
  title: string
  parentId: number | null
  children: Node[]
}
export const treeData: Node[] = [
  {
    id: 1,
    title: "Europe",
    parentId: null,
    children: [
      {
        id: 4,
        title: "Germany",
        parentId: 1,
        children: [],
      },
      {
        id: 5,
        title: "France",
        parentId: 1,
        children: [],
      },
    ],
  },
  {
    id: 2,
    title: "Asia",
    parentId: null,
    children: [
      {
        id: 6,
        title: "China",
        parentId: 2,
        children: [],
      },
      {
        id: 7,
        title: "Japan",
        parentId: 2,
        children: [],
      },
    ],
  },
  {
    id: 3,
    title: "Africa",
    parentId: null,
    children: [
      {
        id: 8,
        title: "Egypt",
        parentId: 3,
        children: [],
      },
      {
        id: 9,
        title: "Kenya",
        parentId: 3,
        children: [],
      },
    ],
  },
]

function findNode(treeData: Node[], id: number, node: Node | null = null) {
  for (const n of treeData) {
    if (n.id === id) {
      return n
    }
    if (n.children.length > 0) {
      node = findNode(n.children, id, node)
      if (node) {
        return node
      }
    }
  }

  return null
}

function findNode2(treeData, id) {
  let foundNode = treeData.find((node) => node.id === id)
  if (foundNode) {
    return foundNode
  } else {
    for (let i = 0; i < treeData.length; i++) {
      if (treeData[i].children) {
        foundNode = findNode(treeData[i].children, id)
        if (foundNode) {
          return foundNode
        }
      }
    }
  }
  return null
}

console.log(findNode2(treeData, 5))
