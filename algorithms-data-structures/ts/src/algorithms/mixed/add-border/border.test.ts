import {test, expect} from "bun:test";
import {addBorder} from "./border";

test("addBorder", () => {
  expect(addBorder(["abc", "def"])).toEqual([
    "*****",
    "*abc*",
    "*def*",
    "*****",
  ]);
  expect(addBorder(["a"])).toEqual(["***", "*a*", "***"]);
});
