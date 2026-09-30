import { createFileRoute, Outlet } from "@tanstack/react-router"
import { SidebarInset } from "@/components/ui/sidebar.tsx"
import { DashHeader } from "@/components/layout/dash/dash-header.tsx"
import { DashSidebar } from "@/components/layout/dash/dash-sidebar.tsx"
import {platform} from "@/lib/platform.ts";

export const Route = createFileRoute("/_authenticated/dash")({
    component: RouteComponent,
})

function RouteComponent() {
    return (
        <>
            {platform.isMacOS() &&  <DashHeader /> }
            <DashSidebar />
            <SidebarInset className="min-h-0 bg-background pt-[calc(var(--header-height)+1rem)] pb-4">
                <main className="flex min-h-0 flex-1 flex-col overflow-y-auto px-6">
                    <div className="mx-auto flex w-full min-w-0 flex-1 flex-col">
                        <Outlet />
                    </div>
                </main>
            </SidebarInset>
        </>
    )
}
