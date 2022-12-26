function snail(list: number[][]): number[] {
  const result: number[] = []
  for (let i = 0; i < list.length; i++) {
    const rows = list[i]
    // const colsLength = rows.length
    console.log(rows)
  }
  return result
}
const list = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
]
console.log(snail(list))
