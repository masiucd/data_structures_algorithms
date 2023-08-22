import {assertEquals} from "~/dev_deps.ts"

import {BinarySearchTree} from "../bst/bst.ts"

Deno.test("Bst Operations works", () => {
  const bst = new BinarySearchTree<number>()
  bst.insert(10)
  bst.insert(3)
  bst.insert(7)
  bst.insert(15)
  bst.insert(12)
  bst.insert(20)
  const inOrderList = bst.printInOrder()
  assertEquals(inOrderList, [3, 7, 10, 12, 15, 20])

  const preOrderList = bst.printPreOrder()
  assertEquals(preOrderList, [10, 3, 7, 15, 12, 20])

  const postOrderList = bst.printPostOrder()
  assertEquals(postOrderList, [7, 3, 12, 20, 15, 10])
})
