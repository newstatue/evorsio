import { Button } from "@/components/ui/button.tsx"
import { Square, X, Minus } from "lucide-react"
import { cn } from "cn"
import {
  Menubar,
  MenubarContent,
  MenubarGroup,
  MenubarItem,
  MenubarMenu,
  MenubarSeparator,
  MenubarShortcut,
  MenubarTrigger,
} from "@/components/ui/menubar"
import { Kbd } from "@/components/ui/kbd.tsx"
import { Window } from "@wailsio/runtime"
import type { CSSProperties } from "react"
export function TitleBar() {
  const windowButtonClass =
    "h-full w-12 rounded-none [--wails-draggable:no-drag] active:not-aria-[haspopup]:translate-y-0"

  return (
    <header className="flex h-(--titlebar-height) shrink-0 items-center bg-sidebar [--wails-draggable:drag]">
      <div className="flex items-center">
        <Menubar className="rounded-sm border-0 font-normal [--wails-draggable:no-drag]">
          <MenubarMenu>
            <MenubarTrigger>File</MenubarTrigger>
            <MenubarContent>
              <MenubarGroup>
                <MenubarItem>
                  New Tab
                  <MenubarShortcut>
                    <Kbd className="rounded-sm">⌘K</Kbd>
                  </MenubarShortcut>
                </MenubarItem>
                <MenubarItem>New Window</MenubarItem>
              </MenubarGroup>
              <MenubarSeparator />
              <MenubarGroup>
                <MenubarItem>Share</MenubarItem>
                <MenubarItem>Print</MenubarItem>
              </MenubarGroup>
            </MenubarContent>
          </MenubarMenu>
          <MenubarMenu>
            <MenubarTrigger>File</MenubarTrigger>
            <MenubarContent>
              <MenubarGroup>
                <MenubarItem>
                  New Tab
                  <MenubarShortcut>
                    <Kbd className="rounded-sm">⌘K</Kbd>
                  </MenubarShortcut>
                </MenubarItem>
                <MenubarItem>New Window</MenubarItem>
              </MenubarGroup>
              <MenubarSeparator />
              <MenubarGroup>
                <MenubarItem>Share</MenubarItem>
                <MenubarItem>Print</MenubarItem>
              </MenubarGroup>
            </MenubarContent>
          </MenubarMenu>
        </Menubar>
      </div>

      <div className="flex-1" />

      <div className="flex h-full items-center">
        <Button
          size="icon"
          variant="ghost"
          className={windowButtonClass}
          aria-label="最小化"
          onClick={() => Window.Minimise()}
          style={{ "--wails-non-client-region": "minimize" } as CSSProperties}
        >
          <Minus className="size-4" />
        </Button>

        <Button
          size="icon"
          variant="ghost"
          className={windowButtonClass}
          aria-label="最大化"
          onClick={() => Window.ToggleMaximise()}
          style={{ "--wails-non-client-region": "maximize" } as CSSProperties}
        >
          <Square className="size-3" />
        </Button>

        <Button
          size="icon"
          variant="ghost"
          className={cn(windowButtonClass, "hover:bg-red-600 hover:text-white")}
          aria-label="关闭"
          onClick={() => Window.Close()}
          style={{ "--wails-non-client-region": "close" } as CSSProperties}
        >
          <X className="size-4" />
        </Button>
      </div>
    </header>
  )
}
