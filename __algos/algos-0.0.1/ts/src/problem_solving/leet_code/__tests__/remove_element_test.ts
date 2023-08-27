import {assertEquals} from "~/dev_deps.ts";
import {removeElement} from "../remove_element/solution.ts";

Deno.test("removeElement works", () => {
  const nums = [3, 2, 2, 3];
  const val = 3;
  const result = removeElement(nums, val);
  assertEquals(result, 2);
  assertEquals(nums, [2, 2]);
});

Deno.test("removeElement works longer list", () => {
  const nums = [0, 1, 2, 2, 3, 0, 4, 2];
  const val = 2;
  const result = removeElement(nums, val);
  assertEquals(result, 5);
  assertEquals(nums, [0, 1, 3, 0, 4]);
});
