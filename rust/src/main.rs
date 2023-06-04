// Name of the project: rust
// use rust::algos::bits::count;
// use rust::common;
use rust::math;
// use rust::problem_solving::kata;

fn main() {
  let xs = vec![1, 2, 3, 4, 5, 6, 7, 8, 9];
  let res = math::basics::median(&xs);

  println!("{res}");
}