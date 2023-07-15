use rust::strings::is_anagram::is_anagram;

fn main() {
    let s = String::from("anagram");
    let t = String::from("nagaram");
    println!("{}", is_anagram(s, t));
}
