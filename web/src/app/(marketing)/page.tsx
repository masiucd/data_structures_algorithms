import Link from "next/link";

import {Icons} from "@/components/icons";
import {PageWrapper} from "@/components/page_wrapper";
import {Tooltip} from "@/components/tooltip";

const TopicsExamples = [
  {
    title: "Algorithms",
    description:
      "Different algorithms are used to solve different problems. From sorting to searching, algorithms are used to solve problems with data. ",
    icon: Icons.PieChart,
    path: "/topics/algorithms",
  },
  {
    title: "Data Structures",
    description:
      "Different data structures are used to store different types of data. From linked lists to stacks, data structures are used to store data.",
    icon: Icons.Database,
    path: "topics/data-structures",
  },
  {
    title: "Sorting",
    description:
      "Sorting is the process of arranging a list of items in a particular order. Sorting algorithms are used to arrange a list of items in a particular order.",
    icon: Icons.SortAsc,
    path: "topics/algorithms/sorting",
  },
  {
    title: "Searching",
    description:
      "Searching is the process of finding a particular item in a list of items. Searching algorithms are used to find a particular item in a list of items. They are used in every part of life, from making a cup of tea to landing a plane.",
    icon: Icons.Search,
    path: "/topics/algorithms/searching",
  },
  {
    title: "Linked Lists",
    description:
      "Linked lists are a linear data structure that stores data in a non-contiguous manner. They are used to implement other data structures such as queues and stacks. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: Icons.Queue,
    path: "/topics/data-structures/linked-lists",
  },
  {
    title: "Trees",
    description:
      "Trees are a non-linear data structure that stores data in a hierarchical manner. They are used to implement other data structures such as heaps and graphs. They are used in every part of life, from making a cup of tea to landing a plane. In computer science, algorithms are used to solve problems with data. They are a fundamental part of computer science, and understanding them is essential to becoming a good programmer.",
    icon: Icons.Tree,
    path: "/topics/data-structures/trees",
  },
];

function Top() {
  return (
    <div className="flex flex-1 items-center bg-main-hero dark:bg-main-hero-white">
      <div className="flex flex-col items-center justify-center gap-2 rounded bg-gray-950/70">
        <div className="p-2">
          <h1 className="text-4xl font-extrabold  text-gray-200 drop-shadow-2xl dark:text-gray-100 sm:text-5xl lg:text-6xl">
            <span className="block">
              Learn <span className="text-blue-200">Algorithms</span> and{" "}
              <span className="text-blue-200">Data Structures</span>
            </span>
            <span className="block">with code examples and visualization </span>
          </h1>
        </div>
        <div className="flex w-full flex-col gap-2 rounded bg-blue-950/20 p-2">
          <p className=" max-w-3xl text-xl text-gray-300 drop-shadow-xl dark:text-gray-400">
            Algorithms and data structures are the building blocks of computer
            science. They are used to solve problems with data. They are a
            fundamental part of computer science, and understanding them is
            essential to becoming a good programmer.
          </p>
          <div className="flex w-full gap-5">
            <Link href="/about">
              <span className="text-sm font-semibold text-gray-100 transition-opacity duration-150 hover:opacity-50">
                About
              </span>
            </Link>
            <Link href="/topics">
              <span className="text-sm font-semibold text-gray-100 transition-opacity duration-150 hover:opacity-50">
                Topics
              </span>
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

const DESCRIPTION_LENGTH = 100;
function Grid() {
  return (
    <div className="flex max-h-96 flex-col gap-1 overflow-scroll p-2 md:max-h-full">
      <div className="p-2">
        <h3>Common Topics</h3>
      </div>
      <ul className="grid flex-1 grid-cols-1 gap-5 p-2 sm:grid-cols-2 lg:grid-cols-3 ">
        {TopicsExamples.sort((a, b) => a.title.localeCompare(b.title)).map(
          (topic) => (
            <li
              key={topic.title}
              className="h-32 rounded-md border border-gray-800 bg-gray-100 p-1 shadow-md"
            >
              <div className="flex h-full flex-col">
                <div className="flex items-center gap-1">
                  <topic.icon size={20} />
                  <strong className="font-bold">{topic.title}</strong>
                </div>
                <Tooltip
                  text={topic.description}
                  disabled={topic.description.length < DESCRIPTION_LENGTH}
                >
                  <p className="text-sm text-gray-700">
                    {topic.description.length > DESCRIPTION_LENGTH
                      ? `${topic.description.substring(
                          0,
                          DESCRIPTION_LENGTH
                        )}... `
                      : topic.description}
                  </p>
                </Tooltip>
                <div className="mt-auto">
                  <Link
                    className="text-sm font-semibold text-blue-950 transition-opacity duration-150 hover:opacity-50"
                    href={topic.path}
                  >
                    To {topic.title}
                  </Link>
                </div>
              </div>
            </li>
          )
        )}
      </ul>
    </div>
  );
}

export default function Home() {
  return (
    <PageWrapper fluid className="border">
      <Top />
      <Grid />
    </PageWrapper>
  );
}
