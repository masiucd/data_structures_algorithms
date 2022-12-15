import {assertEquals} from "../../../dev_deps.ts"
import {sum} from "../math/sum.ts"

Deno.test(function sumItWorks() {
  assertEquals(sum([]), 0)
  assertEquals(sum([1, 2]), 3)
  assertEquals(sum([1, 2, 3]), 6)
})
