import {zipWith} from "../zip_with/solution.ts"
import {assertEquals} from "../../../dev_deps.ts"

Deno.test("zipWith it works +!", () => {
  const res = zipWith((a, b) => a + b, [1, 2, 3], [4, 5, 6])
  assertEquals(res, [5, 7, 9])
})

Deno.test("zipWith it works *!", () => {
  const res = zipWith((a, b) => a * b, [1, 2, 3], [4, 5, 6, 20])
  assertEquals(res, [4, 10, 18])
})

Deno.test("zipWith it works -!", () => {
  const res = zipWith((a, b) => a - b, [1, 2, 3], [4, 5, 6, 20])
  assertEquals(res, [-3, -3, -3])
})
