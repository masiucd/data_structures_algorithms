import {useMDXComponent} from "next-contentlayer/hooks";
import {type ReactNode} from "react";

function Box({children}: {children: ReactNode}) {
  return (
    <div className="mb-3 flex flex-col gap-1 rounded-md p-1 shadow-md dark:border-gray-700">
      {children}
    </div>
  );
}

export function Mdx({code}: {code: string}) {
  let MDXContent = useMDXComponent(code);
  return (
    <article className="prose prose-neutral mx-auto py-1  dark:prose-invert prose-h2:mb-1 prose-h2:mt-0">
      <MDXContent
        components={{
          // TODO - copy button
          Box,
        }}
      />
    </article>
  );
}
