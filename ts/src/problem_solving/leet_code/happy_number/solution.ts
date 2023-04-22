function isHappy(n: number): boolean {
  const store: Record<number, number> = {}
  while (n > 1) {
    if (store[n] > 1) return false
    const f = n
      .toString()
      .split("")
      .map(x => parseInt(x))
      .reduce((a, b) => a + b * b, 0)
    n = f
    if (store[n]) {
      store[n]++
    } else {
      store[n] = 1
    }
  }
  return n === 1
}

export {isHappy}
