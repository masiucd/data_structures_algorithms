import "./globals.css";
import "@radix-ui/themes/styles.css";

import {Theme} from "@radix-ui/themes";
import type {Metadata} from "next";
import {Inter} from "next/font/google";
import {type ReactNode} from "react";

const inter = Inter({subsets: ["latin"]});

export const metadata: Metadata = {
	title: "Data Structures and Algorithms",
	description: "Data Structures and Algorithms learning resources",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: ReactNode;
}>) {
	return (
		<html lang="en">
			<body className={inter.className}>
				<Theme
					accentColor="gold"
					grayColor="sage"
					radius="medium"
					scaling="90%"
				>
					{children}
				</Theme>
			</body>
		</html>
	);
}
