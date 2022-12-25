mod algo;
mod data_structures;
mod problem_solving;

fn main() {
  println!("{:?}", problem_solving::kata::positive_sum::solution(&vec![1, 2, 3, 4, -2]));
}