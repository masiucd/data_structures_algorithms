use std::collections::HashMap;
pub fn is_valid(s: &str) -> bool {
    let mut stack: Vec<char> = Vec::new();
    let closing = HashMap::from([
        ('}', '{'),
        (']', '['),
        (')', '('),
    ]);

    for c in s.chars() {
        if closing.contains_key(&c) {
            if !stack.is_empty() && closing.get(&c).unwrap() == stack.get(stack.len() - 1).unwrap() {
                stack.pop();
            } else {
                return false;
            }
        } else {
            stack.push(c);
        }
    }
    return stack.is_empty();
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
        assert!(is_valid("()"));
        assert!(is_valid("()[]{}"));
    }
}
