class Node<T> {
  value: T
  left: Node<T> | null
  right: Node<T> | null
  constructor(value: T) {
    this.value = value
    this.left = null
    this.right = null
  }
}

interface BstTreeAble<T> {
  insert(value: T): void
  search(value: T): boolean
  delete(value: T): void
  printInOrder(): T[]
  printPreOrder(): T[]
  printPostOrder(): T[]
  dfs(value: T): boolean
  bfs(value: T): boolean
  printDFS(): T[]
  printBFS(): T[]
}

export class BinarySearchTree<T> implements BstTreeAble<T> {
  private root: Node<T> | null = null

  constructor() {
    this.root = null
  }

  insert(value: T): void {
    this.root = this.insertHelper(this.root, value)
  }

  private insertHelper(node: Node<T> | null, value: T): Node<T> {
    if (node === null) {
      return new Node(value)
    } else if (value < node.value) {
      node.left = this.insertHelper(node.left, value)
    } else {
      node.right = this.insertHelper(node.right, value)
    }
    return node
  }

  search(value: T): boolean {
    throw new Error("Method not implemented.")
  }
  delete(value: T): void {
    throw new Error("Method not implemented.")
  }
  printInOrder(): T[] {
    const result: T[] = []
    const traverse = (node: Node<T> | null) => {
      if (node === null) return
      traverse(node.left)
      result.push(node.value)
      traverse(node.right)
    }
    traverse(this.root)
    return result
  }
  printPreOrder(): T[] {
    const result: T[] = []
    const traverse = (node: Node<T> | null) => {
      if (node === null) return
      result.push(node.value)
      traverse(node.left)
      traverse(node.right)
    }
    traverse(this.root)
    return result
  }
  printPostOrder(): T[] {
    const result: T[] = []
    const traverse = (node: Node<T> | null) => {
      if (node === null) return
      traverse(node.left)
      traverse(node.right)
      result.push(node.value)
    }
    traverse(this.root)
    return result
  }
  dfs(value: T): boolean {
    throw new Error("Method not implemented.")
  }
  bfs(value: T): boolean {
    throw new Error("Method not implemented.")
  }

  printDFS(): T[] {
    throw new Error("Method not implemented.")
  }
  printBFS(): T[] {
    throw new Error("Method not implemented.")
  }
}

const bst = new BinarySearchTree<number>()
bst.insert(10)
bst.insert(3)
bst.insert(22)
console.log(bst)
