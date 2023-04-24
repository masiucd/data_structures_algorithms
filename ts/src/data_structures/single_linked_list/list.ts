class Node<T> {
  value: T
  next: Node<T> | null
  constructor(value: T) {
    this.value = value
    this.next = null
  }
}

interface ListAble<T> {
  prepend(value: T): void
  append(value: T): void
  get(index: number): T | null
  delete(index: number): void
  insert(index: number, value: T): void
  print(): T[]
  getSize(): number
  getHead(): T | null
  getTail(): T | null
}

export class LinkedList<T> implements ListAble<T> {
  private head: Node<T> | null = null
  private tail: Node<T> | null = null
  private len = 0

  constructor() {
    this.head = null
    this.tail = null
    this.len = 0
  }

  prepend(value: T): void {
    const newNode: Node<T> = new Node(value)
    if (!this.head) {
      this.head = newNode
      this.tail = newNode
    } else {
      newNode.next = this.head
      this.head = newNode
    }
    this.len++
  }
  append(value: T): void {
    const newNode: Node<T> = new Node(value)
    if (!this.head) {
      this.head = newNode
      this.tail = newNode
    } else {
      this.tail!.next = newNode
      this.tail = newNode
    }
    this.len++
  }
  get(index: number): T | null {
    throw new Error("Method not implemented.")
  }
  delete(index: number): void {
    throw new Error("Method not implemented.")
  }
  insert(index: number, value: T): void {
    throw new Error("Method not implemented.")
  }
  print(): T[] {
    const result: T[] = []
    let current = this.head
    while (current) {
      result.push(current.value)
      current = current.next
    }
    return result
  }
  getSize(): number {
    return this.len
  }
  getHead(): T | null {
    return this.head ? this.head.value : null
  }
  getTail(): T | null {
    return this.tail ? this.tail.value : null
  }
}
