export function matrixElementsSum(matrix: number[][]): number {
  let sum = 0;
  for (let column = 0; column < matrix[0].length; column++) {
    for (let row = 0; row < matrix.length; row++) {
      let value = matrix[row][column];
      if (value === 0) {
        break;
      }
      sum += value;
    }
  }
  return sum;
}

export function matrixElementsSumTwo(matrix: number[][]): number {
  let bannedIndex: number[] = [];
  let sum = 0;
  for (let row = 0; row < matrix.length; row++) {
    let currentRow = matrix[row];
    for (let column = 0; column < currentRow.length; column++) {
      let value = currentRow[column];
      if (value === 0) {
        bannedIndex.push(column);
      } else if (!bannedIndex.includes(column)) {
        sum += value;
      }
    }
  }
  return sum;
}
