import {createFileRoute, Link, Outlet} from '@tanstack/react-router'
import {
    Sidebar,
    SidebarContent, SidebarFooter,
    SidebarGroup, SidebarGroupAction, SidebarGroupContent, SidebarGroupLabel,
    SidebarHeader, SidebarInput, SidebarMenu, SidebarMenuBadge, SidebarMenuButton, SidebarMenuItem,
    SidebarProvider, SidebarRail, SidebarSeparator, SidebarTrigger
} from "@/components/ui/sidebar";
import {HardDrive, Lock, User} from "lucide-react";
import {Field} from "@/components/ui/field.tsx";

export const Route = createFileRoute('/dash')({
    component: RouteComponent,
})

function RouteComponent() {
    // @ts-ignore
    const SIDEBAR_KEYBOARD_SHORTCUT = "b"
    // @ts-ignore
    const SIDEBAR_WIDTH = "16rem"
    // @ts-ignore
    const SIDEBAR_WIDTH_MOBILE = "18rem"
    return (
        <SidebarProvider>
            <Sidebar variant="sidebar" collapsible="offcanvas">
                <SidebarHeader />
                <SidebarContent>
                    <SidebarGroup>
                        <SidebarGroupLabel>应用</SidebarGroupLabel>
                        <SidebarGroupAction></SidebarGroupAction>
                        <SidebarGroupContent>
                            <SidebarInput/>
                        </SidebarGroupContent>

                        <SidebarMenu>
                            <SidebarMenuItem>
                                <SidebarMenuButton render={<Link to="/dash/drive"  activeProps={{
                                    "data-active": true,
                                }}/>}>
                                    <HardDrive/> <span>存储</span>
                                </SidebarMenuButton>
                                <SidebarMenuBadge/>
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
                <SidebarRail/>
            </Sidebar>
            <main className="flex-1 p-6">
                <header>
                    <SidebarTrigger size="icon"/>
                </header>
                <div className="mx-auto w-full max-w-7xl">
                    <Outlet />
                </div>
            </main>
        </SidebarProvider>
    )
}
