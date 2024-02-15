export function arrayChange(inputArray: number[]): number {
  let count = 0;
  for (let i = 1; i < inputArray.length; i++) {
    if (inputArray[i] <= inputArray[i - 1]) {
      let difference = inputArray[i - 1] - inputArray[i];
      count += difference + 1;
      inputArray[i] = inputArray[i - 1] + 1;
    }
  }
  return count;
}
