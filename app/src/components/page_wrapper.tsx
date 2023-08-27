import {type PropsWithChildren} from "react";

import {cn} from "@/app/lib/css";

type Props = {
  className?: string;
  fluid?: boolean;
};

export function PageWrapper({
  children,
  className,
  fluid = false,
}: PropsWithChildren<Props>) {
  return (
    <div
      className={cn(
        "mx-auto flex w-full max-w-6xl flex-1 flex-col",
        className,
        fluid && "max-w-none"
      )}
    >
      {children}
    </div>
  );
}
