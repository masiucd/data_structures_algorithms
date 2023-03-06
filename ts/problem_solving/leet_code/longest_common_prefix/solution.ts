export function longestCommonPrefix(input: string[]): string {
  let result = ""
  for (let i = 0; i < input[0].length; i++) {
    for (const s of input) {
      if (i === s.length || s[i] !== input[0][i]) {
        return result
      }
    }
    result += input[0][i]
  }

  return result
}
