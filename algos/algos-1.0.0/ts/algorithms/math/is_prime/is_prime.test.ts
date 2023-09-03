import {isPrime, isPrimeFunctional, isPrimeRecursive} from "./is_prime";
import {expect, test} from "bun:test";

test("isPrime should return false for 0", () => {
  expect(isPrime(0)).toBe(false);

  expect(isPrimeFunctional(0)).toBe(false);

  expect(isPrimeRecursive(0)).toBe(false);
});

test("isPrime should return false for 1", () => {
  expect(isPrime(1)).toBe(false);

  expect(isPrimeFunctional(1)).toBe(false);

  expect(isPrimeRecursive(1)).toBe(false);
});

test("isPrime should return true for 2", () => {
  expect(isPrime(2)).toBe(true);

  expect(isPrimeFunctional(2)).toBe(true);

  expect(isPrimeRecursive(2)).toBe(true);
});

test("isPrime should return true for 3", () => {
  expect(isPrime(3)).toBe(true);

  expect(isPrimeFunctional(3)).toBe(true);

  expect(isPrimeRecursive(3)).toBe(true);
});

test("isPrime should return false for 4", () => {
  expect(isPrime(4)).toBe(false);

  expect(isPrimeFunctional(4)).toBe(false);

  expect(isPrimeRecursive(4)).toBe(false);
});

test("isPrime should return true for 5", () => {
  expect(isPrime(5)).toBe(true);

  expect(isPrimeFunctional(5)).toBe(true);

  expect(isPrimeRecursive(5)).toBe(true);
});
