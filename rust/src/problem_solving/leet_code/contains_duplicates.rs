use std::collections::HashMap;

pub fn contains_duplicates(nums: &Vec<i32>) -> bool {
    let mut store = HashMap::new();
    for n in nums {
        *store.entry(n).or_insert(0) += 1;
    }
    !store.values().all(|x| x == &1)
}

#[cfg(test)]

mod tests {
    use super::*;

    #[test]
    fn test_contains_duplicates() {
        let res = contains_duplicates(&vec![1, 2, 3, 4, 5, 2, 1]);
        assert_eq!(res, true);
        let res = contains_duplicates(&vec![1, 2, 3]);
        assert_eq!(res, false);
    }
}
