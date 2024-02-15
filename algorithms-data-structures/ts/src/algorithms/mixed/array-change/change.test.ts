import {test, describe, expect} from "bun:test";
import {arrayChange} from "./change";

describe("arrayChange", () => {
  test("[1,1,1]", () => {
    let res = arrayChange([1, 1, 1]);
    expect(res).toBe(3);
  });

  test("[-1000,0,-2,0]", () => {
    let res = arrayChange([-1000, 0, -2, 0]);
    expect(res).toBe(5);
  });
});
