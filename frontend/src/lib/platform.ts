import { System } from "@wailsio/runtime"

export const platform = {
    isMacOS: () => System.IsMac(),
    isWindows: () => System.IsWindows(),
    isLinux: () => System.IsLinux(),
    isDesktop: () => System.IsDesktop(),

    isWindowsLike: () =>
        System.IsWindows() || System.IsLinux(),
}