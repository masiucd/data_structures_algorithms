import {assertEquals} from "~/dev_deps.ts"
import {LinkedList} from "../list.ts"

Deno.test("Prepend to list", () => {
  const list = new LinkedList<number>()
  list.prepend(10)
  list.prepend(3)
  list.prepend(7)
  assertEquals(list.print(), [7, 3, 10])
  assertEquals(list.getHead(), 7)
  assertEquals(list.getSize(), 3)
  assertEquals(list.getTail(), 10)
})

Deno.test("Append to list", () => {
  const list = new LinkedList<number>()
  list.append(10)
  list.append(3)
  list.append(7)
  assertEquals(list.print(), [10, 3, 7])
  assertEquals(list.getHead(), 10)
  assertEquals(list.getSize(), 3)
  assertEquals(list.getTail(), 7)
})

Deno.test("Get specific node from the list", () => {
  const list = new LinkedList<number>()
  list.append(10)
  list.prepend(2)
  list.append(3)
  list.prepend(100)
  list.append(7)
  list.prepend(22)
  // [22,100,2,10,3,7]
  assertEquals(list.print(), [22, 100, 2, 10, 3, 7])
  // Get the head
  assertEquals(list.get(0), 22)
  // Get the tail
  assertEquals(list.get(5), 7)
  // Get the middle
  assertEquals(list.get(2), 2)
})

Deno.test("Delete node from the list", () => {
  const list = new LinkedList<number>()
  list.append(10)
  list.prepend(2)
  list.append(3)
  list.prepend(100)
  list.append(7)
  list.prepend(22)
  // [22,100,2,10,3,7]
  assertEquals(list.print(), [22, 100, 2, 10, 3, 7])
  // Delete the head
  list.delete(0)
  assertEquals(list.print(), [100, 2, 10, 3, 7])
  // Delete the tail
  list.delete(4)
  assertEquals(list.print(), [100, 2, 10, 3])
  // Delete the middle
  list.delete(2)
  assertEquals(list.print(), [100, 2, 3])
  // Delete the last
  list.delete(2)
  assertEquals(list.print(), [100, 2])
  // Delete the first
  list.delete(0)
  assertEquals(list.print(), [2])
  // Delete the only
  list.delete(0)
  assertEquals(list.print(), [])
  assertEquals(list.getSize(), 0)
  assertEquals(list.getHead(), null)
  assertEquals(list.getTail(), null)
})

Deno.test("Insert node to the list", () => {
  const list = new LinkedList<number>()
  list.append(10)
  list.prepend(2)
  list.append(3)
  list.prepend(100)
  list.append(7)
  list.prepend(22)
  // [22,100,2,10,3,7]
  assertEquals(list.print(), [22, 100, 2, 10, 3, 7])
  // Insert at head
  list.insert(0, 1)
  assertEquals(list.print(), [1, 22, 100, 2, 10, 3, 7])
  // Insert at tail
  list.insert(7, 9)
  assertEquals(list.print(), [1, 22, 100, 2, 10, 3, 7, 9])
  // Insert at middle
  list.insert(4, 8)
  assertEquals(list.print(), [1, 22, 100, 2, 8, 10, 3, 7, 9])
  // Insert at last
  list.insert(9, 11)
  assertEquals(list.print(), [1, 22, 100, 2, 8, 10, 3, 7, 9, 11])
  // Insert at first
  list.insert(0, 0)
  assertEquals(list.print(), [0, 1, 22, 100, 2, 8, 10, 3, 7, 9, 11])
})

Deno.test("Reverse list", () => {
  const list = new LinkedList<number>()
  list.append(10)
  list.prepend(2)
  list.append(3)
  list.prepend(100)
  list.append(7)
  list.prepend(22)
  assertEquals(list.print(), [22, 100, 2, 10, 3, 7])
  list.reverse()
  assertEquals(list.print(), [7, 3, 10, 2, 100, 22])
})
