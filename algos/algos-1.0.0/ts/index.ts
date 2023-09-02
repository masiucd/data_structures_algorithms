import {expect, test} from "bun:test";

function myPow(n: number, exponent: number) {
  let res = 1;
  let i = 1;
  while (i <= exponent) {
    res *= n;
    i++;
  }

  return res;
}

let res = myPow(2, 3);
console.log("res", res);

test("myPow", () => {
  expect(myPow(2, 3)).toBe(8);
});
