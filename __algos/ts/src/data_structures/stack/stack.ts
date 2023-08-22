import {LinkedList} from "../single_linked_list/list.ts"

interface StackAble<T> {
  push(value: T): void
  pop(): T | null
  peek(): T | null
  getSize(): number
  isEmpty(): boolean
}

export class Stack<T> implements StackAble<T> {
  private container: LinkedList<T>

  constructor() {
    this.container = new LinkedList<T>()
  }
  push(value: T): void {
    this.container.append(value)
  }
  pop(): T | null {
    const removed = this.container.getTail()
    this.container.delete(this.container.getSize() - 1)
    return removed
  }
  peek(): T | null {
    return this.container.getTail()
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
