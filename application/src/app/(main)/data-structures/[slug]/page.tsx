import {getMdxContents} from "@/app/db/content";

export async function generateStaticParams() {
	return getMdxContents().map(({slug}) => ({slug}));
}

function getPost(slug: string) {
	return getMdxContents().find((post) => post.slug === slug);
}

export default function DataStructuresSlugPage({
	params,
}: {
	params: {
		slug: string;
	};
}) {
	let post = getPost(params.slug);
	if (!post) {
		return <div>Post not found</div>;
	}
	return (
		<div>
			<p>asdas</p>
			{post.content}
		</div>
	);
}
