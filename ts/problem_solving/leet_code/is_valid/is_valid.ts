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

export function isValidTwo(s: string) {
  const stack: string[] = []
  for (const c of s) {
    switch (c) {
      case ")":
        return stack.pop() === "("
      case "]":
        return stack.pop() === "["
      case "}":
        return stack.pop() === "{"
      default:
        stack.push(c)
    }
  }
  return stack.length === 0
}

console.log(isValidTwo("()")) // true
console.log(isValidTwo("()[]{}")) // true
console.log(isValidTwo("(]")) // false
