#[allow(dead_code)]
pub fn factors(n: i32) -> Vec<i32> {
  let mut res = vec![];
  for i in 1..n + 1 {
    if n % i == 0 {
      res.push(i);
    }
  }
  res
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_factors() {
    assert_eq!(factors(1), vec![1]);
    assert_eq!(factors(20), vec![1, 2, 4, 5, 10, 20]);
    assert_eq!(factors(10), vec![1, 2, 5, 10]);
    assert_eq!(factors(9), vec![1, 3, 9]);
    assert_eq!(factors(7), vec![1, 7]);
  }
}