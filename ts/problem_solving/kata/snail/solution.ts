export function snail(list: number[][]): number[] {
  const result: number[] = []
  while (list.length > 0) {
    //get the first row (first array in the array)
    result.push(...list.shift()!)
    for (let i = 0; i < list.length; i++) {
      //get the items at the end of each array (right side)
      result.push(list[i].pop()!)
    }
    //get the bottom row from end to front (bottom row reversed)
    result.push(...(list.pop() || []).reverse())
    for (let i = list.length - 1; i >= 0; i--) {
      //get the items at the beginning of each array (left side)
      result.push(list[i].shift()!)
    }
  }
  return result
}
const list = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
]
console.log(snail(list))
