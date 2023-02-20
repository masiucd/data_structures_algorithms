pub fn two_sum(nums: Vec<i32>, target: i32) -> (usize, usize) {
    let mut store = std::collections::HashMap::new();
    for (i, v) in nums.iter().enumerate() {
        if store.contains_key(&(target - v)) {
            let stored_index = store.get(&(target - v)).unwrap();
            return (*stored_index, i);
        }
        store.insert(v, i);
    }
    return (0, 0);
}

#[cfg(test)]

mod tests {

    use super::*;

    #[test]

    fn test_two_sum() {
        assert_eq!(two_sum(vec![2, 3, 1, 4], 6), (0, 3));
        assert_eq!(two_sum(vec![2, 3, 1, 4], 3), (0, 2));
        assert_eq!(two_sum(vec![2, 3, 1, 5], 7), (0, 3));
        assert_eq!(two_sum(vec![2, 3, 1, 5], 4), (1, 2));
    }
}
