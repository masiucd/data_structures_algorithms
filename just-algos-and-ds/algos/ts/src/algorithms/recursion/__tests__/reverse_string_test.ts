import {assertEquals} from "../../../dev_deps.ts"
import {reverse, reverse2} from "../reverse_string.ts"

Deno.test(function reverseStringRecursiveItWorks() {
  assertEquals(reverse(""), "")
  assertEquals(reverse("marcell"), "llecram")
  assertEquals(reverse("1916"), "6191")
})
Deno.test(function reverseStringRecursiveTwoItWorks() {
  assertEquals(reverse2(""), "")
  assertEquals(reverse2("marcell"), "llecram")
  assertEquals(reverse2("1916"), "6191")
})
