import {assertEquals} from "~/dev_deps.ts"

import {rotate} from "../rotate/solution.ts"

Deno.test("rotate it works with fully list", () => {
  assertEquals(rotate("Hello"), ["elloH", "lloHe", "loHel", "oHell", "Hello"])
})

Deno.test("rotate it works with empty string", () => {
  assertEquals(rotate(""), [""])
})

Deno.test("rotate it works with one letter", () => {
  assertEquals(rotate("a"), ["a"])
})
