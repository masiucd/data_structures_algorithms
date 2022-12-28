pub fn count_bits(n: u32) -> u32 {
    let mut count: u32 = 0;
    let mut n = n;
    while n > 0 {
        count += n & 1;
        n >>= 1;
    }
    count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count_bits() {
        assert_eq!(count_bits(0), 0);
        assert_eq!(count_bits(4), 1);
        assert_eq!(count_bits(7), 3);
        assert_eq!(count_bits(9), 2);
        assert_eq!(count_bits(10), 2);
    }
}
