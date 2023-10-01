export class MinStack {
  stack: number[];
  minStack: number[];
  size: number;

  constructor() {
    this.stack = [];
    this.minStack = [];
    this.size = 0;
  }

  push(n: number): void {
    if (this.size === 0) {
      this.stack.push(n);
      this.minStack.push(n);
    } else {
      let currentMin = this.minStack.at(-1)!;
      if (n < currentMin) {
        this.minStack.push(n);
      } else {
        this.minStack.push(currentMin);
      }
      this.stack.push(n);
    }
    this.size++;
  }
  pop(): void {
    this.stack.pop();
    this.minStack.pop();
    this.size--;
  }

  top(): number {
    return this.stack.at(-1)!;
  }
  getMin(): number {
    return this.minStack.at(-1)!;
  }
}
