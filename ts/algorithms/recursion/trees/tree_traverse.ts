const root = {
  data: "A",
  children: [
    {data: "B", children: [{data: "D", children: []}]},
    {
      data: "C",
      children: [
        {
          data: "E",
          children: [
            {data: "G", children: []},
            {data: "H", children: []},
          ],
        },
        {data: "F", children: []},
      ],
    },
  ],
}

export function preorderTraverse(node: typeof root, nodes: string[]) {
  nodes.push(node["data"])
  if (node["children"].length > 0) {
    // Recursive case
    for (let i = 0; i < node["children"].length; i++) {
      preorderTraverse(node["children"][i], nodes) // Traverse child nodes.
    }
    return nodes
  }
}

export function postorderTraverse(node: typeof root, nodes: string[]) {
  for (let i = 0; i < node["children"].length; i++) {
    // RECURSIVE CASE
    postorderTraverse(node["children"][i], nodes) // Traverse child nodes.
  }
  nodes.push(node["data"])
  // BASE CASE
  return nodes
}

export function inorderTraverse(node: typeof root, nodes: string[]) {
  if (node["children"].length >= 1) {
    // RECURSIVE CASE
    inorderTraverse(node["children"][0], nodes) // Traverse the left child.
  }
  nodes.push(node["data"])
  if (node["children"].length >= 2) {
    // RECURSIVE CASE
    inorderTraverse(node["children"][1], nodes) // Traverse the right child.
  }
  // BASE CASE
  return nodes
}

// const xs = inorderTraverse(root, [])
