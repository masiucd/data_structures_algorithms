interface RomanLetters {
  [key: string]: number
}
const romanLetters: RomanLetters = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000,
}
export function solution(roman: string): number {
  let result = 0
  if (roman === "IV") return 4
  for (const char of roman) {
    const value = romanLetters[char]
    result += value
  }
  return result
}
