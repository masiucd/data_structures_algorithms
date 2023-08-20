pub fn is_palindrome(s: String) -> bool {
    // Remove all white space and special characters from the 's' string
    let s = s
        .chars()
        .filter(|c| c.is_alphanumeric())
        .map(|c| c.to_lowercase().to_string())
        .collect::<String>();
    let reversed = s.chars().rev().collect::<String>();
    s == reversed
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_one() {
        let res = is_palindrome("A man, a plan, a canal: Panama".to_string());
        assert!(res);
    }

    #[test]
    fn test_two() {
        let res = !is_palindrome("Marcello Is the best".to_string());
        assert!(res);
    }
}
