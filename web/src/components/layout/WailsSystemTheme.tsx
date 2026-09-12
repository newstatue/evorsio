import { useEffect } from "react"
import { Events, System } from "@wailsio/runtime"
import { useTheme } from "@/components/theme-provider"

export function WailsSystemTheme() {
    const { theme } = useTheme()

    useEffect(() => {
        if (theme !== "system") {
            return
        }

        const syncTheme = async () => {
            const isDark = await System.IsDarkMode()
            const root = document.documentElement

            root.classList.toggle("dark", isDark)
            root.classList.toggle("light", !isDark)
        }

        void syncTheme()

        return Events.On("windows:SystemThemeChanged", () => {
            void syncTheme()
        })
    }, [theme])

    return null
}