import {assert, assertFalse} from "../../../dev_deps.ts"
import {isPalindrome} from "../palindrome.ts"

Deno.test(function isPalindromeItWorks() {
  assert(isPalindrome(""))
  assert(isPalindrome("racecar"))
  assert(isPalindrome("otto"))
})

Deno.test(function isPalindromeItFails() {
  assertFalse(isPalindrome("marcell"))
  assertFalse(isPalindrome("cwks"))
})
