function topKFrequent(nums: number[], k: number): number[] {
  let store = new Map<number, number>();
  for (let n of nums) {
    if (store.has(n)) {
      let amount = store.get(n);
      if (amount) {
        store.set(n, amount + 1);
      }
    } else {
      store.set(n, 1);
    }
  }

  let buckets: number[][] = new Array(nums.length + 1).map(() => []);
  for (let [k, v] of store) {
    if (!buckets[v]) buckets[v] = [];
    buckets[v].push(k);
  }

  let result: number[] = [];
  for (let i = buckets.length - 1; i >= 0; i--) {
    let bucket = buckets[i];

    if (bucket && k > 0) {
      result.push(...bucket);
      k -= bucket.length;
    }
  }
  return result;
}

console.log(topKFrequent([1, 1, 1, 2, 2, 3], 2));
