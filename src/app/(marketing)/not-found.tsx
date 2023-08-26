import Link from "next/link";

import {PageWrapper} from "@/components/page_wrapper";

export default function NotFound() {
  return (
    <PageWrapper>
      <section className="flex flex-1 flex-col items-center justify-center space-y-4 border">
        <h1>
          Oops We got a
          <span className="relative text-red-500 after:absolute after:bottom-2 after:left-4 after:h-5 after:w-14 after:rotate-12  after:bg-red-500/30 after:content-[''] ">
            {" "}
            404
          </span>{" "}
        </h1>
        <p>Sorry, we could not find the page you were looking for</p>
        <Link
          className="text-sm font-semibold text-blue-950 transition-opacity duration-150 hover:opacity-50"
          href="/"
        >
          Return Home
        </Link>
      </section>
    </PageWrapper>
  );
}
