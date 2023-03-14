pub fn plus_one(xs: Vec<i32>) -> Vec<i32> {
    let mut result = xs.clone();
    for i in (0..result.len()).rev() {
        if result[i] < 9 {
            result[i] += 1;
            return result;
        }
        result[i] = 0;
    }
    result = vec![0; result.len() + 1];
    result[0] = 1;
    result
}

#[cfg(test)]

mod tests {
    use super::*;

    #[test]
    fn test_plus_one() {
        assert_eq!(plus_one(vec![1, 2, 3]), vec![1, 2, 4]);
        assert_eq!(plus_one(vec![9, 9, 9]), vec![1, 0, 0, 0]);
    }
}
