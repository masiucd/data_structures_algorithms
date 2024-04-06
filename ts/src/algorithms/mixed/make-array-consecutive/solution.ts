export function makeArrayConsecutive(statues: number[]): number {
  let xs = statues.toSorted((a, b) => a - b);
  let res = 0;
  for (let i = 0; i < xs.length - 1; i++) {
    let diff = xs[i + 1] - xs[i];
    if (diff !== 1) {
      res += diff - 1;
    }
  }
  return res;
}
