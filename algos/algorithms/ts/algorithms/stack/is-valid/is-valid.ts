export function isValid(s: string): boolean {
  let stack: string[] = [];
  let closingBrackets = new Map<string, string>([
    [")", "("],
    ["}", "{"],
    ["]", "["],
  ]);
  for (const char of s) {
    // check if char is closing
    if (closingBrackets.has(char)) {
      // check that stack is not empty and that the last element is the opening bracket
      if (stack.length > 0 && closingBrackets.get(char) === stack.at(-1)) {
        // pop the last element
        stack.pop();
      } else {
        // return false if the stack is empty or the last element is not the opening bracket
        return false;
      }
    } else {
      // push to the stack
      stack.push(char);
    }
  }
  return stack.length === 0;
}
