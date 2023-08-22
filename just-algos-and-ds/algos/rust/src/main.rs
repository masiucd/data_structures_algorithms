use rust::strings::anagram::is_anagram;

fn main() {
    let res = is_palindrome("A man, a plan, a canal: Panama".to_string());
    println!("{res}")
}

fn is_palindrome(s: String) -> bool {
    // Remove all white space and special characters from the 's' string
    let s = s
        .chars()
        .filter(|c| c.is_alphanumeric())
        .map(|c| c.to_lowercase().to_string())
        .collect::<String>();
    let reversed = s.chars().rev().collect::<String>();
    s == reversed
}
