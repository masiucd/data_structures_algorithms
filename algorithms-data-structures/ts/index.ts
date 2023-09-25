import {isValid} from "./algorithms/stack/is-valid/is-valid";

console.log(
  isValid("()"), // true
  isValid("()[]{}"), // true
  isValid("(]"), // false
  isValid("([)]"), // false
  isValid("{[]}"), // true
  isValid("]"), // false
  isValid("(("), // false
  isValid("(("), // false
  isValid("){") // false
);
