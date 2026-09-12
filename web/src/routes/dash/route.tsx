import type { CSSProperties } from "react"
import { createFileRoute, Outlet } from "@tanstack/react-router"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import { DashHeader } from "@/components/layout/dash/dash-header.tsx"
import { DashSidebar } from "@/components/layout/dash/dash-sidebar.tsx"
import {System} from "@wailsio/runtime";

export const Route = createFileRoute("/dash")({
    component: RouteComponent,
})

function RouteComponent() {
    const isWindows = System.IsWindows()

    const sidebarStyle = {
        "--header-height": "3.5rem",
        "--header-macos-pd": "5.5rem",
        ...(!isWindows ? { "--sidebar": "transparent" } : {}),
    } as CSSProperties

    return (
        <SidebarProvider style={sidebarStyle} className={`h-svh overflow-hidden ${isWindows ? "windows-sidebar" : ""}`}>
            <DashHeader />
            <DashSidebar />
            <SidebarInset className="min-h-0 flex-1 bg-background pt-(--header-height)">
                <main className="flex flex-1 flex-col overflow-y-auto p-2">
                    <div className="mx-auto w-5/6 max-w-2xl">
                        <Outlet />
                    </div>
                </main>
            </SidebarInset>
        </SidebarProvider>
    )
}