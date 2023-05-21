pub mod basics {
  pub fn average(xs: &Vec<i32>) -> i32 {
    if xs.is_empty() {
      return 0;
    }
    let size = xs.len() as i32;
    let sum: i32 = xs.iter().sum();
    sum / size
  }

  pub fn median(xs: &Vec<i32>) -> f64 {
    if xs.len() == 0 {
      return 0.0;
    }
    let xs = xs
      .iter()
      .map(|x| *x as f64)
      .collect::<Vec<f64>>();

    if xs.len() == 2 {
      return (xs[0] + xs[1]) / 2.0;
    }
    if xs.len() == 3 {
      return xs[xs.len() / 2];
    }
    let size = xs.len();
    let middle = (size / 2) as usize;
    if !is_even(size as i32) {
      return xs[middle] as f64;
    }
    let a = xs[middle - 1];
    let b = xs[middle];
    return (a + b) / 2.0;
  }

  pub fn is_even(n: i32) -> bool {
    n % 2 == 0
  }

  #[cfg(test)]
  mod tests {
    use super::*;
    #[test]
    fn test_average() {
      assert_eq!(average(&vec![1, 2]), 1);
      assert_eq!(average(&vec![1, 2, 3]), 2);
      assert_eq!(average(&vec![1, 2, 4]), 2);
      assert_eq!(average(&vec![]), 0);
      assert_eq!(average(&vec![1]), 1);
      assert_eq!(average(&vec![2]), 2);
    }
    #[test]
    fn test_median() {
      assert_eq!(median(&vec![]), 0.0);
      assert_eq!(median(&vec![1]), 1.0);
      assert_eq!(median(&vec![1, 2]), 1.5);
      assert_eq!(median(&vec![1, 2, 3]), 2.0);
      assert_eq!(median(&vec![1, 2, 3, 4, 5]), 3.0);
      assert_eq!(median(&vec![1, 2, 3, 4]), 2.5);
      assert_eq!(median(&vec![1, 2, 3, 4, 5, 6]), 3.5);
    }
  }
}
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