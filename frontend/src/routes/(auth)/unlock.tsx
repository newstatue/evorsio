import {createFileRoute, useNavigate} from '@tanstack/react-router'
import {useState} from "react";
import {Unlock} from "@bindings/github.com/newstatue/evorsio/internal/vault/service.ts";
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Field, FieldError, FieldLabel} from "@/components/ui/field.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";

export const Route = createFileRoute('/(auth)/unlock')({
  component: RouteComponent,
})

function RouteComponent() {
  const [masterPass, setMasterPass] = useState("")
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")

  const navigate = useNavigate()

  const handleSubmit = async () => {
    setError("")
    if (!masterPass) {
      setError("请输入主密码")
      return
    }

    setLoading(true)
    try {
      await Unlock(masterPass)

      setMasterPass("")

      await navigate({
        to: "/dash",
      })
    }catch {
      setError("解锁错误")
    }finally {
      setLoading(false)
    }
  }

  return (
      <div className="flex min-h-0 w-full flex-1 items-center justify-center pb-6 px-6">
        <Card className="w-full max-w-sm">
          <CardHeader className="text-center">
            <CardTitle className="text-xl">欢迎回来</CardTitle>
            <CardDescription>使用主密码解锁 Evorsio</CardDescription>
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
                <Button
                    type="submit"
                    disabled={loading}
                >
                  确认
                </Button>
                <FieldError>{error}</FieldError>
              </Field>
            </form>
          </CardContent>
        </Card>

      </div>
  )
}
