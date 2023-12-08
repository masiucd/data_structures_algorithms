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

export class BinarySearchTree {
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

  contains(value: number): boolean {
    return false;
  }
  max(): number {
    return -1;
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
}
