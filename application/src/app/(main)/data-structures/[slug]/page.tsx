import {getMdxContents} from "@/app/db/content";

export default function DataStructuresSlugPage({
	params,
}: {
	params: {
		slug: string;
	};
}) {
	let x = getMdxContents();

	console.log("x", x);

	return (
		<div>
			<p>asdas</p>
			{params.slug}
		</div>
	);
}
