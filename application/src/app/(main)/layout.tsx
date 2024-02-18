import {type ReactNode} from "react";

export default function MainLayout({
	children,
}: Readonly<{children: ReactNode}>) {
	return <main>{children}</main>;
}
