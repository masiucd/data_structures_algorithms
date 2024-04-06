export function alternatingSums(weights: number[]): number[] {
  let left = [];
  let right = [];
  for (let [i, weight] of weights.entries()) {
    if (i % 2 === 0) {
      left.push(weight);
    } else {
      right.push(weight);
    }
  }
  return [left.reduce((a, b) => a + b, 0), right.reduce((a, b) => a + b, 0)];
}

export function alternatingSumsV2(weights: number[]): number[] {
  let leftSum = 0;
  let rightSum = 0;
  for (let [i, weight] of weights.entries()) {
    if (i % 2 === 0) {
      leftSum += weight;
    } else {
      rightSum += weight;
    }
  }
  return [leftSum, rightSum];
}
