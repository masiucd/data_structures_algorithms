import Link from "next/link";

export default function NotFound() {
  return (
    <div>
      <h1>
        Oopps We got a<span className="text-red-500"> 404</span>{" "}
      </h1>
      <p>Sorry, we could not find the page you were looking for</p>
      <Link href="/">Return Home</Link>
    </div>
  );
}
