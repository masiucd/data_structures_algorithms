export function isPrime(n: number): boolean {
  if (n < 2) {
    return false;
  }
  for (let i = 2; i <= Math.sqrt(n); i++) {
    if (n % i === 0) return false;
  }
  return n > 1;
}

export function isPrimeFunctional(n: number): boolean {
  if (n < 2) {
    return false;
  }
  return !Array.from({length: n - 2}, (_, i) => i + 2).some((i) => n % i === 0);
}

export function isPrimeRecursive(n: number, i = 2): boolean {
  if (n < 2) {
    return false;
  }
  if (i > Math.sqrt(n)) {
    return true;
  }
  if (n % i === 0) {
    return false;
  }
  return isPrimeRecursive(n, i + 1);
}
