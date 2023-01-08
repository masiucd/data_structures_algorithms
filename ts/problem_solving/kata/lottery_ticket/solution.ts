function getCharCode(x: string): number {
  return x.charCodeAt(0)
}

type Ticket = [string, number]
export function solution(ticket: Ticket[], win: number): string {
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

export const solution2 = (ticket: Ticket[], win: number): string => {
  if (
    ticket.filter(a => a[0].split("").some(b => b.charCodeAt(0) == a[1]))
      .length >= win
  ) {
    return "Winner!"
  }
  return "Loser!"
}
