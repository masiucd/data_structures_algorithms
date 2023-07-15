pub mod stats {
  pub fn calculate_median(nums: &Vec<i32>) -> f64 {
    let mut nums = nums.clone();
    nums.sort();
    let mid = nums.len() / 2;
    if nums.len() % 2 == 0 {
      ((nums[mid] + nums[mid - 1]) as f64) / 2.0
    } else {
      nums[mid] as f64
    }
  }

  pub fn calculate_average(nums: &Vec<i32>) -> f64 {
    let mut sum = 0;
    for num in nums {
      sum += num;
    }
    (sum as f64) / (nums.len() as f64)
  }
  #[cfg(test)]
  mod tests {
    use super::*;

    #[test]
    fn test_calculate_median_it_works() {
      assert_eq!(calculate_median(&vec![1, 2, 3, 4, 5]), 3.0);
      assert_eq!(calculate_median(&vec![1, 2, 3, 4, 5, 6]), 3.5);
    }

    #[test]
    fn test_calculate_average_it_works() {
      assert_eq!(calculate_average(&vec![1, 2, 3, 4, 5]), 3.0);
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
