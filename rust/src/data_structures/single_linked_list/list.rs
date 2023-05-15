pub struct Node {
  data: i32,
  next: Option<Box<Node>>,
}

pub trait List {
  fn new() -> Self;
  fn append(&mut self, data: i32);
  fn prepend(&mut self, data: i32);
  fn remove(&mut self, data: i32);
  fn get(&self, index: usize) -> Option<&i32>;
  fn print(&self);
}

pub struct SingleLinkedList {
  head: Option<Box<Node>>,
  tail: Option<Box<Node>>,
  length: usize,
}

impl List for SingleLinkedList {
  fn new() -> Self {
    SingleLinkedList {
      head: None,
      tail: None,
      length: 0,
    }
  }

  fn append(&mut self, data: i32) {
    let new_node = Box::new(Node {
      data,
      next: None,
    });

    match self.tail.take() {
      Some(old_tail) => {
        old_tail.next = Some(new_node);
      }
      None => {
        self.head = Some(new_node);
      }
    }

    self.length += 1;
    self.tail = Some(new_node);
  }

  fn prepend(&mut self, data: i32) {
    let new_node = Box::new(Node {
      data,
      next: self.head.take(),
    });

    if self.head.is_none() {
      self.tail = Some(new_node.clone());
    }

    self.head = Some(new_node);
    self.length += 1;
  }

  fn remove(&mut self, data: i32) {
    let mut current_node = self.head.take();

    if let Some(mut node) = current_node {
      if node.data == data {
        self.head = node.next.take();
        self.length -= 1;
        return;
      }

      while let Some(next_node) = node.next.take() {
        if next_node.data == data {
          node.next = next_node.next;
          self.length -= 1;
          break;
        }
        node = next_node;
      }
    }
  }

  fn get(&self, index: usize) -> Option<&i32> {
    if index >= self.length {
      return None;
    }

    let mut current_node = self.head.as_ref();

    for _ in 0..index {
      current_node = current_node.unwrap().next.as_ref();
    }

    current_node.map(|node| &node.data)
  }

  fn print(&self) {
    let mut current_node = self.head.as_ref();

    while let Some(node) = current_node {
      println!("{}", node.data);
      current_node = node.next.as_ref();
    }
  }
}