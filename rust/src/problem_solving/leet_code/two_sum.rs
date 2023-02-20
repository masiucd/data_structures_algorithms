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
