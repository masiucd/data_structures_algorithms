use std::collections::HashMap;

#[allow(dead_code)]
pub fn contains_duplicates(nums: &Vec<i32>) -> bool {
    let mut store = HashMap::new();
    for n in nums {
        *store.entry(n).or_insert(0) += 1;
    }
    !store.values().all(|x| x == &1)
}

#[allow(dead_code)]
pub fn contains_duplicates_two(nums: &Vec<i32>) -> bool {
    let mut store = HashMap::new();
    for n in nums {
        if store.contains_key(n) {
            return true;
        }
        store.insert(n, 1);
    }
    false
}

#[allow(dead_code)]
pub fn contains_duplicates_three(nums: &Vec<i32>) -> bool {
    let mut set = std::collections::HashSet::new();
    for n in nums {
        set.insert(n);
    }
    set.len() != nums.len()
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

    #[test]
    fn test_contains_duplicates_two() {
        let res = contains_duplicates_two(&vec![1, 2, 3, 4, 5, 2, 1]);
        assert_eq!(res, true);
        let res = contains_duplicates_two(&vec![1, 2, 3]);
        assert_eq!(res, false);
    }

    #[test]
    fn test_contains_duplicates_three() {
        let res = contains_duplicates_three(&vec![1, 2, 3, 4, 5, 2, 1]);
        assert_eq!(res, true);
        let res = contains_duplicates_three(&vec![1, 2, 3]);
        assert_eq!(res, false);
    }
}
