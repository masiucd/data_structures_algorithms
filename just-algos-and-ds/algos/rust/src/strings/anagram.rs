pub fn is_anagram(s: String, t: String) -> bool {
    let mut s = s.chars().collect::<Vec<char>>();
    let mut t = t.chars().collect::<Vec<char>>();
    s.sort();
    t.sort();
    s == t
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_anagram() {
        let s = String::from("anagram");
        let t = String::from("nagaram");
        assert_eq!(is_anagram(s, t), true);

        let s = String::from("rat");
        let t = String::from("car");
        assert_eq!(is_anagram(s, t), false);
    }
}
