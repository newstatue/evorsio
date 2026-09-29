import type { CSSProperties } from "react"
import { createRootRoute, Outlet } from "@tanstack/react-router"
import { TitleBar } from "@/components/layout/title-bar.tsx"
import {platform} from "@/lib/platform.ts";
import {SidebarProvider} from "@/components/ui/sidebar.tsx";

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {

  const style = {
    "--titlebar-height": platform.isWindowsLike() ? "2rem" : "0rem",
    "--header-height": platform.isMacOS() ? "2.75rem":"0rem",
    "--header-macos-pd": "5.5rem",
    ...(platform.isDesktop() ? { "--sidebar": "transparent" } : {}),
  } as CSSProperties

  return (
      <SidebarProvider style={style} className={"flex h-svh flex-col overflow-hidden"}>
        { platform.isWindowsLike() && <TitleBar />}
        <div className="flex min-h-0 min-w-0 w-full flex-1">
          <Outlet />
        </div>
        {/*<TanStackRouterDevtools position="bottom-right" />*/}
      </SidebarProvider>
  )
}
