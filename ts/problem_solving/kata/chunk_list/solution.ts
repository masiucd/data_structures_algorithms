function chunkList(list: number[], k: number): number[][] {
  const result = []
  for (let i = 0; i < list.length; i += k) {
    const xs = list.slice(i, i + k)
    result.push(xs)
  }
  return result
}

console.log(chunkList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 3))
