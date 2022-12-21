const root = {
  key: "A",
  children: [
    {key: "B", children: [{key: "D", children: []}]},
    {
      key: "C",
      children: [
        {
          key: "E",
          children: [
            {key: "F", children: []},
            {key: "G", children: [{key: "H", children: []}]},
          ],
        },
      ],
    },
  ],
}
function recursiveSum(xs: number[]) {
  if (xs.length === 0) {
    return 0
  }
  const head = xs[0]
  const tail = xs.slice(1)
  return head + recursiveSum(tail)
}

console.log(recursiveSum([1, 2, 3, 4]))
