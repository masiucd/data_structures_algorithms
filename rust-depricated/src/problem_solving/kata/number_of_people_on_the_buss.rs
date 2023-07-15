#[allow(dead_code)]
pub fn number_of_people_on_the_buss(buss_stops: &[(i32, i32)]) -> i32 {
  let mut res = 0;
  for (a, b) in buss_stops.iter() {
    res += a - b;
  }
  res
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_buss_stops() {
    let xs = &[
      (10, 0),
      (3, 5),
      (5, 8),
    ];
    assert_eq!(number_of_people_on_the_buss(xs), 5)
  }
}
