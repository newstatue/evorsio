import type { CSSProperties } from "react"
import { createRootRoute, Outlet } from "@tanstack/react-router"
import { TitleBar } from "@/components/layout/title-bar.tsx"
import { System } from "@wailsio/runtime"

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  const isWindows = System.IsWindows()
  const isDesktop = System.IsDesktop()

  const style = {
    "--titlebar-height": isWindows ? "2rem" : "0px",
    "--header-height": "3.75rem",
    "--header-macos-pd": "5.5rem",
    ...(isDesktop ? { "--sidebar": "transparent" } : {}),
  } as CSSProperties

  return (
    <div style={style} className={"flex h-svh flex-col overflow-hidden"}>
      {isWindows && <TitleBar />}
      <main className="min-h-0 min-w-0 flex-1 overflow-y-auto">
        <Outlet />
      </main>
      {/*<TanStackRouterDevtools position="bottom-right" />*/}
    </div>
  )
}
