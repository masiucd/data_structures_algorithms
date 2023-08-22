import {type PropsWithChildren} from "react";

// const Tree = Object.freeze( [
//   {id: 1, title: "Algorithms", parentId: null},
//   {id: 2, title: "Data Structures", parentId: null},
//   {id: 3, title: "Sorting", parentId: 1},
//   {id: 4, title: "Searching", parentId: 1},
//   {id: 5, title: "Strings", parentId: 1},
//   {id: 6, title: "Arrays", parentId: 2},
//   {id: 7, title: "Linked Lists", parentId: 2},
//   {id: 8, title: "Stacks", parentId: 2},
//   {id: 9, title: "Queues", parentId: 2},
//   {id: 10, title: "Trees", parentId: 2},
//   {id: 11, title: "Graphs", parentId: 2},
//   {id: 12, title: "Hashing", parentId: 2},
//   {id: 13, title: "Recursion", parentId: 1},
//   {id: 14, title: "Backtracking", parentId: 1},
//   {id: 15, title: "Dynamic Programming", parentId: 1},
//   {id: 16, title: "Greedy Algorithms", parentId: 1},
//   {id: 17, title: "Divide and Conquer", parentId: 1},
//   {id: 18, title: "Bit Manipulation", parentId: 1},
//   {id: 19, title: "Mathematics", parentId: 1},
//   {id: 20, title: "Puzzles", parentId: 1},
// ]);

export default function Layout({children}: PropsWithChildren) {
  return (
    <>
      <main className="flex min-h-screen flex-col ">
        <div className="grid flex-1 grid-cols-1  sm:grid-cols-12">
          <aside className=" sm:col-span-3">
            <ul>
              <li>Home</li>
              <li>About</li>
              <li>FAQ</li>
              <li>Contact</li>
            </ul>
          </aside>
          <section className="flex flex-col  sm:col-span-9">{children}</section>
        </div>
      </main>
    </>
  );
}
