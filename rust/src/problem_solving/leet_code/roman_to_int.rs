use std::collections::HashMap;

pub fn roman_letters_map() -> HashMap<char, i32> {
  let mut roman = HashMap::new();
  roman.insert('I', 1);
  roman.insert('V', 5);
  roman.insert('X', 10);
  roman.insert('L', 50);
  roman.insert('C', 100);
  roman.insert('D', 500);
  roman.insert('M', 1000);
  roman
}

pub fn roman_to_int(s: String) -> i32 {
  let roman = roman_letters_map();
  let mut result = 0;
  for (i, c) in s.chars().enumerate() {
    if i + 1 < s.len() {
      let next = s
        .chars()
        .nth(i + 1)
        .unwrap();
      if roman[&c] < roman[&next] {
        result -= roman[&c];
      } else {
        result += roman[&c];
      }
    } else {
      result += roman[&c];
    }
  }
  result
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_roman_to_int() {
    let s = String::from("III");
    let res = roman_to_int(s);
    assert_eq!(res, 3);

    let s = String::from("IV");
    let res = roman_to_int(s);
    assert_eq!(res, 4);

    let s = String::from("IX");
    let res = roman_to_int(s);
    assert_eq!(res, 9);

    let s = String::from("LVIII");
    let res = roman_to_int(s);
    assert_eq!(res, 58);

    let s = String::from("MCMXCIV");
    let res = roman_to_int(s);
    assert_eq!(res, 1994);
  }
}