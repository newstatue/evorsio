import { createFileRoute } from "@tanstack/react-router"
import {Button} from "@/components/ui/button.tsx";
import {CreateEntry, GeneratePassword} from "@bindings/github.com/newstatue/evorsio/internal/vault/service.ts";
import {ButtonGroup} from "@/components/ui/button-group.tsx";
import {Input} from "@/components/ui/input.tsx";
import {useState} from "react";

export const Route = createFileRoute("/dash/vault/")({
  component: RouteComponent,
  staticData: {
    breadcrumb: "密钥",
  },
})

function RouteComponent() {
  const [password, setPassword] = useState("")

  const create = async () => {
    await CreateEntry({
      Title:"hello", Notes: "beizhu", Password: "pass", URL: "https://github.com", Username: "wyx"
    })
  }


  const generate = async () => {
    const result = await GeneratePassword({Length: 16, Lowercase: true, Numbers: true, Symbols: true, Uppercase: true})
    setPassword(result)
  }

  return (
      <>
        <ButtonGroup>
          <Button onClick={create}>创建</Button>


          <Button onClick={generate}>生成密码</Button>
        </ButtonGroup>
        <Input value={password} readOnly />
      </>
  )
}
