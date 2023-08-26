import {allPosts, Post} from "contentlayer/generated";
import {format, parseISO} from "date-fns";
import {notFound} from "next/navigation";
import {useMDXComponent} from "next-contentlayer/hooks";
import {ReactNode} from "react";

import {buildAbsolutePath} from "@/app/lib/tree";
import {Icons} from "@/components/icons";
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

function Box({children}: {children: ReactNode}) {
  return (
    <div className="mb-3 flex flex-col gap-1 rounded-md p-1 shadow-md dark:border-gray-700">
      {children}
    </div>
  );
}

function Mdx({code}: {code: string}) {
  let MDXContent = useMDXComponent(code);

  return (
    <article className="prose prose-neutral mx-auto py-1  dark:prose-invert prose-h2:mb-1 prose-h2:mt-0">
      <MDXContent
        components={{
          // pre: (props) => (
          //   // @ts-ignore
          //   <Code
          //     {...props}
          //     theme={{
          //       dark: "github-dark",
          //       light: "github-light",
          //     }}
          //     // lineNumbers
          //   />
          // ),
          Box,
        }}
      />
    </article>
  );
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
          <h1 className="">{post.title}</h1>
          <p className="">{post.description}</p>
        </div>
        <div className="flex flex-col justify-between ">
          <time className="">{format(parseISO(date), "LLL do, yyyy")}</time>
          <ul className="m-0 flex list-none gap-3 p-0">
            {post.tags.map((tag) => (
              <li
                className="m-0 flex items-center gap-1 p-0 capitalize"
                key={tag}
              >
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
