const arr1 = [
  {name: "Joe Brown", goals: 0, assists: 0, points: 0},
  {name: "Jim Bob", goals: 2, assists: 1, points: 3},
  {name: "Harry Styles", goals: 1, assists: 1, points: 2},
  {name: "Craig Mack", goals: 5, assists: 7, points: 12},
  {name: "Marcell Ciszek Druzynski 🔥", goals: 5, assists: 7, points: 12},
  {name: "DOG", bones: 100000, goals: 5, assists: 7, points: 12},
]
// prettier-ignore
const arr2 = [
  { name: 'Craig Mack', goals: 3, assists: 3, points: 6, ppg: 0, ppa: 0, pims: 0, },
  { name: 'Jim Bob', goals: 1, assists: 4, points: 5, ppg: 0, ppa: 1, pims: 0 },
  { name: 'Joe Brown', goals: 0, assists: 0, points: 0, ppg: 0, ppa: 0, pims: 0, },
  { name: 'Harry Styles', goals: 0, assists: 0, points: 0, ppg: 0, ppa: 0, pims: 0, },
];

export function addItUp(...items: any[]) {
  return items.flat().reduce((acc, item) => {
    const {name, ...rest} = item
    acc[name] = acc[name] || {...rest}
    Object.entries(rest).forEach(([key, value]) => {
      if (acc[name][key]) {
        acc[name][key] += value
      } else {
        acc[name][key] = value
      }
    })
    return acc
  }, {})
}

const result = addItUp(arr1, arr2)
console.log(result)
