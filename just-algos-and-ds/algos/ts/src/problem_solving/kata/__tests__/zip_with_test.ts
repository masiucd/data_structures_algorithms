import {zipWith, zipWith2} from "../zip_with/solution.ts"
import {assertEquals} from "../../../dev_deps.ts"

Deno.test("zipWith it works +!", () => {
  let res = zipWith((a, b) => a + b, [1, 2, 3], [4, 5, 6])
  assertEquals(res, [5, 7, 9])
  res = zipWith2((a, b) => a + b, [1, 2, 3], [4, 5, 6])
  assertEquals(res, [5, 7, 9])
})

Deno.test("zipWith it works *!", () => {
  let res = zipWith((a, b) => a * b, [1, 2, 3], [4, 5, 6, 20])
  assertEquals(res, [4, 10, 18])
  res = zipWith2((a, b) => a * b, [1, 2, 3], [4, 5, 6, 20])
  assertEquals(res, [4, 10, 18])
})

Deno.test("zipWith it works -!", () => {
  let res = zipWith((a, b) => a - b, [1, 2, 3], [4, 5, 6, 20])
  assertEquals(res, [-3, -3, -3])
  res = zipWith2((a, b) => a - b, [1, 2, 3], [4, 5, 6, 20])
  assertEquals(res, [-3, -3, -3])
})
