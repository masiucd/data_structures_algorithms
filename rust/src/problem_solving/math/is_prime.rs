pub fn is_prime(n: i32) -> bool {
  for i in 2..(n as f64).sqrt() as i32 {
    if n % i == 0 {
      return false;
    }
  }
  true
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_is_prime() {
    assert!(is_prime(2));
    assert!(is_prime(3));
    assert!(is_prime(7));
    assert!(is_prime(11));
    assert!(is_prime(17));
  }
}