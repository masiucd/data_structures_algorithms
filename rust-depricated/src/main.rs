// Name of the project: rust
// use rust::algos::bits::count;
// use rust::common;
pub mod strings;

// use rust::problem_solving::strings;

fn main() {
  // let res = strings::pig_latin("apple");
  // let res = strings::pig_latin("first");
  // let res = strings::pig_latin("Здравствуйте");

  let res = is_anagram("anagram", "nagaram");
  println!("{res}");
  let res = is_anagram("car", "ram");
  println!("{res}");
}

fn is_anagram(s1: &str, s2: &str) -> bool {
  s1.chars().into_iter().collect::<Vec<char>>().sort() ==
    s2.chars().into_iter().collect::<Vec<char>>().sort()
}
