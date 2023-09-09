import {LinkedList} from "../single_linked_list/list.ts"

interface Queable<T> {
  enqueue(value: T): void
  dequeue(): T | null
  peek(): T | null
  getSize(): number
  isEmpty(): boolean
}

export class Queue<T> implements Queable<T> {
  private container: LinkedList<T>

  constructor() {
    this.container = new LinkedList<T>()
  }

  enqueue(value: T): void {
    this.container.append(value)
  }
  dequeue(): T | null {
    const removed = this.container.getHead()
    this.container.delete(0)
    return removed
  }
  peek(): T | null {
    return this.container.getHead()
  }
  getSize(): number {
    return this.container.getSize()
  }
  isEmpty(): boolean {
    return this.container.getSize() === 0
  }

  print(): T[] {
    return this.container.print()
  }
}
