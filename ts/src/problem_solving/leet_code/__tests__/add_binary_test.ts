import {assertEquals} from "~/dev_deps.ts"
import {addBinary} from "../add_binary/solution.ts"

Deno.test("addBinary it works", () => {
  assertEquals(addBinary("11", "1"), "100")
})

Deno.test("addBinary it works", () => {
  assertEquals(addBinary("1010", "1011"), "10101")
})
