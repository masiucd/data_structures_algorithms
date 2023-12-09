interface BinarySearchTreeAble {
  insert(value: number): TreeNode | null;
  insertIter(value: number): TreeNode | null;
  find(value: number): TreeNode | null;
  remove(value: number): TreeNode | null;
  contains(value: number): boolean;
  size(): number;
  depth(): number;
  min(): number;
  max(): number;
  isEmpty(): boolean;
  toArray(): number[];
  toString(): string;
  bfs(): number[];
  dfs(order: "PRE" | "POST" | "IN"): number[];
}

class TreeNode {
  value: number;
  left: TreeNode | null;
  right: TreeNode | null;

  constructor(value: number) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

export class BinarySearchTree implements BinarySearchTreeAble {
  root: TreeNode | null;

  constructor() {
    this.root = null;
  }

  insertIter(value: number): TreeNode | null {
    if (this.root === null) {
      this.root = new TreeNode(value);
      return this.root;
    }
    let node = this.root;
    while (node !== null) {
      if (value < node.value) {
        if (node.left === null) {
          node.left = new TreeNode(value);
          return node.left;
        }
        node = node.left;
      } else {
        if (node.right === null) {
          node.right = new TreeNode(value);
          return node.right;
        }
        node = node.right;
      }
    }
    return null;
  }

  insert(value: number): TreeNode | null {
    if (this.root === null) {
      this.root = new TreeNode(value);
      return this.root;
    }
    return this.#insertHelper(this.root, value);
  }

  contains(value: number): boolean {
    return false;
  }
  max(): number {
    if (this.root === null) return -1;
    let current = this.root;
    while (current.right !== null) {
      current = current.right;
    }
    return current.value;
  }
  bfs(): number[] {
    if (this.root === null) return [];
    let result: number[] = [];
    let queue: TreeNode[] = []; // TODO make a real Q data structure
    queue.push(this.root);
    while (queue.length > 0) {
      let node = queue.shift();
      if (node) {
        result.push(node?.value);
        if (node.left) {
          queue.push(node.left);
        }
        if (node.right) {
          queue.push(node.right);
        }
      }
    }
    return result;
  }

  find(value: number): TreeNode | null {
    if (this.root === null) return null;
    return this.#findHelper(this.root, value);
  }

  remove(value: number): TreeNode | null {
    throw new Error("Method not implemented.");
  }
  size(): number {
    if (this.root === null) return 0;
    return this.#sizeHelper(this.root);
  }

  depth(): number {
    throw new Error("Method not implemented.");
  }
  min(): number {
    if (this.root === null) return -1;
    let current = this.root;
    while (current.left !== null) {
      current = current.left;
    }
    return current.value;
  }
  isEmpty(): boolean {
    return this.root === null;
  }
  toArray(): number[] {
    throw new Error("Method not implemented.");
  }
  toString(): string {
    throw new Error("Method not implemented.");
  }
  dfs(order: "PRE" | "POST" | "IN" = "PRE"): number[] {
    let result: number[] = [];
    switch (order) {
      case "PRE":
        this.#preOrder(this.root, result);
        break;
      case "POST":
        this.#postOrder(this.root, result);
        break;
      case "IN":
        this.#inOrder(this.root, result);
        break;
      default:
        break;
    }
    return result;
  }
  #inOrder(node: TreeNode | null, result: number[]) {
    if (node === null) return;
    this.#inOrder(node.left, result);
    result.push(node.value);
    this.#inOrder(node.right, result);
  }

  #preOrder(node: TreeNode | null, result: number[]) {
    if (node === null) return;
    result.push(node.value);
    this.#preOrder(node.left, result);
    this.#preOrder(node.right, result);
  }

  #postOrder(node: TreeNode | null, result: number[]) {
    if (node === null) return;
    this.#postOrder(node.left, result);
    this.#postOrder(node.right, result);
    result.push(node.value);
  }

  #sizeHelper(node: TreeNode | null): number {
    if (node === null) return 0;
    return 1 + this.#sizeHelper(node.left) + this.#sizeHelper(node.right);
  }
  #insertHelper(node: TreeNode | null, value: number) {
    if (node === null) {
      node = new TreeNode(value);
      return node;
    }
    if (value < node.value) {
      node.left = this.#insertHelper(node.left, value);
    } else {
      node.right = this.#insertHelper(node.right, value);
    }
    return node;
  }
  #findHelper(node: TreeNode | null, value: number): TreeNode | null {
    if (node === null) return null;
    if (node.value === value) return node;
    if (value < node.value) {
      return this.#findHelper(node.left, value);
    }
    return this.#findHelper(node.right, value);
  }
}
