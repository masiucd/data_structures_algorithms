export function solution(haystack: string, needle: string): number {
  const needleLen = needle.length
  for (let i = 0; i < haystack.length - needleLen + 1; i++) {
    if (haystack.substring(i, i + needleLen) === needle) {
      return i
    }
  }
  return -1
}
