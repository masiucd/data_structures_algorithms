mod problem_solving;

fn main() {
    let x = problem_solving::mcd_strings::reverse_string::reverse_recursive("abc");
    let xs = vec![6, 1, 3, 3, 3, 6, 6];
    let res = problem_solving::daily_coding_problems::find_single_digit::solution(xs);
    println!("{:?}", res)
}
