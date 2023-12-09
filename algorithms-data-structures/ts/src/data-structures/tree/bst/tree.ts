// - `insert(value)`: inserts a new value in the tree, recursively
// - `insertIter(value)`: inserts a new value in the tree, iteratively
// - `find(value)`: returns the node with the given value
// - `remove(value)`: removes the node with the given value
// - `contains(value)`: returns true if the value exists in the tree, otherwise false
// - `size()`: returns the number of nodes in the tree
// - `depth()`: returns the maximum depth of the tree
// - `min()`: returns the minimum value in the tree
// - `max()`: returns the maximum value in the tree
// - `isEmpty()`: returns true if the tree is empty, otherwise false
// - `toArray()`: returns an array with all the values in the tree, in order
// - `toString()`: returns a string representation of the tree
// - `bfs()`: returns an array with all the values in the tree, in breadth-first order
// - `dfs()`: returns an array with all the values in the tree, in depth-first order
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
  dfs(): number[];
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
  find(value: number): TreeNode | null {
    throw new Error("Method not implemented.");
  }
  remove(value: number): TreeNode | null {
    throw new Error("Method not implemented.");
  }
  size(): number {
    throw new Error("Method not implemented.");
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
  dfs(): number[] {
    throw new Error("Method not implemented.");
  }
}
