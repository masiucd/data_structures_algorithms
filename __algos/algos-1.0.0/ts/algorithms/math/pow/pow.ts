export function myPow(base: number, exponent: number) {
  let res = 1;
  let i = 1;
  while (i <= exponent) {
    res *= base;
    i++;
  }

  return res;
}

export function myPow2(base: number, exponent: number): number {
  if (exponent === 0) {
    return 1;
  }

  if (exponent === 1) {
    return base;
  }

  return base * myPow2(base, exponent - 1);
}
