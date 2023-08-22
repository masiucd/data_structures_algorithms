import {assertEquals} from "../../../dev_deps.ts"
import {sum, sumDivideAndConcur} from "../math/sum.ts"

Deno.test(function sumItWorks() {
  assertEquals(sum([]), 0)
  assertEquals(sum([1, 2]), 3)
  assertEquals(sum([1, 2, 3]), 6)
})
Deno.test(function sumDivideAndConcurItWorks() {
  assertEquals(sumDivideAndConcur([]), 0)
  assertEquals(sumDivideAndConcur([1, 2]), 3)
  assertEquals(sumDivideAndConcur([1, 2, 3]), 6)
})
