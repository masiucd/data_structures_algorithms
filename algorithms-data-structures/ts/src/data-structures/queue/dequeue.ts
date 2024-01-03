export class Deque {
  private items: number[];

  constructor() {
    this.items = [];
  }

  isEmpty(): boolean {
    return this.items.length === 0;
  }

  append(value: number): void {
    this.items.push(value);
  }

  appendleft(value: number): void {
    this.items.unshift(value);
  }

  pop(): number {
    let removedItem = this.items.pop();
    if (removedItem) return removedItem;
    return -1;
  }

  popleft(): number | undefined {
    let removedItem = this.items.shift();
    if (removedItem) return removedItem;
    return -1;
  }
}
