import {assertEquals} from "../../../dev_deps.ts"
import {descendingOrder} from "../decending_order/solution.ts"

Deno.test("descendingOrder it works!", () => {
  assertEquals(descendingOrder(0), 0)
  assertEquals(descendingOrder(1), 1)
  assertEquals(descendingOrder(123456789), 987654321)
})
