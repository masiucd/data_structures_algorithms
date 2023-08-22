import "./globals.css";

import type {Metadata} from "next";
import {Inter} from "next/font/google";
import {type ReactNode} from "react";

import {cn} from "./lib/css";

const inter = Inter({subsets: ["latin"]});

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
          "bg-gray-50 text-gray-950 dark:bg-gray-950 dark:text-gray-50 h-full",
          inter.className
        )}
      >
        {children}
      </body>
    </html>
  );
}
