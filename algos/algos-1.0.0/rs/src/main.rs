use rs::algorithms::strings::is_palindrome::is_palindrome;
use rs::algorithms::strings::count_vowels::{ count_vowels, count_vowels_rec };
// use rs::data_structures::lists::lists::Node;

fn main() {
    let r = count_vowels_rec("hello");
    println!("count_vowels: {}", r);
}
