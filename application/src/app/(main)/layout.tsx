import {type ReactNode} from "react";

export default function MainLayout({
	children,
}: Readonly<{children: ReactNode}>) {
	return (
		<>
			<header className="h-[5rem]">
				<div className="app-width mx-auto">
					<strong>Data Structures and Algorithms</strong>
				</div>
			</header>
			<main className="flex min-h-[calc(100dvh-10rem)] flex-col">
				{children}
			</main>
			<footer className="h-[5rem]">
				<div className="app-width mx-auto flex items-center gap-2">
					<strong>
						© {new Date().getFullYear()} Data Structures and Algorithms
					</strong>
					<p>
						Built with
						<a
							target="_blank"
							rel="noopener noreferrer"
							href="https://nextjs.org"
							className="mx-1"
						>
							<span className="text-primary-600 hover:underline focus:underline active:underline">
								Next.js
							</span>
						</a>
					</p>
				</div>
			</footer>
		</>
	);
}
