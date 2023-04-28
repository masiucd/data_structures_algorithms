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
    throw new Error("Method not implemented.")
  }
  pop(): T | null {
    throw new Error("Method not implemented.")
  }
  peek(): T | null {
    throw new Error("Method not implemented.")
  }
  getSize(): number {
    throw new Error("Method not implemented.")
  }
  isEmpty(): boolean {
    throw new Error("Method not implemented.")
  }
}
