mod algo;
mod data_structures;
mod problem_solving;

fn main() {
  let res = problem_solving::leet_code::solutions::Solution::length_of_last_word(
    "Hello World".to_string()
  );
  println!("{}", res);
}