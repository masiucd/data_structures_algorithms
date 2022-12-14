pub fn fibonacci_dp(n: i32) -> i32 {
    let mut a = 1;
    let mut b = 1;
    for _i in 0..n - 1 {
        let temp = a;
        a = a + b;
        b = temp;
    }
    a
}
