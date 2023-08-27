export function sum(xs: number[]): number {
  if (xs.length === 0) return 0
  const head = xs[0]
  const tail = xs.slice(1)
  return head + sum(tail)
}

export function sumDivideAndConcur(xs: number[]) {
  // base case  --> Either an array containing zero numbers (where we return 0) or an array containing one number (where we return the number)
  // What argument is passed to the recursive function call?  --> Either the left half or the right half of the array of numbers.
  // How does this argument become closer to the base case?  --> The size of the array of numbers is halved each time, eventually becoming an array containing zero or one number.
  if (xs.length === 0) return 0
  if (xs.length === 1) return xs[0]
  const first = sumDivideAndConcur(xs.slice(0, Math.floor(xs.length / 2)))
  const second = sumDivideAndConcur(
    xs.slice(Math.floor(xs.length / 2), xs.length + 1)
  )
  return first + second
}
