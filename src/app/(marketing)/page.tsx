import {PageWrapper} from "@/components/page_wrapper";

const PopularTopics = Object.freeze([
  {
    title: "Algorithms",
    description:
      "Different algorithms are used to solve different problems. From sorting to searching, algorithms are used to solve problems with data. ",
    icon: null,
    path: "/algorithms",
  },
  {
    title: "Data Structures",
    description:
      "Different data structures are used to store different types of data. From linked lists to stacks, data structures are used to store data.",
    icon: null,
    path: "/data-structures",
  },
  {
    title: "Sorting",
    description:
      "Sorting is the process of arranging a list of items in a particular order. Sorting algorithms are used to arrange a list of items in a particular order.",
    icon: null,
    path: "/algorithms/sorting",
  },
  {
    title: "Searching",
    description:
      "Searching is the process of finding a particular item in a list of items. Searching algorithms are used to find a particular item in a list of items. They are used in every part of life, from making a cup of tea to landing a plane.",
    icon: null,
    path: "/algorithms/searching",
  },
  {
    title: "Linked Lists",
    description:
      "Linked lists are a linear data structure that stores data in a non-contiguous manner. They are used to implement other data structures such as queues and stacks. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: null,
    path: "/data-structures/linked-lists",
  },
  {
    title: "Stacks",
    description:
      "A stack is a linear data structure that follows the Last In First Out (LIFO) principle. It is a collection of elements that are added and removed in a specific order. Stacks are used to implement other data structures such as queues and linked lists. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: null,
    path: "/data-structures/stacks",
  },
]);

export default function Home() {
  return (
    <PageWrapper fluid className="border">
      <div className="flex flex-1 items-center bg-main-hero dark:bg-main-hero-white">
        <div className="flex flex-col items-center justify-center rounded-sm bg-gray-950/70 p-2">
          <h1 className="text-4xl font-extrabold tracking-tighter text-gray-200 drop-shadow-2xl dark:text-gray-100 sm:text-5xl lg:text-6xl">
            <span className="block">Learn Algorithms and Data Structures</span>
            <span className=" block">
              with code examples and visualization{" "}
            </span>
          </h1>
          <p className="mt-6 max-w-3xl text-xl text-gray-300 drop-shadow-xl dark:text-gray-400">
            Algorithms and data structures are the building blocks of computer
            science. They are used to solve problems with data. They are a
            fundamental part of computer science, and understanding them is
            essential to becoming a good programmer.
          </p>
        </div>
      </div>
      <ul className="grid flex-1 grid-cols-1 gap-5 border p-2 sm:grid-cols-2 lg:grid-cols-3">
        {PopularTopics.map((topic) => (
          <li
            key={topic.title}
            className="h-40 rounded-md border border-gray-800 p-1 shadow-md"
          >
            <div className="h-full">
              <strong className="font-bold">{topic.title}</strong>
              {/* TODO tooltip */}
              <p className="text-sm text-gray-700">
                {topic.description.length > 150
                  ? `${topic.description.substring(0, 150)}... `
                  : topic.description}
              </p>
            </div>
          </li>
        ))}
      </ul>
    </PageWrapper>
  );
}
