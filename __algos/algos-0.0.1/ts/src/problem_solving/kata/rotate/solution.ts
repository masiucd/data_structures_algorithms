function rotate(str: string) {
  if (str.length === 0) {
    return [""]
  }
  const res = []
  for (let i = 0; i < str.length; i++) {
    const word = str.slice(i) + str.slice(0, i)
    res.push(word)
  }
  const first = res.shift()
  res.push(first)
  return res
}

export {rotate}
