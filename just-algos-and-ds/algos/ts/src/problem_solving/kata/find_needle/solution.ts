export function solution(haystack: string[]): string {
  let index = -1
  for (let i = 0; i < haystack.length; i++) {
    if (haystack[i] === "needle") {
      index = i
      break
    }
  }
  return index > -1 ? `found the needle at position ${index}` : ""
}
