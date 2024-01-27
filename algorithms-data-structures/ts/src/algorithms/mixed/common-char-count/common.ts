export function commonCharacterCount(s1: string, s2: string): number {
  let charMap = new Map<string, number>();
  for (let c of s1) {
    if (!charMap.has(c)) {
      charMap.set(c, 1);
    } else {
      let n = charMap.get(c)!;
      charMap.set(c, (n += 1));
    }
  }
  let count = 0;
  for (let c of s2) {
    if (charMap.has(c)) {
      let n = charMap.get(c) as number;
      if (n > 0) {
        count++;
        charMap.set(c, (n -= 1));
      }
    }
  }
  return count;
}
