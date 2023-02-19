mod algo;
mod data_structures;
mod problem_solving;

fn main() {
  let res = problem_solving::leet_code::single_number::single_number(vec![4, 1, 2, 1, 2]);
  println!("res: {}", res);
}