const globeData = [
  {id: 1, title: "Europe", parentId: null},
  {id: 2, title: "Africa", parentId: null},
  {id: 3, title: "Asia", parentId: null},

  {id: 4, title: "France", parentId: 1},
  {id: 5, title: "Morocco", parentId: 2},
  {id: 6, title: "Poland", parentId: 1},
  {id: 7, title: "Thailand", parentId: 3},
  {id: 8, title: "Japan", parentId: 3},
  {id: 9, title: "Nigeria", parentId: 2},
  {id: 10, title: "Spain", parentId: 1},
  {id: 11, title: "Burma", parentId: 3},
  {id: 12, title: "Italy", parentId: 1},

  {id: 13, title: "Paris", parentId: 4},
  {id: 14, title: "Lyon", parentId: 4},
  {id: 15, title: "Marseille", parentId: 4},

  {id: 16, title: "Rabbat", parentId: 5},
  {id: 17, title: "Casablanca", parentId: 5},

  {id: 18, title: "Warszawa", parentId: 6},
  {id: 19, title: "Poznan", parentId: 6},
  {id: 20, title: "Krakow", parentId: 6},

  {id: 21, title: "Bangkok", parentId: 7},
  {id: 22, title: "Tokyo", parentId: 8},
  {id: 23, title: "Niger", parentId: 9},
  {id: 24, title: "Barcelona", parentId: 10},
  {id: 25, title: "Madrid", parentId: 10},
];

interface GlobeNode {
  id: number;
  title: string;
  parentId: number | null;
}

interface Result extends GlobeNode {
  children: Result[];
}

function constructTree(list: GlobeNode[], parentId: number | null): Result[] {
  return list
    .filter((x) => x.parentId === parentId)
    .map((child) => ({...child, children: constructTree(list, child.id)}));
}

function constructTreeTwo(
  list: GlobeNode[],
  parentId: number | null
): Result[] {
  const parents = list.filter((x) => x.parentId === parentId);
  const xs = [];
  for (const item of parents) {
    xs.push({
      ...item,
      children: constructTreeTwo(list, item.id),
    });
  }
  return xs;
}

// const xs = constructTreeTwo(globeData, null);
// console.log("xs", xs);
