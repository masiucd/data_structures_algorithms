use std::collections::HashMap;

#[allow(dead_code)]
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

#[allow(dead_code)]
pub fn is_valid_two(s: &str) -> bool {
    let mut stack = vec![];
    for ch in s.chars() {
        match ch {
            '(' => stack.push(')'),
            '[' => stack.push(']'),
            '{' => stack.push('}'),
            x => {
                if Some(x) != stack.pop() {
                    return false;
                }
            }
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

    #[test]
    fn test_is_valid_two() {
        assert_eq!(is_valid_two("()"), true);
        assert_eq!(is_valid_two("()[]{}"), true);
        assert_eq!(is_valid_two("(]"), false);
        assert_eq!(is_valid_two("([)]"), false);
        assert_eq!(is_valid_two("{[]}"), true);
    }
}
