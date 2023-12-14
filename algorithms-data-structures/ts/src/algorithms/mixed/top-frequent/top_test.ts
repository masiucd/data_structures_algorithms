import {test, expect} from "bun:test";
import {topKFrequent} from "./top";

test("topKFrequent", () => {
  let input = [1, 1, 1, 2, 2, 3];
  let output = [1, 2];
  expect(topKFrequent(input, 2)).toEqual(output);
});

test("topKFrequent", () => {
  let input = [1];
  let output = [1];
  expect(topKFrequent(input, 1)).toEqual(output);
});
