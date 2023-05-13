pub fn concat_strings(a: &str, b: &str) -> String {
  let mut result = String::new();
  result.push_str(a);
  result.push_str(b);
  result
}