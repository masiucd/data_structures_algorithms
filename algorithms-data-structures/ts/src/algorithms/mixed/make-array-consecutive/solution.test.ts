import {test, expect} from "bun:test";
import {makeArrayConsecutive} from "./solution";

test("makeArrayConsecutive", () => {
  expect(makeArrayConsecutive([6, 2, 3, 8])).toEqual(3);
  expect(makeArrayConsecutive([0, 3])).toEqual(2);
  expect(makeArrayConsecutive([5, 4, 6])).toEqual(0);
  expect(makeArrayConsecutive([6, 3])).toEqual(2);
  expect(makeArrayConsecutive([1])).toEqual(0);
});
