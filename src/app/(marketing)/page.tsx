const PopularTopics = Object.freeze([
  {
    title: "Algorithms",
    description:
      "Algorithms are a set of instructions or rules designed to solve a specific problem. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: null,
    path: "/algorithms",
  },
  {
    title: "Data Structures",
    description:
      "Data structures are a way of organizing and storing data so that it can be accessed and modified efficiently. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: null,
    path: "/data-structures",
  },
  {
    title: "Sorting",
    description:
      "Sorting is the process of arranging a list of items in a particular order. Sorting algorithms are used to sort data into a specific order. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: null,
    path: "/algorithms/sorting",
  },
  {
    title: "Searching",
    description:
      "Searching is the process of finding a particular item in a list of items. Searching algorithms are used to find a particular item in a list of items. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: null,
    path: "/algorithms/searching",
  },
  {
    title: "Linked Lists",
    description:
      "A linked list is a linear data structure that consists of a sequence of nodes. Each node contains data and a pointer to the next node in the sequence. Linked lists are used to implement other data structures such as stacks and queues. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
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
    <div>
      <ul className="grid grid-cols-1 gap-5 border p-2 sm:grid-cols-2 lg:grid-cols-3">
        {PopularTopics.map((topic) => (
          <li key={topic.title} className="max-w-sm">
            <h2>{topic.title}</h2>
            <p className="truncate text-sm text-gray-500">
              {topic.description}
            </p>
          </li>
        ))}
      </ul>
    </div>
  );
}
