import {expect, test} from "bun:test";
import {BinarySearchTree} from "./tree";

function createTree() {
  const tree = new BinarySearchTree();
  tree.insert(5);
  tree.insert(3);
  tree.insert(7);
  tree.insert(2);
  tree.insert(4);
  tree.insert(6);
  tree.insert(8);
  return tree;
}

test("BinarySearchTree", () => {
  const tree = createTree();
  expect(tree.root?.value).toBe(5);
  expect(tree.root?.left?.value).toBe(3);
  expect(tree.root?.right?.value).toBe(7);
  expect(tree.root?.left?.left?.value).toBe(2);
  expect(tree.root?.left?.right?.value).toBe(4);
  expect(tree.root?.right?.left?.value).toBe(6);
  expect(tree.root?.right?.right?.value).toBe(8);
});

test("BinarySearchTree#bfs", () => {
  const tree = createTree();
  expect(tree.bfs()).toEqual([5, 3, 7, 2, 4, 6, 8]);
});

test("BinarySearchTree#min", () => {
  const tree = createTree();
  expect(tree.min()).toBe(2);
});

test("BinarySearchTree#max", () => {
  const tree = createTree();
  expect(tree.max()).toBe(8);
});

test("BinarySearchTree#dfs", () => {
  const tree = createTree();
  expect(tree.dfs("PRE")).toEqual([5, 3, 2, 4, 7, 6, 8]);
  expect(tree.dfs("IN")).toEqual([2, 3, 4, 5, 6, 7, 8]);
  expect(tree.dfs("POST")).toEqual([2, 4, 3, 6, 8, 7, 5]);
});

test("BinarySearchTree#find", () => {
  const tree = createTree();
  expect(tree.find(5)?.value).toBe(5);
  expect(tree.find(3)?.value).toBe(3);
  expect(tree.find(7)?.value).toBe(7);
  expect(tree.find(2)?.value).toBe(2);
  expect(tree.find(4)?.value).toBe(4);
  expect(tree.find(6)?.value).toBe(6);
  expect(tree.find(8)?.value).toBe(8);
  expect(tree.find(9)).toBe(null);
});

test("BinarySearchTree#size", () => {
  const tree = createTree();
  expect(tree.size()).toBe(7);
});
