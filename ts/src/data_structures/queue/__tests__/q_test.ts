import {assertEquals, assert} from "~/dev_deps.ts"
import {Queue} from "../q.ts"

Deno.test("Enqueue", () => {
  const queue = new Queue<number>()
  queue.enqueue(1)
  queue.enqueue(2)
  queue.enqueue(3)
  assertEquals(queue.print(), [1, 2, 3])
})

Deno.test("Dequeue", () => {
  const queue = new Queue<number>()
  queue.enqueue(1)
  queue.enqueue(2)
  queue.enqueue(3)
  assertEquals(queue.dequeue(), 1)
  assertEquals(queue.print(), [2, 3])

  assertEquals(queue.dequeue(), 2)
  assertEquals(queue.print(), [3])

  assertEquals(queue.dequeue(), 3)
  assertEquals(queue.print(), [])
  assert(queue.isEmpty())
})

Deno.test("Peek", () => {
  const queue = new Queue<number>()
  queue.enqueue(1)
  queue.enqueue(2)
  queue.enqueue(3)
  assertEquals(queue.peek(), 1)
  assertEquals(queue.print(), [1, 2, 3])
})
