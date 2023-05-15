pub mod prime {
  #[allow(dead_code)]
  pub fn is_prime(n: i32) -> bool {
    if n <= 1 {
      return false;
    }
    if n <= 3 {
      return true;
    }
    if n % 2 == 0 || n % 3 == 0 {
      return false;
    }
    let mut i = 5;
    while i * i <= n {
      if n % i == 0 || n % (i + 2) == 0 {
        return false;
      }
      i += 6;
    }
    true
  }

  pub fn prime_nums_from_range(start: i32, end: i32) -> Vec<i32> {
    let mut primes = Vec::new();
    for i in start..end {
      if is_prime(i) {
        primes.push(i);
      }
    }
    primes
  }

  #[cfg(test)]
  mod tests {
    use super::*;

    #[test]
    fn test_is_prime_it_works() {
      assert_eq!(is_prime(2), true);
      assert_eq!(is_prime(3), true);
      assert_eq!(is_prime(4), false);
      assert_eq!(is_prime(5), true);
      assert_eq!(is_prime(6), false);
      assert_eq!(is_prime(7), true);
    }

    #[test]
    fn test_prime_nums_from_range() {
      assert_eq!(prime_nums_from_range(1, 10), vec![2, 3, 5, 7]);
    }
  }
}