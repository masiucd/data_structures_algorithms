// Name of the project: rust
use rust::algos::bits::count;
use rust::common;
use rust::math;

fn main() {
  let xs = math::prime::prime_nums_from_range(1, 10);
  println!("{:?}", xs);
  let r = math::prime::is_prime(4);
  println!("{}", r);
}