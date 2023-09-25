use std::collections::HashMap;

// use rs::algorithms::strings::is_palindrome::is_palindrome;
use rs::algorithms::stacks::is_valid::is_valid;
// use rs::data_structures::lists::lists::Node;

fn main() {
    let mut res = is_valid("()");
    println!("{res}");
    res = is_valid("()[]{}");
    println!("{res}");

    let closing = HashMap::from([
        ('}', '{'),
        (']', '['),
        (')', '('),
    ]);

    let f = closing.contains_key(&'}');
    println!("{f}");
}
