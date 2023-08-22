export function plusOne(digits: number[]): number[] {
  const result = [...digits].reverse()
  let carrier = 1
  let i = 0

  while (carrier > 0) {
    if (i < result.length) {
      if (result[i] === 9) {
        result[i] = 0
      } else {
        result[i]++
        carrier = 0
      }
    } else {
      result.push(1)
      carrier = 0
    }
    i++
  }
  result.reverse()
  return result
}

export function plusOne2(digits: number[]): number[] {
  let result = [...digits]
  for (let i = result.length - 1; i >= 0; i--) {
    if (result[i] < 9) {
      result[i]++
      return result
    }
    result[i] = 0
  }
  result = new Array(result.length + 1).fill(0)
  result[0] = 1
  return result
}
