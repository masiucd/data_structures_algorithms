export function isPalindrome(str: string): boolean {
  if (str.length === 0 || str.length === 1) {
    return true
  }
  const firstChar = str.at(0)
  const lastChar = str.at(-1)
  const middle = str.substring(1, str.length - 1)
  return firstChar === lastChar && isPalindrome(middle)
}

console.log(isPalindrome("abc")) // false;
console.log(isPalindrome("racecar")) // true
