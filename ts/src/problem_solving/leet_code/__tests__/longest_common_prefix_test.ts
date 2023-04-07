import {assertEquals} from "~/dev_deps.ts"
import {longestCommonPrefix} from "../longest_common_prefix/solution.ts"

Deno.test("longestCommonPrefix it works with fully list", () => {
  assertEquals(longestCommonPrefix(["flower", "flow", "flight"]), "fl")
})

Deno.test("longestCommonPrefix it works with empty string", () => {
  assertEquals(longestCommonPrefix([""]), "")
})
