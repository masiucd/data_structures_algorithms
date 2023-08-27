#[allow(dead_code)]
pub fn is_palindrome(input: &str) -> bool {
    let mut start: usize = 0;
    let mut end = input.len() - 1;
    while start < end {
        if input.chars().nth(start).unwrap() == input.chars().nth(end).unwrap() {
            start += 1;
            end -= 1;
        } else {
            return false;
        }
    }
    true
}

#[allow(dead_code)]
fn is_palindrome2(input: &str) -> bool {
    return input.chars().rev().collect::<String>() == input;
}

#[cfg(test)]
mod tests {
    use super::is_palindrome;
    #[test]
    fn test_is_palindrome() {
        assert_eq!(is_palindrome("racecar"), true);
        assert_eq!(is_palindrome("marcell"), false);
    }

    use super::is_palindrome2;
    #[test]
    fn test_is_palindrome2() {
        assert_eq!(is_palindrome2("racecar"), true);
        assert_eq!(is_palindrome2("marcell"), false);
    }
}
