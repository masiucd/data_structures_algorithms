pub fn sum(xs: Vec<i32>) -> i32 {
    if xs.len() == 0 {
        return 0;
    }
    let head = xs.first().unwrap();
    let tail = &xs[1..];
    return head + sum(tail.to_vec());
}
