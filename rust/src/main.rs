use crate::problem_solving::leet_code::count_bits;

mod algo;
mod data_structures;
mod problem_solving;

fn main() {
  let res = count_bits::count_bits(5);
  // print as binary
  println!("binary: {:b}", 5);
  println!("res: {}", res);
}