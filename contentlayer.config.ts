import {defineDocumentType, makeSource} from "contentlayer/source-files";
import rehypeAutolinkHeadings from "rehype-autolink-headings";
import rehypePrettyCode from "rehype-pretty-code";
import rehypeSlug from "rehype-slug";
import remarkGfm from "remark-gfm";

export const Post = defineDocumentType(() => ({
  name: "Post",
  filePathPattern: `**/*.mdx`,
  contentType: "mdx",
  fields: {
    title: {type: "string", required: true},
    description: {type: "string", required: true},
    date: {type: "date", required: true},
    updated: {type: "date", required: false},
    problem: {type: "string", required: true},
    tags: {type: "list", of: {type: "string"}, required: true},
  },
  computedFields: {
    url: {
      type: "string",
      resolve: (post) => `/posts/${post._raw.flattenedPath}`,
    },
    slug: {
      type: "string",
      resolve: (post) => post._raw.flattenedPath,
    },
  },
}));

export default makeSource({
  contentDirPath: "posts",
  documentTypes: [Post],
  mdx: {
    remarkPlugins: [remarkGfm],

    rehypePlugins: [
      rehypeSlug,
      [
        rehypePrettyCode,
        {
          // link to different themes https://github.com/shikijs/shiki/blob/main/docs/themes.md
          theme: {
            dark: "github-dark-dimmed",
            light: "github-light",
          },
          // theme: "github-dark-dimmed",
          onVisitLine(node: Record<string, any>) {
            // Prevent lines from collapsing in `display: grid` mode, and allow empty
            // lines to be copy/pasted
            if (node.children.length === 0) {
              node.children = [{type: "text", value: " "}];
            }
          },
          onVisitHighlightedLine(node: Record<string, any>) {
            node.properties.className.push("line--highlighted");
          },
          onVisitHighlightedWord(node: Record<string, any>) {
            node.properties.className = ["word--highlighted"];
          },
        },
      ],

      [
        rehypeAutolinkHeadings,
        {
          properties: {
            className: ["anchor"],
          },
        },
      ],
    ],
  },
});
