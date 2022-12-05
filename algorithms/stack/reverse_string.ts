function reverseString(input: string): string {
  let reversed = ""
  const stack = []
  for (const char of input) {
    stack.push(char)
  }
  while (stack.length > 0) {
    const char = stack.pop()
    reversed += char
  }

  return reversed
}

console.log(reverseString("Hello"))
