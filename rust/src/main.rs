use std::collections::HashMap;

mod algo;
mod data_structures;
mod problem_solving;

fn main() {
    println!(
        "{}",
        problem_solving::kata::clean_string::solution("abc#d##c")
    );

    println!(
        "{:?}",
        problem_solving::kata::group_anagrams::group_anagrams(vec![
            "eat".to_string(),
            "tea".to_string(),
            "tan".to_string(),
            "ate".to_string(),
            "nat".to_string(),
            "bat".to_string()
        ])
    );
}
