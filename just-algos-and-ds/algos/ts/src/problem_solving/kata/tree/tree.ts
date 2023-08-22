export type Node = {
  id: number
  title: string
  parentId: number | null
  children: Node[]
}
export const treeData: Node[] = [
  {
    id: 1,
    title: "Europe",
    parentId: null,
    children: [
      {
        id: 4,
        title: "Germany",
        parentId: 1,
        children: [],
      },
      {
        id: 5,
        title: "France",
        parentId: 1,
        children: [],
      },
    ],
  },
  {
    id: 2,
    title: "Asia",
    parentId: null,
    children: [
      {
        id: 6,
        title: "China",
        parentId: 2,
        children: [
          {
            id: 10,
            title: "Beijing",
            parentId: 6,
            children: [],
          },
        ],
      },
      {
        id: 7,
        title: "Japan",
        parentId: 2,
        children: [],
      },
    ],
  },
  {
    id: 3,
    title: "Africa",
    parentId: null,
    children: [
      {
        id: 8,
        title: "Egypt",
        parentId: 3,
        children: [],
      },
      {
        id: 9,
        title: "Kenya",
        parentId: 3,
        children: [],
      },
    ],
  },
]
