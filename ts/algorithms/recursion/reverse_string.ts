export function reverse(input: string): string {
  if (input.length === 0) {
    return ""
  }
  const head = input.charAt(0)
  const tail = input.slice(1)
  // all the chars in the string except the first and add it to the end!
  return reverse(tail) + head
}
