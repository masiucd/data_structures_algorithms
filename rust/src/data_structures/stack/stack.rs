#[derive(Debug)]
pub struct Stack {
    stack: Vec<i32>,
}

pub trait StackPush {
    fn push(&mut self, value: i32);
    fn pop(&mut self) -> Option<i32>;
    fn peek(&self) -> Option<i32>;
}

impl StackPush for Stack {
    fn push(&mut self, value: i32) {
        self.stack.push(value);
    }

    fn pop(&mut self) -> Option<i32> {
        let last = self.stack.pop();
        match last {
            Some(last) => Some(last),
            None => None,
        }
    }

    fn peek(&self) -> Option<i32> {
        let last = self.stack.last();
        if let Some(last) = last {
            Some(*last)
        } else {
            None
        }
    }
}

pub fn new() -> Stack {
    Stack { stack: Vec::new() }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_stack_push() {
        let mut stack = new();
        stack.push(1);
        stack.push(2);
        stack.push(3);
        assert_eq!(stack.peek(), Some(3));

        let removed = stack.pop();
        assert_eq!(removed, Some(3));

        assert_eq!(stack.peek(), Some(2));
    }
}
