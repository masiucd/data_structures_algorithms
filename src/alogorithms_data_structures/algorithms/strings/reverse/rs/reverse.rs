fn reverse_string(s: &mut Vec<char>) {
    let mut i = 0;
    let mut j = s.len() - 1;
    while i < j {
        s.swap(i, j);
        i += 1;
        j -= 1;
    }
}

fn reverse_string_two(s: &mut Vec<char>) -> String {
    s.reverse();
    s.iter().collect()
}
