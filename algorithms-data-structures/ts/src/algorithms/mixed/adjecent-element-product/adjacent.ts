export function adjacentElementsProduct(inputArray: number[]): number {
  let max = inputArray[0] * inputArray[1];
  for (let i = 1; i < inputArray.length - 1; i++) {
    let current = inputArray[i];
    let next = inputArray[i + 1];
    let product = current * next;
    if (max < product) {
      max = product;
    }
  }
  return max;
}
