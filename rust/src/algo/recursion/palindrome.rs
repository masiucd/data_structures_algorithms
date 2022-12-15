pub fn is_palindrome(s: String) -> bool {
    if s.len() == 0 || s.len() == 1 {
        return true;
    }
    let start = s.chars().nth(0).unwrap();
    let end = s.chars().nth(s.len() - 1).unwrap();
    let middle: String = s.chars().skip(1).take(s.len() - 2).collect();
    return start == end && is_palindrome(middle.to_string());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_palindrome_it_works() {
        assert_eq!(is_palindrome("".to_string()), true);
        assert_eq!(is_palindrome("racecar".to_string()), true);
        assert_eq!(is_palindrome("otto".to_string()), true);
    }

    #[test]
    fn test_is_palindrome_it_fails() {
        assert_eq!(is_palindrome("cwks".to_string()), false);
        assert_eq!(is_palindrome("1916".to_string()), false);
    }
}
