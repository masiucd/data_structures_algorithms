import {assertEquals} from "https://deno.land/std@0.153.0/testing/asserts.ts"
import {Stack} from "../stack.ts"

Deno.test("Push to stack", () => {
  const stack = new Stack<number>()
  stack.push(10)
  stack.push(3)
  stack.push(7)
  assertEquals(stack.print(), [10, 3, 7])
  assertEquals(stack.getSize(), 3)
})

Deno.test("Pop from stack", () => {
  const stack = new Stack<number>()
  stack.push(10)
  stack.push(3)
  stack.push(7)
  assertEquals(stack.pop(), 7)
  assertEquals(stack.print(), [10, 3])
  assertEquals(stack.getSize(), 2)
})

Deno.test("Peek from stack", () => {
  const stack = new Stack<number>()
  stack.push(10)
  stack.push(3)
  stack.push(7)
  assertEquals(stack.peek(), 7)
  assertEquals(stack.print(), [10, 3, 7])
  assertEquals(stack.getSize(), 3)
})

Deno.test("Check if stack is empty", () => {
  const stack = new Stack<number>()
  assertEquals(stack.isEmpty(), true)
  stack.push(10)
  assertEquals(stack.isEmpty(), false)
  stack.push(3)
  assertEquals(stack.isEmpty(), false)
  stack.push(7)
  assertEquals(stack.isEmpty(), false)
  stack.pop()
  assertEquals(stack.isEmpty(), false)
  stack.pop()
  assertEquals(stack.isEmpty(), false)
  stack.pop()
  assertEquals(stack.isEmpty(), true)
})
