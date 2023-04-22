use std::collections::HashSet;

#[allow(dead_code)]
pub fn happy_number(n: i32) -> bool {
  let mut seen = HashSet::new();
  let mut num = n;
  while num != 1 && !seen.contains(&num) {
    seen.insert(num);
    num = calculate_square_sum(num);
  }
  num == 1
}

#[allow(dead_code)]
fn calculate_square_sum(mut n: i32) -> i32 {
  let mut res = 0;
  while n > 0 {
    let num = n % 10;
    res += i32::pow(num, 2);
    n = n / 10;
  }
  res
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_happy_number() {
    assert!(happy_number(19));
  }

  #[test]
  fn test_happy_number_2() {
    assert!(!happy_number(2));
  }

  #[test]
  fn test_happy_number_3() {
    assert!(!happy_number(3));
  }
}