import {createCounter} from "../counter.ts";
import {assertEquals} from "~/dev_deps.ts";

Deno.test("createCounter", () => {
  const counter = createCounter(10);
  assertEquals(counter(), 10);
  assertEquals(counter(), 11);
  assertEquals(counter(), 12);
});
