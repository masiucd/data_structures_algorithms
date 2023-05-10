import {assertEquals, assert, assertFalse} from "~/dev_deps.ts";
import {List} from "../list.ts";

Deno.test("Append", () => {
  const list = new List<number>();
  list.append(1);
  list.append(2);
  list.append(3);
  assertEquals(list.toArray(), [1, 2, 3]);
});
