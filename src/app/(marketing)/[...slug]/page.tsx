import {readFile} from "node:fs/promises";
import {join} from "node:path";

import {notFound} from "next/navigation";

import {buildAbsolutePath} from "@/app/lib/tree";
import {TreeData} from "@/data/tree_data";

export async function generateStaticParams() {
  let paths = TreeData.map((node) => {
    let path = buildAbsolutePath(TreeData, node.id);
    return {
      params: {
        slug: path.split("/"),
      },
    };
  });
  return paths;
}

async function getPost(slugs: string[]) {
  // let post = await readFile(join("posts", slugs.join("")));

  let slug = slugs.at(-1);
  if (!slug) {
    return null;
  }
  let absolutPath = process.cwd();
  let post = await readFile(join(absolutPath, "posts", slug + ".mdx"), "utf8");
  return post;
}

async function Page({params}: {params: {slug: string[]}}) {
  let post = await getPost(params.slug);
  if (!post) {
    notFound();
  }
  // We can get the correct MDX post from the last index when accessing the list
  // let leafNodes = TreeData.filter((node) => {
  //   return !TreeData.some((node2) => node2.parentId === node.id);
  // });
  return (
    <div>
      <h1>{params.slug.at(-1)}</h1>
    </div>
  );
}

export default Page;
