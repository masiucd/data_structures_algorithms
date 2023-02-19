pub fn get_concatenation(nums: Vec<i32>) -> Vec<i32> {
  let mut result = nums.clone();
  result.extend(nums);
  result
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_get_concatenation() {
    assert_eq!(get_concatenation(vec![1, 2, 3]), vec![1, 2, 3, 1, 2, 3]);
    assert_eq!(get_concatenation(vec![1, 3, 2, 1]), vec![1, 3, 2, 1, 1, 3, 2, 1]);
  }
}