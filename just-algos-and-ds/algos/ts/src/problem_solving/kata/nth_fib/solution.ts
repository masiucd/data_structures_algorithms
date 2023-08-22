export function nthFibo(n: number) {
  if (n === 1) return 0
  const list = []
  for (let i = 0; i < n; i++) {
    list.push(fib(i))
  }
  console.log(list)
  return list.at(-1)
}

export function fib(n: number): number {
  let a = 0
  let b = 1
  let temp = 0
  for (let i = 0; i < n; i++) {
    temp = a
    a = b
    b = temp + b
  }
  return a
}
