pub mod search {
  pub fn linnear_search(xs: Vec<i32>, x: i32) -> Option<usize> {
    for (i, &v) in xs.iter().enumerate() {
      if v == x {
        return Some(i);
      }
    }
    return None;
  }

  pub fn binary_search(xs: Vec<i32>, x: i32) -> Option<usize> {
    let mut low = 0;
    let mut high = xs.len() - 1;
    while low <= high {
      let mid = (low + high) / 2;
      let guess = xs[mid];
      if guess == x {
        return Some(mid);
      }
      if guess > x {
        high = mid - 1;
      } else {
        low = mid + 1;
      }
    }
    return None;
  }
  #[cfg(test)]
  mod tests {
    use super::*;

    #[test]
    fn test_binary_search_it_works() {
      assert_eq!(binary_search(vec![1, 2, 3, 4, 5], 3), Some(2));
    }

    #[test]
    fn test_binary_search_it_fails() {
      assert_eq!(binary_search(vec![1, 2, 3, 4, 5], 6), None);
    }

    #[test]
    fn test_linnear_search_it_works() {
      assert_eq!(linnear_search(vec![1, 2, 3, 4, 5], 3), Some(2));
    }
  }
}
pub mod recursion {
  pub fn sum(xs: Vec<i32>) -> i32 {
    if xs.len() == 0 {
      return 0;
    }
    let head = xs.first().unwrap();
    let tail = &xs[1..];
    return head + sum(tail.to_vec());
  }

  pub fn is_palindrome(s: String) -> bool {
    if s.len() == 0 || s.len() == 1 {
      return true;
    }
    let start = s.chars().nth(0).unwrap();
    let end = s
      .chars()
      .nth(s.len() - 1)
      .unwrap();
    let middle: String = s
      .chars()
      .skip(1)
      .take(s.len() - 2)
      .collect();
    return start == end && is_palindrome(middle.to_string());
  }

  #[cfg(test)]
  mod tests {
    use super::*;

    #[test]
    fn test_is_palindrome_it_works() {
      assert_eq!(is_palindrome("".to_string()), true);
      assert_eq!(is_palindrome("racecar".to_string()), true);
      assert_eq!(is_palindrome("otto".to_string()), true);
    }

    #[test]
    fn test_is_palindrome_it_fails() {
      assert_eq!(is_palindrome("cwks".to_string()), false);
      assert_eq!(is_palindrome("1916".to_string()), false);
    }

    #[test]
    fn test_sum_it_works() {
      assert_eq!(sum(vec![1, 2, 3, 4, 5]), 15);
    }
  }

  pub fn fibonacci_dp(n: i32) -> i32 {
    let mut a = 1;
    let mut b = 1;
    for _i in 0..n - 1 {
      let temp = a;
      a = a + b;
      b = temp;
    }
    a
  }
}
pub mod bits {
  pub fn count(n: u32) -> u32 {
    let mut count: u32 = 0;
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
    fn test_count() {
      assert_eq!(count(0), 0);
      assert_eq!(count(4), 1);
      assert_eq!(count(7), 3);
      assert_eq!(count(9), 2);
      assert_eq!(count(10), 2);
    }
  }
  pub mod search {}
}