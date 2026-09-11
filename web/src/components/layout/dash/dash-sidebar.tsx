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
  SidebarRail,
} from "@/components/ui/sidebar.tsx"
import { Link } from "@tanstack/react-router"
import { HardDrive, User, Lock } from "lucide-react"
import { Button } from "@/components/ui/button.tsx"

export function DashSidebar() {
  return (
    <Sidebar variant="sidebar" collapsible="offcanvas" className="z-20">
      <div className="h-(--header-height)"/>
      <SidebarHeader>
        <Button/>
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
      <SidebarRail />
    </Sidebar>
  )
}
