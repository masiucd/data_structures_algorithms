import {allPosts, Post} from "contentlayer/generated";
import {format, parseISO} from "date-fns";
import {notFound} from "next/navigation";
import {useMDXComponent} from "next-contentlayer/hooks";

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

// TODO
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

function Mdx({code, post}: {code: string; post: Post}) {
  let MDXContent = useMDXComponent(code);
  let date = getDate(post);
  return (
    <article className="prose prose-neutral mx-auto border border-red-400 py-1 dark:prose-invert">
      <aside className="mb-5 flex justify-between border border-red-400">
        <div className="flex flex-col">
          <h1 className="m-0 p-0">{post.title}</h1>
          <p className="m-0 p-0">{post.description}</p>
        </div>
        <div className="flex flex-col justify-between ">
          <time className="m-0 p-0">
            {format(parseISO(date), "LLL do, yyyy")}
          </time>
          <ul className="m-0 flex list-none gap-3 p-0">
            {post.tags.map((tag) => (
              <li
                className="m-0 flex items-center gap-1 p-0 capitalize"
                key={tag}
              >
                <span>
                  <Icons.RulerHorizontal />
                </span>
                <span>{tag}</span>
              </li>
            ))}
          </ul>
        </div>
      </aside>
      <MDXContent components={{}} />
    </article>
  );
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
    <section>
      <Mdx code={post.body.code} post={post} />
    </section>
  );
}

export default Page;
