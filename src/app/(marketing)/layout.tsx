import Link from "next/link";
import {type PropsWithChildren} from "react";

import {Icons} from "@/components/icons";
import {TreeDataType, TreeNode} from "@/types/tree";

import {TreeList} from "./components/tree_list";

const TreeData = Object.freeze([
  {id: 1, title: "Algorithms", parentId: null},
  {id: 2, title: "Data Structures", parentId: null},
  {id: 3, title: "Sorting", parentId: 1},
  {id: 4, title: "Searching", parentId: 1},
  {id: 5, title: "Strings", parentId: 1},
  {id: 6, title: "Arrays", parentId: 2},
  {id: 7, title: "Linked Lists", parentId: 2},
  {id: 8, title: "Stacks", parentId: 2},
  {id: 9, title: "Queues", parentId: 2},
  {id: 10, title: "Trees", parentId: 2},
  {id: 11, title: "Graphs", parentId: 2},
  // {id: 12, title: "Hashing", parentId: 2},
  {id: 13, title: "Recursion", parentId: 1},
  // {id: 14, title: "Backtracking", parentId: 1},
  // {id: 15, title: "Dynamic Programming", parentId: 1},
  // {id: 16, title: "Greedy Algorithms", parentId: 1},
  // {id: 17, title: "Divide and Conquer", parentId: 1},
  // {id: 18, title: "Bit Manipulation", parentId: 1},
  // {id: 19, title: "Mathematics", parentId: 1},
  // {id: 20, title: "Puzzles", parentId: 1},
  {id: 21, title: "Bubble sort", parentId: 3},
  {id: 22, title: "Insertion sort", parentId: 3},
  {id: 23, title: "Dynamic Arrays", parentId: 6},
  {id: 24, title: "Static Arrays", parentId: 6},
  {id: 25, title: "Single Linked List", parentId: 7},
  {id: 26, title: "Double Linked List", parentId: 7},
]);

function generateSlug(title: string) {
  return title.toLowerCase().replace(/ /g, "-");
}

function getChildren(treeData: readonly TreeNode[], parentId: number | null) {
  return treeData.filter((item) => item.parentId === parentId);
}

function buildAbsolutePath(
  treeData: readonly TreeNode[],
  id: number | null = null
): string {
  let node = treeData.find((item) => item.id === id);
  if (!node) {
    return "";
  }
  let parentAbsolutePath = buildAbsolutePath(treeData, node.parentId);
  return `${parentAbsolutePath}/${generateSlug(node.title)}`;
}

function getTree(
  treeData: readonly TreeNode[],
  parentId: number | null,
  absolutePath: string = ""
): TreeDataType[] {
  let children = getChildren(treeData, parentId);
  return children.map((child) => {
    return {
      ...child,
      children: getTree(treeData, child.id, absolutePath),
      href: buildAbsolutePath(treeData, child.id),
    };
  });
}

const PAGE_LINKS = Object.freeze([
  {
    title: "Home",
    href: "/",
  },
  {
    title: "About",
    href: "/about",
  },
  {
    title: "FAQ",
    href: "/faq",
  },
  {
    title: "Contact",
    href: "/contact",
  },
]);

const SOCIAL_LINKS = Object.freeze([
  {
    title: "Twitter",
    href: "https://twitter.com/masiu_cd",
    icon: Icons.TwitterLogo,
  },
  {
    title: "GitHub",
    href: "https://github.com/masiucd",
    icon: Icons.GitHubLogo,
  },
  {
    title: "Discord",
    href: "https://discord.com/users/masiu1916",
    icon: Icons.Discord,
  },

  {
    title: "Instagram",
    href: "https://www.instagram.com/masiu_cd",
    icon: Icons.Instagram,
  },
]);

export default function Layout({children}: PropsWithChildren) {
  let tree = getTree(TreeData, null);

  return (
    <>
      <main className="flex min-h-screen flex-col ">
        <div className="grid flex-1 grid-cols-1 sm:grid-cols-12">
          <aside className="flex flex-col sm:col-span-2 md:col-span-3">
            <ul className="flex max-h-[38rem] flex-col gap-1 overflow-auto">
              <TreeList tree={tree} />
            </ul>
            <div className="mt-auto flex flex-col gap-1 px-1 pb-1">
              <ul className="flex gap-2">
                {PAGE_LINKS.map((link) => (
                  <li key={link.title}>
                    <Link
                      href={link.href}
                      className="text-sm font-semibold text-blue-950 transition-opacity duration-150 hover:opacity-50"
                    >
                      <span>{link.title}</span>
                    </Link>
                  </li>
                ))}
              </ul>
              <ul className="flex gap-2">
                {SOCIAL_LINKS.map((link) => (
                  <li
                    key={link.title}
                    className="text-sm font-semibold text-blue-950 transition-opacity duration-150 hover:opacity-50"
                  >
                    <a href={link.href}>
                      <link.icon fontSize={30} />
                    </a>
                  </li>
                ))}
              </ul>
            </div>
          </aside>
          <section className="flex flex-col  sm:col-span-10 md:col-span-9">
            {children}
          </section>
        </div>
      </main>
    </>
  );
}
