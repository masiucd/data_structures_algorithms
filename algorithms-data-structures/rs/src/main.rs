use std::collections::HashMap;

// use rs::algorithms::strings::is_palindrome::is_palindrome;
use rs::algorithms::stacks::is_valid::is_valid;
use rs::algorithms::arrays::sliding_window::target_sum::target_sum;
// use rs::data_structures::lists::lists::Node;

fn main() {
    let xs = vec![1, 10, 1, 2, 3];
    let size = 3;
    let result = target_sum(xs, size);
    println!("result: {}", result);
}
