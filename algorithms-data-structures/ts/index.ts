import {Node, merge, printList} from "./algorithms/linked-list/merge";

let t1 = new Node(1);
t1.next = new Node(3);
t1.next.next = new Node(5);

let t2 = new Node(2);
t2.next = new Node(4);
t2.next.next = new Node(6);

let merged = merge(t1, t2);
console.log("🚀 ~ file: merge.test.ts:14 ~ test ~ merged:", printList(merged));
