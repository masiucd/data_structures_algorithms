#[allow(dead_code)]
pub fn single_number(nums: Vec<i32>) -> i32 {
  let mut store = std::collections::HashMap::new();
  for num in nums {
    let count = store.entry(num).or_insert(0);
    *count += 1;
  }
  for (key, value) in store {
    if value == 1 {
      return key;
    }
  }
  -1
}

#[allow(dead_code)]
pub fn single_number2(nums: Vec<i32>) -> i32 {
  let mut ret = 0;
  for num in nums {
    ret = ret ^ num;
  }
  return ret;
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_single_number() {
    assert_eq!(single_number(vec![2, 2, 1]), 1);
    assert_eq!(single_number(vec![4, 1, 2, 1, 2]), 4);
  }

  #[test]
  fn test_single_number2() {
    assert_eq!(single_number2(vec![2, 2, 1]), 1);
    assert_eq!(single_number2(vec![4, 1, 2, 1, 2]), 4);
  }
}