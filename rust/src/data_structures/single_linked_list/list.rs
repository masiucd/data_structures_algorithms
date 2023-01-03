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
