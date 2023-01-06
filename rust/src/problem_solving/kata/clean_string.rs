pub fn solution(input: &str) -> String {
    let mut result = String::new();
    for c in input.chars() {
        if c == '#' {
            if result.len() > 0 {
                result.pop();
            }
        } else {
            result.push(c);
        }
    }
    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        assert_eq!(solution("abc#d##c"), "ac");
        assert_eq!(solution("abc####d##c#"), "");
        assert_eq!(solution("#cwks"), "cwks");
        assert_eq!(solution("#####"), "");
    }
}
