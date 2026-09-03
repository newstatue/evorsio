import {createFileRoute, Link, Outlet} from '@tanstack/react-router'
import {
    Sidebar,
    SidebarContent, SidebarFooter,
    SidebarGroup, SidebarGroupContent, SidebarGroupLabel,
    SidebarHeader, SidebarMenu, SidebarMenuButton, SidebarMenuItem,
    SidebarProvider
} from "@/components/ui/sidebar";
import {HardDrive, Lock, User} from "lucide-react";

export const Route = createFileRoute('/dash')({
    component: RouteComponent,
})

function RouteComponent() {
    // @ts-ignore
    const SIDEBAR_KEYBOARD_SHORTCUT = "b"
    return (
        <SidebarProvider>
            <Sidebar variant="floating" collapsible="icon">
                <SidebarHeader />
                <SidebarContent>
                    <SidebarGroup>
                        <SidebarGroupLabel>应用</SidebarGroupLabel>
                        <SidebarGroupContent></SidebarGroupContent>
                        <SidebarMenu>
                            <SidebarMenuItem>
                                <SidebarMenuButton render={<Link to="/dash/drive"  activeProps={{
                                    "data-active": true,
                                }}/>}>
                                    <HardDrive/> <span>存储</span>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                            <SidebarMenuItem>
                                <SidebarMenuButton render={<Link to="/dash/vault"  activeProps={{
                                    "data-active": true,
                                }}/>}>
                                    <Lock/> <span>密码</span>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                        </SidebarMenu>
                    </SidebarGroup>
                </SidebarContent>
                <SidebarFooter>
                    <SidebarMenu>
                        <SidebarMenuItem>
                            <SidebarMenuButton>
                                <User /> <span>本地用户</span>
                            </SidebarMenuButton>
                        </SidebarMenuItem>
                    </SidebarMenu>
                </SidebarFooter>
            </Sidebar>
            <main className="flex-1 p-6">
                <div className="mx-auto w-full max-w-7xl">
                    <Outlet />
                </div>
            </main>
        </SidebarProvider>
    )
}
