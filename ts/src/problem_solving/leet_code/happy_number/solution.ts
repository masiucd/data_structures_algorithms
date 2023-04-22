function isHappy(n: number): boolean {
  const store = new Set()
  while (!store.has(n)) {
    store.add(n)
    n = n
      .toString()
      .split("")
      .map(x => parseInt(x))
      .reduce((a, b) => a + b * b, 0)
    if (n === 1) {
      return true
    }
    if (store.has(n)) return false
  }
  return false
}

function isHappy2(n: number): boolean {
  const store = new Set()
  while (!store.has(n)) {
    store.add(n)
    n = numOfSquare(n)
    if (n === 1) return true
  }
  return false
}

function numOfSquare(n: number) {
  let res = 0
  while (n > 0) {
    let num = n % 10
    num = num ** 2
    res += num
    n = Math.floor(n / 10)
  }
  return res
}

export {isHappy, isHappy2, numOfSquare}
