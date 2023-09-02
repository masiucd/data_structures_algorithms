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
