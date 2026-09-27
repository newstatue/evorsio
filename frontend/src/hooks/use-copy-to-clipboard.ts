import { useState } from "react"
import { Clipboard } from "@wailsio/runtime"

export function useCopyToClipboard() {
    const [copiedValue, setCopiedValue] = useState<string | null>(null)

    async function copyToClipboard(text: string) {
        if (!text) return

        await Clipboard.SetText(text)
        setCopiedValue(text)

        setTimeout(() => {
            setCopiedValue(null)
        }, 1500)
    }

    return {
        copyToClipboard,
        copiedValue,
    }
}