import {assertEquals, assert, assertFalse} from "~/dev_deps.ts";
import {List} from "../list.ts";

Deno.test("Append", () => {
  const list = new List<number>();
  list.append(1);
  list.append(2);
  list.append(3);
  assertEquals(list.toArray(), [1, 2, 3]);
});

Deno.test("Prepend", () => {
  const list = new List<number>();
  list.prepend(1);
  list.prepend(2);
  list.prepend(3);
  assertEquals(list.toArray(), [3, 2, 1]);
});

Deno.test("Get", () => {
  const list = new List<number>();
  list.append(1);
  list.append(2);
  list.append(3);
  assertEquals(list.get(0), 1);
  assertEquals(list.get(1), 2);
  assertEquals(list.get(2), 3);
});
