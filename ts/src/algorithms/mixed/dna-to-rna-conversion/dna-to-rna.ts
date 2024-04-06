export function dnaToRna(dna: string) {
  let res = "";
  for (let char of dna) {
    switch (char) {
      case "T":
        res += "U";
        break;
      case "U":
        res += "T";
        break;

      default:
        res += char;
    }
  }
  return res;
}
