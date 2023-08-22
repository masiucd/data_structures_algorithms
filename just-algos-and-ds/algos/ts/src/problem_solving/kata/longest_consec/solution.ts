export function solution(input: string[], k: number): string {
  let longest = ""
  if (k > 0 && input.length >= k) {
    for (let i = 0; i < input.length - k + 1; i++) {
      const s = input.slice(i, i + k).join("")
      if (s.length > longest.length) {
        longest = s
      }
    }
  }
  return longest
}
