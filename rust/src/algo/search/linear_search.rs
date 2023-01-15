#[allow(dead_code)]
pub fn linear_search<T: PartialEq>(xs: &Vec<T>, needle: T) -> i32 {
  for i in 0..xs.len() {
    if xs[i] == needle {
      return i as i32;
    }
  }
  -1
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_linear_search() {
    let xs = vec![1, 2, 3, 4, 5];
    let needle = 3;
    let res = linear_search(&xs, needle);
    assert_eq!(res, 2);
  }

  #[test]
  fn test_linear_search_not_found() {
    let xs = vec![1, 2, 3, 4, 5];
    let needle = 6;
    let res = linear_search(&xs, needle);
    assert_eq!(res, -1);
  }
}