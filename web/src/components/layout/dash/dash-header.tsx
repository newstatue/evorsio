import { SidebarTrigger, useSidebar } from "@/components/ui/sidebar.tsx"
import { Link, useMatches } from "@tanstack/react-router"
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb.tsx"
import { Separator } from "@/components/ui/separator.tsx"
import { Fragment } from "react"
import { cn } from "cn"
import { System } from "@wailsio/runtime"

export function DashHeader() {
  const { open, isMobile } = useSidebar()

  const matches = useMatches()
  const breadcrumbMatches = matches.filter(
    (match) => match.staticData?.breadcrumb
  )
  return (
    <header className="fixed inset-x-0 top-0 z-10 flex h-(--header-height) items-center p-2">
      <div
        className={cn(
          "shrink-0 bg-transparent transition-[width] duration-200 ease-linear",
          open && !isMobile
            ? "w-(--sidebar-width)"
            : System.IsMac()
              ? "w-(--header-macos-pd)"
              : "w-0"
        )}
      />
      <div className="flex h-full min-w-0 flex-1 items-center gap-2 bg-background">
        <SidebarTrigger size="icon" />
        <Separator orientation="vertical" />
        <Breadcrumb>
          <BreadcrumbList>
            {breadcrumbMatches.map((match, index) => {
              const isLast = index === breadcrumbMatches.length - 1
              const label = match.staticData!.breadcrumb as string
              return (
                <Fragment key={match.id}>
                  {index > 0 && <BreadcrumbSeparator />}
                  <BreadcrumbItem>
                    {isLast ? (
                      <BreadcrumbPage>{label}</BreadcrumbPage>
                    ) : (
                      <BreadcrumbLink render={<Link to={match.pathname} />}>
                        {label}
                      </BreadcrumbLink>
                    )}
                  </BreadcrumbItem>
                </Fragment>
              )
            })}
          </BreadcrumbList>
        </Breadcrumb>
      </div>
    </header>
  )
}
