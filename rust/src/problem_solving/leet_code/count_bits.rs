#[allow(dead_code)]
pub fn count_bits(n: i32) -> i32 {
  format!("{:b}", n)
    .chars()
    .map(|c| c.to_string())
    .map(|s| s.parse::<i32>().unwrap())
    .filter(|x| x == &1)
    .collect::<Vec<_>>()
    .len() as i32
}

#[allow(dead_code)]
pub fn count_bits_two(n: i32) -> i32 {
  let mut count = 0;
  let mut n = n;
  while n > 0 {
    count += n & 1;
    n >>= 1;
  }
  count
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_count_bits() {
    assert_eq!(count_bits(5), 2);
    assert_eq!(count_bits(10), 2);
    assert_eq!(count_bits(15), 4);
  }

  #[test]
  fn test_count_bits_two() {
    assert_eq!(count_bits_two(5), 2);
    assert_eq!(count_bits_two(10), 2);
    assert_eq!(count_bits_two(15), 4);
  }
}