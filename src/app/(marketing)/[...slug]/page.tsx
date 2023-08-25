import {allPosts, Post} from "contentlayer/generated";
import {format, parseISO} from "date-fns";
import {notFound} from "next/navigation";
import {useMDXComponent} from "next-contentlayer/hooks";

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

function Mdx({code, post}: {code: string; post: Post}) {
  let MDXContent = useMDXComponent(code);
  return (
    <article className="prose prose-neutral mx-auto border border-red-400 py-1 dark:prose-invert">
      <aside className="mb-5 border border-red-400">
        <h1 className="m-0 p-0">{post.title}</h1>
        <p className="m-0 p-0">{post.description}</p>
        <p className="m-0 p-0">{format(parseISO(post.date), "yyyy/M/do")}</p>
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
      {/* <aside className="mb-5">
        <h1>{post.title}</h1>
        <p>{format(parseISO(post.date), "yyyy/M/do")}</p>
      </aside> */}
      <Mdx code={post.body.code} post={post} />
    </section>
  );
}

export default Page;
