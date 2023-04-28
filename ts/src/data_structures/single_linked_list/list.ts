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
  reverse(): void
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
  reverse(): void {
    if (!this.head) return
    let prev = null
    let current = this.head as Node<T> | null
    let next = null
    while (current !== null) {
      next = current.next
      current.next = prev
      prev = current
      current = next
    }
    this.tail = this.head
    this.head = prev
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
    if (!this.head) return null
    if (index < 0 || index >= this.len) throw new Error("Index out of range")
    let current: Node<T> | null = this.head
    let count = 0
    while (current !== null) {
      if (index === count) {
        return current.value
      }
      count++
      current = current.next
    }
    return null
  }
  delete(index: number): void {
    if (!this.head) return
    if (index < 0 || index >= this.len) throw new Error("Index out of range")
    if (index === 0) {
      this.head = this.head!.next
      this.tail = this.head === null ? null : this.tail
      this.len--
      return
    }
    if (index === this.getSize() - 1) {
      const prevTail = this.getNode(index - 1) as Node<T>
      prevTail.next = null
      this.tail = prevTail
      this.len--
      return
    }
    const prevNode = this.getNode(index - 1) as Node<T>
    let nodeToRemove = prevNode.next
    prevNode.next = nodeToRemove!.next
    nodeToRemove = null
    this.len--
  }
  insert(index: number, value: T): void {
    if (index < 0) {
      throw new Error("Index out of range")
    }
    if (index === 0) {
      this.prepend(value)
      return
    }
    if (index >= this.getSize() - 1) {
      this.append(value)
      return
    }
    const newNode = new Node(value)
    const prevNode = this.getNode(index - 1) as Node<T>
    newNode.next = prevNode.next
    prevNode.next = newNode
    this.len++
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

  private getNode(index: number): Node<T> | null {
    if (!this.head) return null
    if (index < 0 || index >= this.len) throw new Error("Index out of range")
    let current: Node<T> | null = this.head
    let count = 0
    while (current !== null) {
      if (index === count) {
        return current
      }
      count++
      current = current.next
    }
    return null
  }
}
