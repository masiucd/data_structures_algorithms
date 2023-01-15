use std::collections::HashMap;

#[allow(dead_code)]
pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
  let mut map: HashMap<String, Vec<String>> = HashMap::new();
  for s in strs {
    let mut chars: Vec<char> = s.chars().collect();
    chars.sort();
    let key = chars.into_iter().collect();
    map
      .entry(key)
      .or_insert(vec![])
      .push(s);
  }
  map.values().cloned().collect()
}