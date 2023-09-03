fn is_vowel(c: char) -> bool {
    match c {
        'a' | 'e' | 'i' | 'o' | 'u' => true,
        _ => false,
    }
}

pub fn count_vowels(s: &str) -> usize {
    s.chars()
        .filter(|&c| is_vowel(c))
        .count()
}

pub fn count_vowels_rec(s: &str) -> usize {
    if s.is_empty() {
        return 0;
    }
    let c = s.chars().nth(0).unwrap();
    let count = if is_vowel(c) { 1 } else { 0 };
    return count + count_vowels_rec(&s[1..]);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count_vowels() {
        assert_eq!(count_vowels("hello"), 2);
        assert_eq!(count_vowels("why"), 0);
        assert_eq!(count_vowels("mississippi"), 4);
    }

    #[test]
    fn test_count_vowels_rec() {
        assert_eq!(count_vowels_rec("hello"), 2);
        assert_eq!(count_vowels_rec("why"), 0);
        assert_eq!(count_vowels_rec("mississippi"), 4);
    }
}
