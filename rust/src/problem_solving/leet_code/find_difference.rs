use std::collections::HashMap;

#[allow(dead_code)]
pub fn find_the_difference(s: String, t: String) -> char {
  let mut store = HashMap::new();
  for c in s.chars() {
    let count = store.entry(c).or_insert(0);
    *count += 1;
  }

  for c in t.chars() {
    if (store.contains_key(&c) && store.get(&c).unwrap().eq(&0)) || !store.contains_key(&c) {
      return c;
    } else {
      let count = store.entry(c).or_insert(0);
      *count -= 1;
    }
  }
  ' '
}

#[allow(dead_code)]
pub fn find_differences_two(s: String, t: String) -> char {
  let mut alpha: [i8; 26] = [0; 26];

  for char in s.chars() {
    let diff = (char as usize) - ('a' as usize);
    alpha[diff] += 1;
  }
  for char in t.chars() {
    let diff = (char as usize) - ('a' as usize);
    alpha[diff] -= 1;
  }

  let mut i = 0;
  while i < alpha.len() {
    if alpha[i] == -1 {
      break;
    }
    i += 1;
  }

  let char = (b'a' + (i as u8)) as char;
  char
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_find_the_difference() {
    let s = "abcd".to_string();
    let t = "abcde".to_string();
    let res = find_the_difference(s, t);
    assert_eq!(res, 'e');
  }

  #[test]
  fn test_find_differences_two() {
    let s = "abcd".to_string();
    let t = "abcde".to_string();
    let res = find_differences_two(s, t);
    assert_eq!(res, 'e');
  }
}