export class Node {
  value: number;
  next: Node | null;

  constructor(value: number) {
    this.value = value;
    this.next = null;
  }
}

export function merge(t1: Node | null, t2: Node | null): Node | null {
  if (!t1) return t2;
  if (!t2) return t1;

  let dummy = new Node(0);
  let tail = dummy;

  while (t1 !== null && t2 !== null) {
    if (t1.value < t2.value) {
      tail.next = t1;
      t1 = t1.next;
    } else {
      tail.next = t2;
      t2 = t2.next;
    }
    tail = tail.next;
  }

  if (t1 != null) {
    tail.next = t1;
  } else {
    tail.next = t2;
  }
  return dummy.next;
}

export function printList(list: Node | null) {
  let xs = [];
  while (list !== null) {
    xs.push(list.value);
    list = list.next;
  }
  return xs;
}
