use rs::algorithms::strings::is_palindrome::is_palindrome;
// use rs::data_structures::lists::lists::Node;

fn main() {
    let res = is_palindrome("racecar"); // true
    println!("{res}");
    let res = is_palindrome("marcell"); // false
    println!("{res}");
}
