import {assertEquals} from "~/dev_deps.ts";
import {last} from "../last.ts";

Array.prototype.last = function () {
  if (this.length === 0) return -1;
  return this[this.length - 1];
};

Deno.test("Last prototype", () => {
  let list = [1, 2, 3];
  const last = list.last();
  assertEquals(last, 3);

  list = [];
  const lastEmpty = list.last();
  assertEquals(lastEmpty, -1);
});

Deno.test("Last function", () => {
  let list = [1, 2, 3];
  const lastItem = last(list);
  assertEquals(lastItem, 3);

  list = [];
  const lastEmpty = last(list);
  assertEquals(lastEmpty, -1);
});
