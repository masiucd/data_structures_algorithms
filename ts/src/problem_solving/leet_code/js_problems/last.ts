export function last<T>(list: T[]) {
  if (list.length === 0) return -1;
  return list[list.length - 1];
}
