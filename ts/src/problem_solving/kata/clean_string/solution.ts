export function solution(input: string): string {
  const stack: string[] = []

  for (const char of input) {
    if (char === "#") {
      if (stack.length > 0) {
        stack.pop()
      }
    } else {
      stack.push(char)
    }
  }
  return stack.join("")
}

console.log(solution("abc#d##c"))
console.log(solution("abc##d######"))
console.log(solution("#######"))
