use std::collections::HashMap;

pub fn is_valid(s: &str) -> bool {
    let mut stack = Vec::new();
    let mut map = HashMap::new();
    map.insert(')', '(');
    map.insert(']', '[');
    map.insert('}', '{');
    for c in s.chars() {
        if map.contains_key(&c) {
            if stack.len() > 0 && stack[stack.len() - 1] == map[&c] {
                stack.pop();
            } else {
                return false;
            }
        } else {
            stack.push(c);
        }
    }
    stack.is_empty()
}

#[cfg(test)]

mod tests {
    use super::*;

    #[test]
    fn test_is_valid() {
        assert_eq!(is_valid("()"), true);
        assert_eq!(is_valid("()[]{}"), true);
        assert_eq!(is_valid("(]"), false);
        assert_eq!(is_valid("([)]"), false);
        assert_eq!(is_valid("{[]}"), true);
    }
}
