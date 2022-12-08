const all = [
  {
    id: 797331,
    code: "3O",
    name: "Air Arabia Maroc",
  },
  {
    id: 384,
    code: "DY",
    name: "Norwegian",
  },
]
const selected = [
  {
    id: 797331,
    code: "3O",
    name: "Air Arabia Maroc",
  },
]

type Item = {id: number; code: string; name: string}
// we want to merge both list into one list but where we prefix each object if  is selected or not
type Vc = {id: number; code: string; name: string; isSelected: boolean}
function merged(all: Item[], selected: Item[]): Vc[] {
  const result: Vc[] = []
  for (const vc of all) {
    if (selected.find((item) => item.id === vc.id)) {
      result.push({...vc, isSelected: true})
    } else {
      result.push({...vc, isSelected: false})
    }
  }
  console.log(result)
  return result
}

// merged(all, selected)

function getLabeledValidatingCarriers(all: Item[], selected: Item[]): Vc[] {
  const xs = all.map((item) =>
    selected.find((s) => item.id === s.id)
      ? {...item, isSelected: true}
      : {...item, isSelected: false}
  )
  console.log(xs)

  return xs
}
getLabeledValidatingCarriers(all, selected)
