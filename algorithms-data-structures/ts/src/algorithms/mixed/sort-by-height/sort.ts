export function sortByHeight(heights: number[]) {
  let sorted = heights.filter((n) => n !== -1).sort((a, b) => a - b);
  let i = 0;
  return heights.map((n) => (n === -1 ? -1 : sorted[i++]));
}
