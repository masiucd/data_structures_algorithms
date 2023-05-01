#[allow(dead_code)]
pub struct StackWithList<T> {
  stack: std::collections::LinkedList<T>,
}

pub trait StackWithListAble<T> {
  fn new() -> Self;
  fn push(&mut self, value: T);
  fn pop(&mut self) -> Option<T>;
  fn peek(&self) -> Option<&T>;
  fn is_empty(&self) -> bool;
}

impl<T> StackWithListAble<T> for StackWithList<T> {
  fn new() -> Self {
    StackWithList { stack: std::collections::LinkedList::new() }
  }

  fn push(&mut self, value: T) {
    self.stack.push_back(value);
  }

  fn pop(&mut self) -> Option<T> {
    self.stack.pop_back()
  }

  fn peek(&self) -> Option<&T> {
    self.stack.back()
  }

  fn is_empty(&self) -> bool {
    self.stack.is_empty()
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn it_works_to_create_a_stack_with_list() {
    let stack = StackWithList::<i32>::new();
    assert!(stack.is_empty());
  }

  #[test]
  fn it_works_to_push_to_a_stack_with_list() {
    let mut stack = StackWithList::<i32>::new();
    stack.push(1);
    assert!(!stack.is_empty());
  }

  #[test]
  fn it_works_to_pop_from_a_stack_with_list() {
    let mut stack = StackWithList::<i32>::new();
    stack.push(1);
    let res = stack.pop();
    assert_eq!(res, Some(1));
  }

  #[test]
  fn it_works_to_peek_from_a_stack_with_list() {
    let mut stack = StackWithList::<i32>::new();
    stack.push(1);
    let res = stack.peek();
    assert_eq!(res, Some(&1));
  }
}