// All Longest Strings solution 1
export function allLongestStrings(inputArray: string[]): string[] {
  let currentMaxLength = inputArray[0].length;
  for (let i = 1; i < inputArray.length; i++) {
    if (currentMaxLength < inputArray[i].length) {
      currentMaxLength = inputArray[i].length;
    }
  }
  return inputArray.filter(({length}) => length === currentMaxLength);
}

// All Longest Strings solution 2
export function allLongestStringsTwo(inputArray: string[]): string[] {
  let currentMaxLength = 0;
  let result: string[] = [];
  for (let s of inputArray) {
    if (s.length > currentMaxLength) {
      currentMaxLength = s.length;
      result = [s];
    } else if (s.length === currentMaxLength) {
      result.push(s);
    }
  }
  return result;
}
