export function reverseString(input: string): string {
  if (input === "") return "";
  return reverseString(input.substring(1)) + input.at(0);
}

export function revStringTwo(input: string): string {
  if (input === "") return "";
  let head = input.at(0);
  let tail = input.substring(1);
  return revStringTwo(tail) + head;
}

export function reverseStringThree(input: string) {
  let reversed = "";
  for (let char of input) {
    reversed = char += reversed;
  }
  return reversed;
}

export function reverseStringFour(input: string) {
  let reversed = "";
  for (let i = input.length - 1; i >= 0; i--) {
    reversed += input.at(i);
  }
  return reversed;
}

export function reverseStringFive(input: string) {
  return input.split("").reverse().join("");
}
