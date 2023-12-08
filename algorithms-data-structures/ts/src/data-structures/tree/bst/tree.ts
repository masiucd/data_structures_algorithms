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

  insert(value: number): boolean {
    let node = new TreeNode(value);
    if (this.root === null) {
      this.root = node;
      return true;
    }
    let current = this.root;
    while (true) {
      if (value === current.value) return false;
      if (value < current.value) {
        // go left
        if (current.left && value < current.value) {
          current = current.left;
        } else {
          current.left = node;
        }
      } else {
        // go right
        if (current.right && value < current.value) {
          current = current.right;
        } else {
          current.right = node;
        }
      }
      return true;
    }
  }

  contains(value: number): boolean {
    return false;
  }
  max(): number {
    return -1;
  }
}
