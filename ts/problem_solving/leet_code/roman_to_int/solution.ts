const roman = new Map()
roman.set("I", 1)
roman.set("V", 5)
roman.set("X", 10)
roman.set("L", 50)
roman.set("C", 100)
roman.set("D", 500)
roman.set("M", 1000)

function fromRomanToInt(s: string): number {
  let result = 0
  for (let i = 0; i < s.length; i++) {
    // check if is in bound
    if (i < s.length && roman.get(s[i]) < roman.get(s[i + 1])) {
      result -= roman.get(s[i])
    } else {
      result += roman.get(s[i])
    }
  }
  return result
}

export {fromRomanToInt}
