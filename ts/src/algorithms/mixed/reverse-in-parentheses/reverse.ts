export function reverseInParentheses(inputString: string): string {
  let stack: string[] = [];
  for (let ch of inputString) {
    if (ch === ")") {
      let temp: string[] = [];
      while (stack[stack.length - 1] !== "(") {
        let popped = stack.pop()!;
        temp.push(popped);
      }
      stack.pop();
      stack.push(...temp);
    } else {
      stack.push(ch);
    }
  }
  return stack.join("");
}
