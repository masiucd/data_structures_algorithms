pub fn longest_common_prefix(strs: Vec<&str>) -> String {
  let mut res = String::new();
  let first = strs.get(0).unwrap();
  for (i, char) in first.chars().enumerate() {
    for s in strs.iter() {
      if s.chars().nth(i).unwrap() != char {
        return res;
      }
    }
    res.push(char);
  }
  res
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_longest_common_prefix() {
    let xs = vec!["flower", "flow", "flyer"];
    let res = longest_common_prefix(xs);
    assert_eq!(res, "fl");

    let xs = vec!["dog", "racecar", "car"];
    let res = longest_common_prefix(xs);
    assert_eq!(res, "");

    let xs = vec!["dog", "dog", "dobb"];
    let res = longest_common_prefix(xs);
    assert_eq!(res, "do");
  }
}