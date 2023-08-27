import {expect, test} from "bun:test";
import {myPow, myPow2} from "./pow";

test(" base= 2 exponent= 4 should be 8  ", () => {
  expect(myPow(2, 3)).toBe(8);
});

test(" base= 2 exponent= 4 should be 16 ", () => {
  expect(myPow(2, 4)).toBe(16);
});

test("base = 2 exponent= 5 should be 32", () => {
  expect(myPow(2, 5)).toBe(32);
});

test(" base= 2 exponent= 4 should be 8  ", () => {
  expect(myPow2(2, 3)).toBe(8);
});

test(" base= 2 exponent= 4 should be 16 ", () => {
  expect(myPow2(2, 4)).toBe(16);
});

test("base = 2 exponent= 5 should be 32", () => {
  expect(myPow2(2, 5)).toBe(32);
});
