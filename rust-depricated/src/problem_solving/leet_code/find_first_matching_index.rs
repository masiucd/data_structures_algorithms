#[allow(dead_code)]
pub fn str_str(haystack: &str, needle: &str) -> i32 {
    if needle.is_empty() {
        return 0;
    }
    for (i, _) in haystack.char_indices() {
        if haystack[i..].starts_with(needle) {
            return i as i32;
        }
    }

    return -1 as i32;
}

#[cfg(test)]

mod tests {
    use super::*;

    #[test]
    fn test_str_str() {
        assert_eq!(str_str("hello", "ll"), 2);
        assert_eq!(str_str("aaaaa", "bba"), -1);
        assert_eq!(str_str("", ""), 0);
        assert_eq!(str_str("a", ""), 0);
    }
}
