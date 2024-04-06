import {expect, test} from "bun:test";
import {findErrorNums} from "./find-error-nums";

const testCases = [
  {input: [1, 2, 2, 4], expected: [2, 3]},
  {input: [1, 1], expected: [1, 2]},
  {input: [3, 2, 3, 4, 6, 5], expected: [3, 1]},
];

testCases.forEach(({input, expected}, idx) => {
  test(`findErrorNums - case ${idx + 1}`, () => {
    expect(findErrorNums(input)).toStrictEqual(expected);
  });
});
