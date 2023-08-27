pub mod lists {
    pub struct Node<T> {
        pub data: T,
        pub next: Option<Box<Node<T>>>,
    }
}
