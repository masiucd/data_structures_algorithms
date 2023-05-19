pub mod kata {
  use std::collections::{ HashMap, HashSet };

  #[allow(dead_code)]
  pub fn int_to_roman(num: i32) -> String {
    let values = vec![1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
    let symbols = vec!["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
    let mut res = String::new();
    let mut num = num;
    for (i, n) in values.iter().enumerate() {
      while num >= *n {
        res.push_str(symbols[i]);
        num -= n;
      }
    }
    res
  }

  #[allow(dead_code)]
  pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut map: HashMap<String, Vec<String>> = HashMap::new();
    for s in strs {
      let mut chars: Vec<char> = s.chars().collect();
      chars.sort();
      let key = chars.into_iter().collect();
      map
        .entry(key)
        .or_insert(vec![])
        .push(s);
    }
    map.values().cloned().collect()
  }

  pub fn clean_string(input: &str) -> String {
    let mut result = String::new();
    for c in input.chars() {
      if c == '#' {
        if result.len() > 0 {
          result.pop();
        }
      } else {
        result.push(c);
      }
    }
    return result;
  }

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
    fn test_int_to_roman() {
      let res = int_to_roman(1000);
      assert_eq!(res, "M");
    }

    #[test]
    fn test_int_to_roman_2() {
      let res = int_to_roman(1990);
      assert_eq!(res, "MCMXC");
    }
  }
}