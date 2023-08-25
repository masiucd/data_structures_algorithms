import Link from "next/link";
import {type PropsWithChildren} from "react";

import {Icons} from "@/components/icons";
import {TreeData} from "@/data/tree_data";
import {TreeDataType, TreeNode} from "@/types/tree";

import {buildAbsolutePath} from "../lib/tree";
import {TreeList} from "./components/tree_list";

function getChildren(treeData: readonly TreeNode[], parentId: number | null) {
  return treeData.filter((item) => item.parentId === parentId);
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
          <aside className="flex border-gray-600/30 shadow-inner sm:col-span-2 sm:flex-col sm:border-r sm:shadow md:col-span-3">
            <ul className="flex max-h-[38rem] flex-col gap-1 overflow-auto">
              <TreeList tree={tree} />
            </ul>
            <div className="ml-auto flex flex-col gap-1 px-1 pb-1 sm:mt-auto  sm:w-full">
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
