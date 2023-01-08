function getCharCode(x: string): number {
  return x.charCodeAt(0)
}

type Ticket = [string, number]
export function bingo(ticket: Ticket[], win: number): string {
  let count = 0
  for (const [chars, target] of ticket) {
    for (const c of chars) {
      if (getCharCode(c) === target) {
        count++
        break
      }
    }
  }
  return count >= win ? "Winner!" : "Loser!"
}
