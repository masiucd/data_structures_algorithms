export function moreZeros(s: string) {
  const charsWithMoreZeros = []
  for (const char of s) {
    const [zeros, ones] = countZerosAndOnesInBinary(
      charToBinary(char.charAt(0))
    )
    console.log({zeros, ones})
    if (zeros > ones) {
      charsWithMoreZeros.push(char)
    }
  }

  return [...new Set(charsWithMoreZeros)]
}

export function charToBinary(char: string) {
  return char.charCodeAt(0).toString(2)
}

export function countZerosAndOnesInBinary(binary: string): [number, number] {
  let zeros = 0
  let ones = 0
  for (const char of binary) {
    if (char === "0") {
      zeros++
    } else {
      ones++
    }
  }
  return [zeros, ones]
}
