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

#[allow(dead_code)]
fn reverse_str_rec(s: &str) -> String {
  if s.len() == 0 {
    return String::from("");
  }
  let head = &s[0..1];
  let tail = &s[1..];
  let res = reverse_str_rec(tail);
  format!("{}{}", res, head)
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

  #[test]
  fn it_works_to_reverse_string_rec() {
    assert_eq!(reverse_str_rec("abc"), "cba");
  }
}