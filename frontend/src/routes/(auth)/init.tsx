import {createFileRoute, useNavigate} from '@tanstack/react-router'
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Field, FieldError, FieldLabel} from "@/components/ui/field.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {useState} from "react";
import {Init} from "@bindings/github.com/newstatue/evorsio/internal/vault/service.ts";
import {toast} from "@/components/ui/toast.tsx";

export const Route = createFileRoute('/(auth)/init')({
    component: RouteComponent,
})

function RouteComponent() {
    const [masterPass, setMasterPass] = useState("")
    const [confirmMasterPass, setConfirmMasterPass] = useState("")
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")

    const navigate = useNavigate()

    const handleSubmit = async () => {
        setError("")
        if (!masterPass) {
            setError("请输入主密码")
            return
        }

        if (masterPass !== confirmMasterPass) {
            setError("两次输入的主密码不一致")
            return
        }

        setLoading(true)
        try {
            await Init(masterPass)

            setMasterPass("")
            setConfirmMasterPass("")

            await navigate({
                to: "/dash",
            })
        }catch(err) {
            toast.add({
                title: "设置失败",
                description: (err as Error).message,
                type: "error",
            })
        }finally {
            setLoading(false)
        }
    }

    return (
        <div className="flex min-h-0 w-full flex-1 items-center justify-center pb-6 px-6">
            <Card className="w-full max-w-sm">
                <CardHeader className="text-center">
                    <CardTitle className="text-xl">设置主密码</CardTitle>
                    <CardDescription>主密码用于解锁 Evorsio，请妥善保管</CardDescription>
                </CardHeader>

                <CardContent>
                    <form
                        className="space-y-4"
                        onSubmit={async (e) => {
                            e.preventDefault()
                            await handleSubmit()
                        }}
                    >
                        <Field>
                            <FieldLabel>主密码</FieldLabel>
                            <Input
                                id="master-pass"
                                type="password"
                                value={masterPass}
                                onChange={(e)=>setMasterPass(e.target.value)}
                                autoFocus={true}
                            />
                        </Field>
                        <Field>
                            <FieldLabel>确认主密码</FieldLabel>
                            <Input
                                id="confirm-master-pass"
                                type="password"
                                value={confirmMasterPass}
                                onChange={(e)=>setConfirmMasterPass(e.target.value)}
                            />
                        </Field>
                        <Field>
                            <Button
                                type="submit"
                                disabled={loading}
                            >
                                初始化
                            </Button>
                            <FieldError>{error}</FieldError>
                        </Field>
                    </form>
                </CardContent>
            </Card>

        </div>
    )
}
