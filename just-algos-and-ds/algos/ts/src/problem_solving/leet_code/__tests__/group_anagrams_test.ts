import {assertEquals} from "~/dev_deps.ts"
import {groupAnagrams} from "../group_anagrams/group_anagrams.ts"

Deno.test("groupAnagrams it works with fully list", () => {
  assertEquals(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]), [
    ["eat", "tea", "ate"],
    ["tan", "nat"],
    ["bat"],
  ])
})

Deno.test("groupAnagrams it works with empty string", () => {
  assertEquals(groupAnagrams([""]), [[""]])
})
