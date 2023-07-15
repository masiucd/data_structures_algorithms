pub mod strings {
  pub fn concat_strings(a: &str, b: &str) -> String {
    let mut result = String::new();
    result.push_str(a);
    result.push_str(b);
    result
  }

  pub fn reverse_string(input: &str) -> String {
    let mut reversed = String::new();
    for c in input.chars().rev() {
      reversed.push(c);
    }
    reversed
  }

  pub fn reverse_recursive(input: &str) -> String {
    reverses_recursion_helper(input.to_string())
  }

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
}

pub mod trees {
  #[derive(Debug)]
  pub struct Item {
    id: u8,
    title: String,
    parent_id: Option<u8>,
  }

  #[derive(Debug, Clone)]
  #[allow(dead_code)]
  pub struct TreeNode {
    id: u8,
    title: String,
    parent_id: Option<u8>,
    children: Vec<TreeNode>,
  }

  pub fn build_tree(xs: &Vec<Item>, parent_id: Option<u8>) -> Vec<TreeNode> {
    xs.iter()
      .filter(|x| x.parent_id == parent_id)
      .map(|x| TreeNode {
        title: x.title.clone(),
        id: x.id,
        parent_id: x.parent_id,
        children: build_tree(xs, Some(x.id)),
      })
      .collect()
  }

  #[cfg(test)]
  mod tests {
    use super::*;
    #[test]
    fn test_build_tree() {
      let xs: Vec<Item> = vec![
        Item { id: 1, title: "Europe".to_string(), parent_id: None },
        Item { id: 2, title: "Asia".to_string(), parent_id: None },
        Item { id: 3, title: "France".to_string(), parent_id: Some(1) },
        Item { id: 4, title: "Thailand".to_string(), parent_id: Some(2) },
        Item { id: 5, title: "Paris".to_string(), parent_id: Some(3) }
      ];
    }
  }
}