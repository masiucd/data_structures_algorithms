mod algo;
mod data_structures;
mod problem_solving;

use rust::algo::search::binary_search;

fn main() {
  let xs = vec![1, 2, 3, 4, 5, 6, 7, 8, 9];
  let res = binary_search(&xs, 5);
  println!("{:?}", res);
}