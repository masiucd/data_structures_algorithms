#[allow(dead_code)]
pub fn solution(nums: Vec<i32>) -> i32 {
  let mut smallest = nums[0];
  for num in nums {
    if num < smallest {
      smallest = num;
    }
  }
  smallest
}

#[cfg(test)]
mod tests {
  use super::*;
  #[test]
  fn test_solution_it_works() {
    assert_eq!(solution(vec![1]), 1);
    assert_eq!(solution(vec![1, 2, 3]), 1);
    assert_eq!(solution(vec![1, 2, -3]), -3);
    assert_eq!(solution(vec![1, -2, 3]), -2);
  }
}