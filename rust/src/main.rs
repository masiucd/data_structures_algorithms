use data_structures::stack::stack::StackPush;

// mod problem_solving;
// mod data_structures;
mod data_structures;
fn main() {
    // let x = problem_solving::mcd_strings::reverse_string::reverse_recursive("abc");
    // let xs = vec![6, 1, 3, 3, 3, 6, 6];
    // let res = problem_solving::daily_coding_problems::find_single_digit::solution(xs);
    // println!("{:?}", res)
    // data_structure
    let mut stack = data_structures::stack::stack::new();
    stack.push(100);
    println!("{:?}", stack)
}
