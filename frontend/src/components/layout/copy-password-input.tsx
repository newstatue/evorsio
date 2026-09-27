import { Check, Copy, Eye, EyeOff, RefreshCw } from "lucide-react"
import { useState } from "react"

import {
    InputGroup,
    InputGroupAddon,
    InputGroupButton,
    InputGroupInput,
} from "@/components/ui/input-group"
import { useCopyToClipboard } from "@/hooks/use-copy-to-clipboard"
import { GeneratePassword } from "@bindings/github.com/newstatue/evorsio/internal/vault/service"

type PasswordInputProps = {
    id: string
    value?: string
    label?: string
    readOnly?: boolean
    onChange?: (value: string) => void
}

export function PasswordInput({
                                  id,
                                  value = "",
                                  label = "密码",
                                  readOnly = false,
                                  onChange,
                              }: PasswordInputProps) {
    const [showPassword, setShowPassword] = useState(false)
    const { copyToClipboard, copiedValue } = useCopyToClipboard()

    const generate = async () => {
        const result = await GeneratePassword({
            Length: 16,
            Lowercase: true,
            Numbers: true,
            Symbols: true,
            Uppercase: true,
        })

        onChange?.(result)
    }

    return (
        <InputGroup>
            <InputGroupInput
                id={id}
                type={showPassword ? "text" : "password"}
                autoComplete="off"
                value={value}
                readOnly={readOnly}
                onChange={(e) => onChange?.(e.target.value)}
                className="min-w-0"
            />

            <InputGroupAddon
                align="inline-end"
                className="shrink-0"
            >
                <InputGroupButton
                    aria-label="生成密码"
                    title="生成密码"
                    size="icon-xs"
                    onClick={generate}
                >
                    <RefreshCw />
                </InputGroupButton>

                <InputGroupButton
                    aria-label={showPassword ? "隐藏密码" : "显示密码"}
                    title={showPassword ? "隐藏密码" : "显示密码"}
                    size="icon-xs"
                    onClick={() => setShowPassword((prev) => !prev)}
                >
                    {showPassword ? <EyeOff /> : <Eye />}
                </InputGroupButton>

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