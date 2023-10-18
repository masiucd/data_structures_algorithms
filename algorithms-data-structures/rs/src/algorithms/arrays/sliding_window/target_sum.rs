pub fn target_sum(xs: Vec<i32>, size: u8) -> i32 {
    let mut temp_sum = 0;

    for i in 0..size {
        let val = xs[i as usize];
        temp_sum += val;
    }
    let mut max_sum = temp_sum;
    for i in size..xs.len() as u8 {
        temp_sum += xs[i as usize] - xs[(i - size) as usize];
        max_sum = std::cmp::max(temp_sum, max_sum);
    }

    max_sum
}
