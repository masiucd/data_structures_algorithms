import {test, expect} from "bun:test";
import {dnaToRna} from "./dna-to-rna";

test("dnaToRna", () => {
  expect(dnaToRna("TTTT")).toEqual("UUUU");
  expect(dnaToRna("GCAT")).toEqual("GCAU");
  expect(dnaToRna("GACCGCCGCC")).toEqual("GACCGCCGCC");
});
