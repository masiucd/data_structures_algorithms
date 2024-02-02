import {test, expect} from "bun:test";

import {groupFruits, groupFruitsV2, cartItems, groupToList} from "./cart";

test("should group fruits", () => {
  let input = cartItems;
  let output = {
    Mango: {id: 1, price: 20, qty: 1},
    Banana: {id: 2, price: 30, qty: 1},
    Apple: {id: 3, price: 40, qty: 4},
  };
  expect(groupFruits(input)).toEqual(output);
  expect(groupFruitsV2(input)).toEqual(output);
});

test("should group fruits to list", () => {
  let input = cartItems;
  let output = [
    {id: 1, name: "Mango", price: 20, qty: 1},
    {id: 2, name: "Banana", price: 30, qty: 1},
    {id: 3, name: "Apple", price: 40, qty: 4},
  ];
  expect(groupToList(input)).toEqual(output);
});
