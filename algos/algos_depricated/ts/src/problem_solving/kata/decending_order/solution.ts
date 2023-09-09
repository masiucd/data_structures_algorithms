export function descendingOrder(n: number): number {
  return parseInt(n.toString().split("").sort().reverse().join(""))
}
