use std::collections::HashMap;

#[allow(dead_code)]
pub fn solution(nums: Vec<i32>) -> i32 {
  let mut nums_map: HashMap<i32, i32> = HashMap::new();
  for n in nums {
    *nums_map.entry(n).or_insert(0) += 1;
  }
  for (key, value) in nums_map.iter() {
    if *value == 1 {
      return *key;
    }
  }
  -1
}

#[cfg(test)]
mod tests {
  use super::*;
  #[test]
  fn test_it_works() {
    let xs = vec![6, 1, 3, 3, 3, 6, 6];
    assert_eq!(solution(xs), 1);
    let xs = vec![13, 19, 13, 13];
    assert_eq!(solution(xs), 19);
  }
}