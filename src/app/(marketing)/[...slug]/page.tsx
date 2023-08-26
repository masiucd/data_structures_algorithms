import {allPosts, Post} from "contentlayer/generated";
import {format, parseISO} from "date-fns";
import {notFound} from "next/navigation";

import {buildAbsolutePath} from "@/app/lib/tree";
import {Icons} from "@/components/icons";
import {TreeData} from "@/data/tree_data";

import {Mdx} from "./components/mdx";

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

export async function generateMetadata({params}: {params: {slug: string[]}}) {
  let slug = params.slug.at(-1);
  if (!slug) {
    return null;
  }
  let post = allPosts.find((post) => post.slug === slug);
  if (!post) {
    return null;
  }
  let {title} = post;
  return {title};
}

async function getPost(slugs: string[]) {
  let slug = slugs.at(-1);
  if (!slug) {
    return null;
  }
  let post = allPosts.find((post) => post.slug === slug);
  return post;
}

function getDate(post: Post) {
  if (!post.updated) {
    return post.date;
  }
  return post.date === post.updated ? post.date : post.updated;
}

async function Page({params}: {params: {slug: string[]}}) {
  let post = await getPost(params.slug);
  if (!post) {
    notFound();
  }

  let date = getDate(post);
  return (
    <section>
      <aside className="mx-auto mb-5 flex w-[650px] justify-between pt-10 ">
        <div className="flex flex-col">
          <h1 className="mb-2">{post.title}</h1>
          <div>
            <p>{post.description}</p>
            <a
              href={post.problem}
              className="flex items-center gap-1 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300"
              target="_blank"
              rel="noopener noreferrer"
            >
              <span>
                <Icons.Link size={20} />
              </span>
              <span>Leet code</span>
            </a>
          </div>
        </div>
        <div className="flex flex-col gap-1">
          <time className="text-gray-500 dark:text-gray-400">
            {format(parseISO(date), "LLL do, yyyy")}
          </time>
          <ul className=" flex list-none gap-1 ">
            {post.tags.map((tag) => (
              <li className=" flex items-center capitalize" key={tag}>
                <span>
                  <Icons.Hash size={20} />
                </span>
                <span>{tag}</span>
              </li>
            ))}
          </ul>
        </div>
      </aside>
      <Mdx code={post.body.code} />
    </section>
  );
}

export default Page;
