export function isPalindrome(input: string) {
  let start = 0;
  let end = input.length - 1;

  while (start < end) {
    if (input[start] === input[end]) {
      start++;
      end--;
    } else {
      return false;
    }
  }
  return true;
}

export function isPalindromeRecursive(input: string) {
  if (input.length === 1) {
    return true;
  }

  if (input.length === 2) {
    return input[0] === input[1];
  }

  if (input[0] === input[input.length - 1]) {
    return isPalindromeRecursive(input.slice(1, input.length - 1));
  }

  return false;
}

let res = isPalindromeRecursive("abba");
console.log("res", res);
