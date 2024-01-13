export function reverseList(numbers: number[]): number[] {
  let stack: number[] = [];
  let res: number[] = [];
  for (let n of numbers) {
    stack.push(n);
  }
  for (let [i] of numbers.entries()) {
    res[i] = stack.pop()!;
  }
  return res;
}

export function reverseListTwo(numbers: number[]): number[] {
  let stack: number[] = [];
  let res: number[] = [];
  for (let n of numbers) {
    stack.push(n);
  }
  while (stack.length > 0) {
    res.push(stack.pop()!);
  }
  return res;
}
