import {Folder, Link} from "lucide-react";

export let Icons = {
	Folder: Folder,
	Link: Link,
};

type Icon = keyof typeof Icons;

export function getIcon(type: Icon) {
	return Icons[type] || Folder;
}
