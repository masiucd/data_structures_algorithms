pub fn int_to_roman(num: i32) -> String {
  let values = vec![1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
  let symbols = vec!["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
  let mut res = String::new();
  let mut num = num;
  for (i, n) in values.iter().enumerate() {
    while num >= *n {
      res.push_str(symbols[i]);
      num -= n;
    }
  }
  res
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_int_to_roman() {
    let res = int_to_roman(1000);
    assert_eq!(res, "M");
  }

  #[test]
  fn test_int_to_roman_2() {
    let res = int_to_roman(1990);
    assert_eq!(res, "MCMXC");
  }
}