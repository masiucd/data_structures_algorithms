// brute force! - O(n^2)
export function maxArea(heights: number[]): number {
  let res = 0;
  for (let i = 0; i < heights.length; i++) {
    for (let j = i + 1; j < heights.length; j++) {
      let width = j - i;
      let height = Math.min(heights[i], heights[j]);
      let area = width * height;
      res = Math.max(res, area);
    }
  }

  return res;
}

// two pointers approach (linear) - O(n)
export function maxAreaLinnear(heights: number[]): number {
  let res = 0;
  let left = 0;
  let right = heights.length - 1;
  while (left < right) {
    let width = right - left;
    let height = Math.min(heights[left], heights[right]);
    let area = width * height;
    res = Math.max(res, area);
    if (heights[left] < heights[right]) {
      left++;
    } else {
      right--;
    }
  }

  return res;
}
