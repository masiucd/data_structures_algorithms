pub struct Solution;

impl Solution {
  #[allow(dead_code)]
  pub fn length_of_last_word(s: String) -> i32 {
    s.trim().split(' ').last().unwrap().len() as i32
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_length_of_last_word() {
    let s = String::from("Hello World");
    let res = Solution::length_of_last_word(s);
    assert_eq!(res, 5);
  }
  #[test]
  fn test_length_of_last_word_2() {
    let s = String::from("Hello World    ");
    let res = Solution::length_of_last_word(s);
    assert_eq!(res, 5);
  }

  #[test]
  fn test_length_of_last_word_3() {
    let s = String::from("");
    let res = Solution::length_of_last_word(s);
    assert_eq!(res, 0);
  }
}