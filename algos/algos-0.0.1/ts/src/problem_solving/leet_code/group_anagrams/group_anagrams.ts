export function groupAnagrams(input: string[]): string[][] {
  const result: string[][] = []
  const store = new Map<string, string[]>()

  for (const word of input) {
    const sorted = word.split("").sort().join("")
    if (!store.has(sorted)) {
      store.set(sorted, [word])
    } else {
      store.get(sorted)!.push(word)
    }
  }

  for (const [_, value] of store) {
    result.push(value)
  }
  return result
}
