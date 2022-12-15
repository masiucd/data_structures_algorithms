import {assertEquals} from "../../../dev_deps.ts"
import {reverse} from "../reverse_string.ts"

Deno.test(function reverseStringRecursiveItWorks() {
  assertEquals(reverse(""), "")
  assertEquals(reverse("marcell"), "llecram")
  assertEquals(reverse("1916"), "6191")
})
