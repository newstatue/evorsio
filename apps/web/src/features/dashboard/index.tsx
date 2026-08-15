import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { Outlet } from "@tanstack/react-router"
import { AppSidebar } from "@/features/dashboard/app-sidebar.tsx"

export default function DashboardLayout() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <main>
        <SidebarTrigger />
        <Outlet />
      </main>
    </SidebarProvider>
  )
}
