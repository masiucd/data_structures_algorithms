#[allow(dead_code)]
pub fn solution(input: &Vec<&str>, k: usize) -> String {
  let mut res = String::from("");
  if k > 0 && input.len() >= k {
    for i in 0..input.len() - k + 1 {
      let xs = &input[i..i + k].join("");
      if xs.len() > res.len() {
        res = xs.to_string();
      }
    }
  }
  return res;
}

#[cfg(test)]
mod tests {
  use super::*;
  #[test]
  fn test_solution_it_works_when_k_is_2() {
    let xs = vec!["zone", "abigail", "theta", "form", "libe", "zas"];
    let res = solution(&xs, 2);
    assert_eq!(res, "abigailtheta".to_string())
  }

  #[test]
  fn test_solution_it_works_when_k_is_3() {
    let xs = vec!["zone", "abigail", "theta", "form", "libe", "zas"];
    let res = solution(&xs, 3);
    assert_eq!(res, "zoneabigailtheta".to_string())
  }
}