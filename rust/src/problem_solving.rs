pub mod strings {
  #[allow(dead_code)]
  pub fn pig_latin(input: &str) -> String {
    let vowels = ['a', 'e', 'i', 'o', 'u'];
    let mut res = input.to_string();
    let first = input.chars().next().unwrap();
    if vowels.contains(&first) {
      // add to the end hay to the end of the word
      res.push_str("-hay");
    } else if !vowels.contains(&first) {
      // move the first consonant to the end of the word and add ay
      res.remove(0);
      res.push_str(&format!("-{}ay", first));
    }
    res
  }

  #[cfg(test)]
  mod tests {
    use super::*;
    #[test]
    fn test_pig_latin() {
      let res = pig_latin("apple");
      assert_eq!(res, "apple-hay");
    }
    #[test]
    fn test_pig_latin_2() {
      let res = pig_latin("first");
      assert_eq!(res, "irst-fay");
    }

    #[test]
    fn test_pig_latin_3() {
      let res = pig_latin("Здравствуйте");
      assert_eq!(res, "дравствуйте-Зay");
    }
  }
}
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

  #[allow(dead_code)]
  pub fn smallest_int(nums: Vec<i32>) -> i32 {
    let mut smallest = nums[0];
    for num in nums {
      if num < smallest {
        smallest = num;
      }
    }
    smallest
  }

  #[allow(dead_code)]
  pub fn positive_sum(slice: &[i32]) -> i32 {
    slice
      .iter()
      .filter(|x| x.is_positive())
      .sum()
  }

  #[allow(dead_code)]
  fn is_anagram_one(s1: &str, s2: &str) -> bool {
    let mut chars = s1.chars().into_iter().collect::<Vec<char>>();
    let mut chars2 = s2.chars().into_iter().collect::<Vec<char>>();
    chars.sort();
    chars2.sort();
    chars == chars2
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

    #[test]
    fn test_is_anagram() {
      let res = is_anagram_one("anagram", "nagaram");
      assert!(res);

      let res = !is_anagram_one("anagram", "nagaras");
      assert!(res);
    }

    #[test]
    fn test_smallest_int() {
      assert_eq!(smallest_int(vec![1]), 1);
      assert_eq!(smallest_int(vec![1, 2, 3]), 1);
      assert_eq!(smallest_int(vec![1, 2, -3]), -3);
      assert_eq!(smallest_int(vec![1, -2, 3]), -2);
    }

    #[test]
    fn test_positive_sum() {
      assert_eq!(positive_sum(&vec![1, 2, 3, 4]), 10);
      assert_eq!(positive_sum(&vec![1, 2, 3, 4, -10]), 10);
      assert_eq!(positive_sum(&vec![1, 2, -3, 4, -10]), 7);
    }
  }
}
