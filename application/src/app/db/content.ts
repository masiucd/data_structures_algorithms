import fs from "node:fs";
import path from "node:path";

type MetaData = {
	title: string;
	published: string;
	summery: string;
};

// export function getMdxContent(slug: string) {
// 	return "";
// }

export function getMdxContents() {
	let x = path.join(process.cwd(), "src/content");
	fs.readdirSync(x).forEach((file) => {
		console.log(file);
	});
	return x;
}
