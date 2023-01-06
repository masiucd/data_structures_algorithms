export const brackets: Record<string, string> = {
  "}": "{",
  "]": "[",
  ")": "(",
}
export function isValid(s: string): boolean {
  const stack = []
  for (const c of s) {
    if (brackets[c]) {
      if (stack.length > 0 && stack.at(-1) === brackets[c]) {
        stack.pop()
      } else {
        return false
      }
    } else {
      stack.push(c)
    }
  }
  return stack.length === 0
}

console.log(isValid("()")) // true
console.log(isValid("()[]{}")) // true
console.log(isValid("(]")) // false
