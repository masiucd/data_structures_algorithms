pub fn target_sum(xs: Vec<i32>, size: u8) -> i32 {
    let mut temp_sum = 0;

    for i in 0..size {
        let val = xs[i as usize];
        temp_sum += val;
    }
    let mut max_sum = temp_sum;
    for i in size..xs.len() as u8 {
        temp_sum += xs[i as usize] - xs[(i - size) as usize];
        max_sum = std::cmp::max(temp_sum, max_sum);
    }

    max_sum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_target_sum() {
        let xs = vec![1, 10, 1, 2, 3];
        let size = 3;
        let result = target_sum(xs, size);
        assert_eq!(result, 14);
    }

    #[test]
    fn test_target_sum_2() {
        let xs = vec![1, 10, 1, 2, 3];
        let size = 2;
        let result = target_sum(xs, size);
        assert_eq!(result, 11);
    }
}
