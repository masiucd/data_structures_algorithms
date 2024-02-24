import {Heading, Text} from "@radix-ui/themes";
import Link from "next/link";

import {Icons} from "@/components/icons";
import {PageWrapper} from "@/components/page-wrapper";
import {cn} from "@/lib/cn";

export default function Home() {
	return (
		<PageWrapper>
			<section className="mb-5 ">
				<Heading as="h1">Data Structures and Algorithms</Heading>
				<aside className="md:pr-60">
					<Text as="p">
						A collection of learning resources for data structures and
						algorithms. From basic to advanced topics.
					</Text>
					<Text>
						Coding problems, algorithms, and data structures are the most
						important part of any software development development. This is why
						we need to have a good understanding of data structures and
						algorithms.
					</Text>
				</aside>
			</section>
			<Tree treeData={buildTree(data)} level={0} />
		</PageWrapper>
	);
}

function Tree({treeData, level}: {treeData: TreeNode[]; level: number}) {
	return (
		<ul
			className={cn(
				"flex flex-col gap-2",
				level === 0 && "mb-10",
				level === 1 && "pl-4",
				level === 2 && "pl-8",
				level === 3 && "pl-12",
			)}
		>
			{treeData.map((node, index) => (
				<li key={index}>
					{node.children && node.children.length > 0 ? (
						<Heading as="h3" size="3" className={cn(level === 0 && "mb-2")}>
							{node.title}
						</Heading>
					) : (
						<Link
							className="flex cursor-pointer items-center gap-1 hover:text-primary-600 hover:underline"
							href={getHref(node.title)}
						>
							<Icons.Link size={14} /> <span>{node.title}</span>
						</Link>
					)}

					{node.children && <Tree treeData={node.children} level={level + 1} />}
				</li>
			))}
		</ul>
	);
}

function getHref(title: string) {
	if (dataStructures.includes(title)) {
		return `/data-structures/${slugify(title)}`;
	} else if (algorithms.includes(slugify(title))) {
		return `/algorithms/${title.toLowerCase()}`;
	}
	return "/";
}

let data = [
	{
		id: 1,
		title: "Data Structures",
		parentId: null,
	},
	{
		id: 2,
		title: "Algorithms",
		parentId: null,
	},
	{
		id: 3,
		title: "Arrays",
		parentId: 1,
	},
	{
		id: 4,
		title: "Linked Lists",
		parentId: 1,
	},
	{
		id: 5,
		title: "Stacks",
		parentId: 1,
	},
	{
		id: 6,
		title: "Queues",
		parentId: 1,
	},
	{
		id: 7,
		title: "Hash Tables",
		parentId: 1,
	},
	{
		id: 8,
		title: "Trees",
		parentId: 1,
	},
	{
		id: 9,
		title: "Graphs",
		parentId: 1,
	},
	{
		id: 10,
		title: "Sorting",
		parentId: 2,
	},
	{
		id: 11,
		title: "Searching",
		parentId: 2,
	},
	{
		id: 12,
		title: "Dynamic Programming",
		parentId: 2,
	},
	{
		id: 13,
		title: "Greedy Algorithms",
		parentId: 2,
	},
	{
		id: 14,
		title: "Backtracking",
		parentId: 2,
	},
	{
		id: 15,
		title: "Divide and Conquer",
		parentId: 2,
	},
	{
		id: 16,
		title: "Bit Manipulation",
		parentId: 2,
	},
];

const PARENT_ID_FOR_DATA_STRUCTURES = 1;
const PARENT_ID_FOR_ALGORITHMS = 2;

function slugify(text: string) {
	return text
		.toString()
		.toLowerCase()
		.replace(/\s+/g, "-") // Replace spaces with -
		.replace(/[^\w-]+/g, "") // Remove all non-word chars
		.replace(/--+/g, "-") // Replace multiple - with single -
		.replace(/^-+/, "") // Trim - from start of text
		.replace(/-+$/, ""); // Trim - from end of text
}

function getTitlesByParentId(data: TreeNode[], parentId: number | null) {
	return data
		.filter((item) => item.parentId === parentId)
		.map((item) => item.title);
}
let dataStructures = getTitlesByParentId(data, PARENT_ID_FOR_DATA_STRUCTURES);
let algorithms = getTitlesByParentId(data, PARENT_ID_FOR_ALGORITHMS);

type TreeNode = {
	id: number;
	title: string;
	parentId: number | null;
	children?: TreeNode[];
};

function buildTree(
	treeData: TreeNode[],
	parentId: number | null = null,
): TreeNode[] {
	let nodes = treeData.filter((node) => node.parentId === parentId);
	return nodes.map((node) => ({
		...node,
		children: buildTree(treeData, node.id),
	}));
}
