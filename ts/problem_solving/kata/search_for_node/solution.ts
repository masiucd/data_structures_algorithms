import {Node, treeData} from "./tree.ts"

export function findNode(
  treeData: Node[],
  id: number,
  node: Node | null = null
): null | Node {
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

export function findNode2(treeData: Node[], id: number): null | Node {
  let foundNode = treeData?.find((node) => node.id === id) ?? null
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

export function findNode3(tree: Node[], id: number): null | Node {
  for (const node of tree) {
    if (node.id === id) {
      return node
    } else if (node.children.length > 0) {
      const child = findNode3(node.children, id)
      if (child) return child
    }
  }
  return null
}

// console.log(findNode3(treeData, 10))
