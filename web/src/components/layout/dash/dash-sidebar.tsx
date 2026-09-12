import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupAction,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarInput,
    SidebarMenu,
    SidebarMenuBadge,
    SidebarMenuButton,
    SidebarMenuItem,
} from "@/components/ui/sidebar"
import { Link } from "@tanstack/react-router"
import {HardDrive, User, Lock} from "lucide-react"
import { cn } from "cn"
import { System } from "@wailsio/runtime"

export function DashSidebar() {
    const variant = System.IsWindows() ? "inset" : "sidebar"
    return (
        <Sidebar variant={variant} collapsible="offcanvas" className="top-(--titlebar-height) z-20 h-[calc(100svh-var(--titlebar-height))] select-none">
            <div className={cn(System.IsMac() ? "h-(--header-height)" : "h-0")} />
            <SidebarHeader>
               <span>Evorsio</span>
            </SidebarHeader>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel>应用</SidebarGroupLabel>
                    <SidebarGroupAction></SidebarGroupAction>
                    <SidebarGroupContent>
                        <SidebarInput />
                    </SidebarGroupContent>

                    <SidebarMenu>
                        <SidebarMenuItem>
                            <SidebarMenuButton
                                render={
                                    <Link
                                        to="/dash/drive"
                                        activeProps={{
                                            "data-active": true,
                                        }}
                                    />
                                }
                            >
                                <HardDrive /> <span>存储</span>
                            </SidebarMenuButton>
                            <SidebarMenuBadge />
                        </SidebarMenuItem>
                        <SidebarMenuItem>
                            <SidebarMenuButton
                                render={
                                    <Link
                                        to="/dash/vault"
                                        activeProps={{
                                            "data-active": true,
                                        }}
                                    />
                                }
                            >
                                <Lock /> <span>密码</span>
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
            {/*<SidebarRail/>*/}
        </Sidebar>
    )
}
