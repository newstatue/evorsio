import { Check, Copy } from "lucide-react"

import {
    InputGroup,
    InputGroupAddon,
    InputGroupButton,
    InputGroupInput,
} from "@/components/ui/input-group"
import {useCopyToClipboard} from "@/hooks/use-copy-to-clipboard.ts";

type CopyInputProps = {
    id: string
    value?: string
    label?: string
    readOnly?: boolean
    onChange?: (value: string) => void
}

export function CopyInput({
                              id,
                              value = "",
                              label = "内容",
                              readOnly = false,
                              onChange
                          }: CopyInputProps) {
    const { copyToClipboard, copiedValue } = useCopyToClipboard()

    return (
        <InputGroup>
            <InputGroupInput
                id={id}
                autoComplete="off"
                value={value}
                readOnly={readOnly}
                onChange={(e) => onChange?.(e.target.value)}
            />

            <InputGroupAddon align="inline-end">
                <InputGroupButton
                    aria-label={`复制${label}`}
                    title={`复制${label}`}
                    size="icon-xs"
                    onClick={() => copyToClipboard(value)}
                >
                    {copiedValue === value ? <Check /> : <Copy />}
                </InputGroupButton>
            </InputGroupAddon>
        </InputGroup>
    )
}