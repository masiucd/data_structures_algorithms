"use client";
import * as RadixTooltip from "@radix-ui/react-tooltip";
import {TooltipProviderProps} from "@radix-ui/react-tooltip";
import {PropsWithChildren} from "react";

import {cn} from "@/app/lib/css";

type Props = {
  text: string;
  className?: string;
  disabled?: boolean;
} & TooltipProviderProps;
export function Tooltip({
  children,
  text,
  className,
  disabled,
  delayDuration = 200,
  disableHoverableContent,
}: PropsWithChildren<Props>) {
  if (disabled) return <>{children}</>;
  return (
    <RadixTooltip.Provider
      delayDuration={delayDuration}
      disableHoverableContent={disableHoverableContent}
    >
      <RadixTooltip.Root>
        <RadixTooltip.Trigger asChild>{children}</RadixTooltip.Trigger>
        <RadixTooltip.Portal>
          <RadixTooltip.Content
            className={cn(
              "bg-gray-900 leading-tight text-gray-50 max-w-xs p-1 rounded-md border border-blue-500 shadow-lg",
              className
            )}
          >
            {text}
          </RadixTooltip.Content>
        </RadixTooltip.Portal>
      </RadixTooltip.Root>
    </RadixTooltip.Provider>
  );
}
