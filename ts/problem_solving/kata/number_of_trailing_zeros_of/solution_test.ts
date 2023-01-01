import {assertEquals} from "../../../dev_deps.ts"
import {zeros} from "./solution.ts"

Deno.test(function zerosWorks() {
  assertEquals(zeros(0), 0)
  assertEquals(zeros(6), 1)
  assertEquals(zeros(30), 7)
  assertEquals(zeros(100), 24)
  assertEquals(zeros(1000), 249)
})
