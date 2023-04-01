import {assertEquals} from "../../../dev_deps.ts"
import {nthFibo, fib} from "../nth_fib/solution.ts"

Deno.test("nthFibo it works", () => {
  assertEquals(nthFibo(3), 1)
  assertEquals(nthFibo(1), 0)
  assertEquals(nthFibo(2), 1)
  assertEquals(nthFibo(4), 2)
})

Deno.test("fib it works", () => {
  assertEquals(fib(0), 0)
  assertEquals(fib(1), 1)
  assertEquals(fib(2), 1)
  assertEquals(fib(3), 2)
  assertEquals(fib(4), 3)
  assertEquals(fib(5), 5)
  assertEquals(fib(6), 8)
  assertEquals(fib(7), 13)
})
