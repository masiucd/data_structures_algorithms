use std::cmp::Ordering;

struct Node<T> {
  value: T,
  left: Option<Box<Node<T>>>,
  right: Option<Box<Node<T>>>,
}

impl<T: Ord> Node<T> {
  fn new(value: T) -> Self {
    Self { value, left: None, right: None }
  }

  fn insert(&mut self, value: T) {
    match value.cmp(&self.value) {
      Ordering::Less => {
        if let Some(left) = &mut self.left {
          left.insert(value);
        } else {
          self.left = Some(Box::new(Node::new(value)));
        }
      }
      Ordering::Greater => {
        if let Some(right) = &mut self.right {
          right.insert(value);
        } else {
          self.right = Some(Box::new(Node::new(value)));
        }
      }
      Ordering::Equal => {}
    }
  }
}

struct BinarySearchTree<T> {
  root: Option<Box<Node<T>>>,
}

impl<T: Ord> BinarySearchTree<T> {
  fn new() -> Self {
    Self { root: None }
  }

  fn insert(&mut self, value: T) {
    if let Some(root) = &mut self.root {
      root.insert(value);
    } else {
      self.root = Some(Box::new(Node::new(value)));
    }
  }
}