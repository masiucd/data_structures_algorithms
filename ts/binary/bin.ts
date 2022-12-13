// Operators
// AND &
//  OR |
//  XOR ^
//  NOT ~
//  SHIFT LEFT <<
//  SHIFT RIGHT >>

let n = 10
// console.log(n)
// console.log(n.toString(2))
let five = 5
// console.log(n >> 1, "=", five.toString(2))

function binaryNumbers(from: number, to: number) {
  const result: Map<string, string> = new Map()
  for (let i = from; i <= to; i++) {
    result.set(i.toString(10), i.toString(2))
  }
  return result
}

// 1001
// 1010
// 1000
// console.log(10 & 9) // 1000 --> 8
