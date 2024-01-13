function adjacentElementsProduct(inputArray: number[]): number {
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

console.log(adjacentElementsProduct([3, 6, -2, -5, 7, 3]));
console.log(adjacentElementsProduct([9, 5, 10, 2, 24, -1, -48]));
