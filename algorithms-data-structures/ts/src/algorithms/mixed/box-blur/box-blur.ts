export function boxBlur(image: number[][]): number[][] {
  let res = [];
  for (let row = 0; row < image.length - 2; row++) {
    let xs = [];
    for (let col = 0; col < image[0].length - 2; col++) {
      console.log(image[row][col]);
      let top = [image[row][col], image[row][col + 1], image[row][col + 2]];
      let middle = [
        image[row + 1][col],
        image[row + 1][col + 1],
        image[row + 1][col + 2],
      ];
      let bottom = [
        image[row + 2][col],
        image[row + 2][col + 1],
        image[row + 2][col + 2],
      ];
      xs[col] = calcAverage(top, middle, bottom);
    }
    res.push(xs);
  }
  return res;
}

function calcAverage(
  top: number[],
  middle: number[],
  bottom: number[]
): number {
  let sum = top.reduce((a, b) => a + b, 0);
  sum += middle.reduce((a, b) => a + b, 0);
  sum += bottom.reduce((a, b) => a + b, 0);
  return Math.floor(sum / 9);
}
