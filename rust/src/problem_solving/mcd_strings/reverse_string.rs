pub fn reverse_string(input: &str) -> String {
    let mut reversed = String::new();
    for c in input.chars().rev() {
        reversed.push(c);
    }
    reversed
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works_toReverse_string() {
        assert_eq!(reverse_string("abc"), "cba");
    }
}