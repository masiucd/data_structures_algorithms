// Name of the project: rust
// use rust::algos::bits::count;
// use rust::common;

use rust::problem_solving::strings;

fn main() {
  let res = strings::pig_latin("apple");
  let res = strings::pig_latin("first");
  let res = strings::pig_latin("Здравствуйте");
  println!("{res}");
}

// Convert strings to pig latin.
//  The first consonant of each word is moved to the end of the word and ay is added,
//   so first becomes irst-fay. Words that start with a vowel have hay added to the end instead (apple becomes apple-hay).
//    Keep in mind the details about UTF-8 encoding!
