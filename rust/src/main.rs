mod algo;
mod data_structures;
mod problem_solving;

fn main() {
  let nums = vec![1, 2, 3, 1];
  let k = 3;
  let result = problem_solving::kata::contains_nearby_duplicate::contains_nearby_duplicate(nums, k);
  println!("result: {}", result);
}