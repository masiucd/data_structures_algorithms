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

function Page({params}: {params: {slug: string[]}}) {
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
