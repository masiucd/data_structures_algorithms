export function addBorder(picture: string[]): string[] {
  let xs = picture.map((x) => "*" + x + "*");
  return ["*".repeat(xs[0].length), ...xs, "*".repeat(xs[0].length)];
}
