mod data_structures;
mod problem_solving;
fn main() {
    let xs = vec!["zone", "abigail", "theta", "form", "libe", "zas"];
    let res = problem_solving::kata::longest_consent::solution(&xs, 2);
    println!("{}", res);
}
