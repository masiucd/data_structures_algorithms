pub fn solution(slice: &[i32]) -> i32 {
  slice
    .iter()
    .filter(|x| x.is_positive())
    .sum()
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_solution() {
    assert_eq!(solution(&vec![1, 2, 3, 4]), 10);
    assert_eq!(solution(&vec![1, 2, 3, 4, -10]), 10);
    assert_eq!(solution(&vec![1, 2, -3, 4, -10]), 7);
  }
}