use std::cmp::Ordering;

struct Node<T> {
    value: T,
    left: Option<Box<Node<T>>>,
    right: Option<Box<Node<T>>>,
}

impl<T: Ord> Node<T> {
    fn new(value: T) -> Self {
        Self {
            value,
            left: None,
            right: None,
        }
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_insert() {
        let mut bst = BinarySearchTree::new();
        bst.insert(5);
        bst.insert(3);
        bst.insert(7);
        bst.insert(1);
        bst.insert(4);
        bst.insert(6);
        bst.insert(8);

        let root = &bst.root.unwrap();
        assert_eq!(root.value, 5);
        assert_eq!(root.left.as_ref().unwrap().value, 3);
        assert_eq!(root.right.as_ref().unwrap().value, 7);
        assert_eq!(root.left.as_ref().unwrap().left.as_ref().unwrap().value, 1);
        assert_eq!(root.left.as_ref().unwrap().right.as_ref().unwrap().value, 4);
        assert_eq!(root.right.as_ref().unwrap().left.as_ref().unwrap().value, 6);
        assert_eq!(
            root.right.as_ref().unwrap().right.as_ref().unwrap().value,
            8
        );
    }
}
