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
