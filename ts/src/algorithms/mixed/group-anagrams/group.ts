export function groupAnagrams(strs: string[]): string[][] {
  let res: string[][] = [];
  let soretedStrs = Array.from(strs.map((x) => [...x].sort().join("")));

  let store: Record<string, string[]> = {};
  for (let [i, word] of soretedStrs.entries()) {
    if (!store[word]) {
      store[word] = [strs[i]];
    } else {
      store[word].push(strs[i]);
    }
  }
  for (let [_, v] of Object.entries(store)) {
    res.push(v);
  }
  return res;
}
