#[allow(dead_code)]
pub fn bs_search(xs: &Vec<i32>, target: i32) -> (bool, isize) {
  let mut start = 0;
  let mut end = xs.len() - 1;
  while start < end {
    let middle = (start + end) / 2;
    if target == xs[middle] {
      return (true, middle as isize);
    }
    if target < xs[middle] {
      end = xs[middle] as usize;
    } else {
      start = xs[middle] as usize;
    }
  }
  (false, -1)
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_bs_search_not_found() {
    let xs = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    let res = bs_search(&xs, 22);
    assert_eq!(res, (false, -1));
  }

  #[test]
  fn test_bs_search_found() {
    let xs = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    let res = bs_search(&xs, 5);
    assert_eq!(res, (true, 4));
  }
}