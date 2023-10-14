// contentlayer.config.ts
import { defineDocumentType, makeSource } from "contentlayer/source-files";
import rehypePrettyCode from "rehype-pretty-code";
import rehypeSlug from "rehype-slug";
var Post = defineDocumentType(() => ({
  name: "Post",
  filePathPattern: `**/*.mdx`,
  contentType: "mdx",
  fields: {
    title: { type: "string", required: true },
    description: { type: "string", required: true },
    date: { type: "date", required: true },
    updated: { type: "date", required: false },
    problem: { type: "string", required: true },
    tags: { type: "list", of: { type: "string" }, required: true }
  },
  computedFields: {
    url: {
      type: "string",
      resolve: (post) => `/posts/${post._raw.flattenedPath}`
    },
    slug: {
      type: "string",
      resolve: (post) => post._raw.flattenedPath
    }
  }
}));
var contentlayer_config_default = makeSource({
  contentDirPath: "posts",
  documentTypes: [Post],
  mdx: {
    rehypePlugins: [
      rehypeSlug,
      [
        rehypePrettyCode,
        {
          // link to different themes https://github.com/shikijs/shiki/blob/main/docs/themes.md
          // theme: {
          //   dark: "github-dark-dimmed",
          //   light: "github-light",
          // },
          theme: "github-dark-dimmed",
          onVisitLine(node) {
            if (node.children.length === 0) {
              node.children = [{ type: "text", value: " " }];
            }
          },
          onVisitHighlightedLine(node) {
            node.properties.className.push("line--highlighted");
          },
          onVisitHighlightedWord(node) {
            node.properties.className = ["word--highlighted"];
          }
          // onVisitHighlightedLine(node: any) {
          //   node.properties.className.push("line--highlighted");
          // },
          // onVisitLine(node: Record<string, any>) {
          // 	// Prevent lines from collapsing in `display: grid` mode, and allow empty
          // 	// lines to be copy/pasted
          // 	if (node.children.length === 0) {
          // 		node.children = [{type: "text", value: " "}];
          // 	}
          // },
          // onVisitHighlightedLine(node: Record<string, any>) {
          // 	node.properties.className.push("line--highlighted");
          // },
          // onVisitHighlightedWord(node: Record<string, any>) {
          // 	node.properties.className = ["word--highlighted"];
          // },
        }
      ]
    ]
  }
});
export {
  Post,
  contentlayer_config_default as default
};
//# sourceMappingURL=compiled-contentlayer-config-M4R5N7VJ.mjs.map
