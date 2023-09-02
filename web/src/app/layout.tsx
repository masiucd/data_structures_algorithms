import "./globals.css";

import type {Metadata} from "next";
import {JetBrains_Mono} from "next/font/google";
import {type ReactNode} from "react";

import {cn} from "./lib/css";

const monoFont = JetBrains_Mono({
  weight: ["400", "700"],
  variable: "--mono",
  display: "swap",
  subsets: ["latin-ext"],
});

// const local = localFont({
//   src: [
//     {
//       path: "../../public/fonts/sourcecodepro-italic-variablefont_wght-webfont.woff2",
//       weight: "400",
//       style: "italic",
//     },
//     {
//       path: "../../public/fonts/sourcecodepro-variablefont_wght-webfont.woff2",
//       weight: "400",
//       style: "normal",
//     },
//   ],
//   variable: "--mono",
//   display: "swap",
// });

// TODO add more meta tags
export const metadata: Metadata = {
  title: "Just alos and data structures",
  description: "Just alos and data structures, nothing more, nothing less.",
};

export default function RootLayout({children}: {children: ReactNode}) {
  return (
    <html lang="en">
      <body
        className={cn(
          "bg-gray-50 text-gray-950 dark:bg-gray-950 dark:text-gray-50 h-full font-mono font-normal",
          monoFont.variable
        )}
      >
        {children}
      </body>
    </html>
  );
}
