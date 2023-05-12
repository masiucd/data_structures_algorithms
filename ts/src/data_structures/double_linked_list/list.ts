interface DoubleLinkedListAble<T> {
  append(value: T): void;
  prepend(value: T): void;
  insert(value: T, index: number): void;
  remove(index: number): void;
  removeFirst(): void;
  removeLast(): void;
  get(index: number): T;
  getFirst(): T | null;
  getLast(): T | null;
  size(): number;
  isEmpty(): boolean;
  toArray(): T[];
  toString(): string;
}

export class Node<T> {
  value: T;
  next: Node<T> | null;
  prev: Node<T> | null;
  constructor(value: T) {
    this.value = value;
    this.next = null;
    this.prev = null;
  }
}

export class List<T> implements DoubleLinkedListAble<T> {
  private head: Node<T> | null;
  private tail: Node<T> | null;
  private len: number;

  constructor() {
    this.head = null;
    this.tail = null;
    this.len = 0;
  }

  append(value: T): void {
    const node = new Node(value);
    if (!this.head) {
      this.head = node;
      this.tail = node;
      this.len++;
      return;
    }
    if (this.tail) {
      this.tail.next = node;
      node.prev = this.tail;
      this.tail = node;
      this.len++;
    }
  }
  prepend(value: T): void {
    const node = new Node(value);
    if (!this.head) {
      this.head = node;
      this.tail = node;
      this.len++;
      return;
    }
    node.next = this.head;
    this.head.prev = node;
    this.head = node;
    this.len++;
  }
  insert(value: T, index: number): void {
    throw new Error("Method not implemented.");
  }
  remove(index: number): void {
    throw new Error("Method not implemented.");
  }
  removeFirst(): void {
    throw new Error("Method not implemented.");
  }
  removeLast(): void {
    throw new Error("Method not implemented.");
  }
  get(index: number): T {
    throw new Error("Method not implemented.");
  }
  getFirst(): T | null {
    if (this.head) {
      return this.head.value;
    }
    return null;
  }
  getLast(): T | null {
    if (this.tail) {
      return this.tail.value;
    }
    return null;
  }
  size(): number {
    return this.len;
  }
  isEmpty(): boolean {
    throw new Error("Method not implemented.");
  }
  toArray(): T[] {
    const xs: T[] = [];
    let node = this.head;
    while (node) {
      xs.push(node.value);
      node = node.next;
    }
    return xs;
  }
  toString(): string {
    throw new Error("Method not implemented.");
  }
}
