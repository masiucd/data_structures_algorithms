type Node = {
  id: number
  name: string
  parentId: number | null
}
const flatData: Node[] = [
  {id: 1, name: "Europe", parentId: null},
  {id: 2, name: "Asia", parentId: null},
  {id: 3, name: "Africa", parentId: null},
  {id: 4, name: "Germany", parentId: 1},
  {id: 5, name: "France", parentId: 1},
  {id: 6, name: "China", parentId: 2},
  {id: 7, name: "Japan", parentId: 2},
  {id: 8, name: "Egypt", parentId: 3},
  {id: 9, name: "Kenya", parentId: 3},
  {id: 10, name: "Beijing", parentId: 6},
  {id: 11, name: "Tokyo", parentId: 7},
  {id: 12, name: "Cairo", parentId: 8},
  {id: 13, name: "Nairobi", parentId: 9},
  {id: 14, name: "Paris", parentId: 5},
  {id: 15, name: "Berlin", parentId: 4},
  {id: 16, name: "Munich", parentId: 4},
  {id: 17, name: "Dortmund", parentId: 4},
  {id: 18, name: "Dresden", parentId: 4},
  {id: 19, name: "Leipzig", parentId: 4},
  {id: 20, name: "Hamburg", parentId: 4},
  {id: 21, name: "Stuttgart", parentId: 4},
  {id: 22, name: "Marseille", parentId: 5},
  {id: 23, name: "Lyon", parentId: 5},
  {id: 24, name: "Toulouse", parentId: 5},
  {id: 25, name: "Nice", parentId: 5},
  {id: 26, name: "Nantes", parentId: 5},
  {id: 27, name: "Strasbourg", parentId: 5},
  {id: 28, name: "Düsseldorf", parentId: 4},
  {id: 29, name: "Essen", parentId: 4},
  {id: 30, name: "Bremen", parentId: 4},
  {id: 31, name: "Duisburg", parentId: 4},
  {id: 32, name: "Bochum", parentId: 4},
  {id: 33, name: "Thailand", parentId: 2},
  {id: 34, name: "Bangkok", parentId: 33},
  {id: 35, name: "Phuket", parentId: 33},
  {id: 36, name: "Pattaya", parentId: 33},
  {id: 37, name: "Chiang Mai", parentId: 33},
  {id: 38, name: "Krabi", parentId: 33},
]

function buildTree(
  data: Node[],
  parentId: number | null = null
): Node[] | null {
  const nodes = data.filter(x => x.parentId === parentId)
  if (nodes.length === 0) {
    return null
  }
  return nodes.map(x => ({...x, children: buildTree(data, x.id)}))
}

const res = buildTree(flatData)
console.log(JSON.stringify(res, null, 2))
