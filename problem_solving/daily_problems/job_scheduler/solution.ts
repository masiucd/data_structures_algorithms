type Fn = (name: string) => string
const jobSchedular = (fn: Fn, n: number): string => {
  setTimeout(() => {
    fn()
  }, n * 1000)
}

const greet = (name: string) => () => {
  console.log(`hello ${name}`)
}
const x = jobSchedular(greet("Bob"), 2)
