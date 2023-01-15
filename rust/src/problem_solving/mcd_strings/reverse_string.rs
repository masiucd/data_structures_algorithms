#[allow(dead_code)]
pub fn reverse_string(input: &str) -> String {
  let mut reversed = String::new();
  for c in input.chars().rev() {
    reversed.push(c);
  }
  reversed
}

#[allow(dead_code)]
pub fn reverse_recursive(input: &str) -> String {
  reverses_recursion_helper(input.to_string())
}

#[allow(dead_code)]
fn reverses_recursion_helper(mut s: String) -> String {
  if s.is_empty() {
    return s;
  }
  let removed_char = s.remove(0);
  let mut s = reverses_recursion_helper(s);
  s.push(removed_char);
  s
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn it_works_to_reverse_string() {
    assert_eq!(reverse_string("abc"), "cba");
  }

  #[test]
  fn it_works_to_reverse_string_recursive() {
    assert_eq!(reverse_recursive("abc"), "cba");
  }
}