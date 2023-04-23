use std::collections::{ HashMap, HashSet };

#[allow(dead_code)]
pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
  let mut store: HashMap<i32, usize> = HashMap::new();
  for (i, num) in nums.iter().enumerate() {
    if store.contains_key(num) && i - store.get(num).unwrap() <= (k as usize) {
      return true;
    }
    store.insert(*num, i);
  }
  false
}

#[allow(dead_code)]
pub fn contains_nearby_duplicate_two(nums: Vec<i32>, k: i32) -> bool {
  let mut lookup: HashSet<i32> = HashSet::new();
  let k = k as usize;
  for i in 0..nums.len() {
    if i > k {
      lookup.remove(&nums[i - k - 1]);
    }
    if !lookup.insert(nums[i]) {
      return true;
    }
  }

  false
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_contains_nearby_duplicate() {
    let nums = vec![1, 2, 3, 1];
    let k = 3;
    let result = contains_nearby_duplicate(nums, k);
    assert_eq!(result, true);

    let nums = vec![1, 0, 1, 1];
    let k = 1;
    let result = contains_nearby_duplicate(nums, k);
    assert_eq!(result, true);

    let nums = vec![1, 2, 3, 1, 2, 3];
    let k = 2;
    let result = contains_nearby_duplicate(nums, k);
    assert_eq!(result, false);
  }

  #[test]
  fn test_contains_nearby_duplicate2() {
    // let nums = vec![1, 2, 3, 1];
    // let k = 3;
    // let result = contains_nearby_duplicate2(nums, k);
    // assert_eq!(result, true);

    // let nums = vec![1, 0, 1, 1];
    // let k = 1;
    // let result = contains_nearby_duplicate2(nums, k);
    // assert_eq!(result, true);

    // let nums = vec![1, 2, 3, 1, 2, 3];
    // let k = 2;
    // let result = contains_nearby_duplicate2(nums, k);
    // assert_eq!(result, false);
  }
}