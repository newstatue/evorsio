import { createFileRoute, redirect } from "@tanstack/react-router"
import DashboardLayout from "@/features/dashboard"
import { authClient } from "@/lib/auth-client.ts"

export const Route = createFileRoute("/dashboard")({
  component: DashboardLayout,

  beforeLoad: async ()=>{
    const { data: session } =  await authClient.getSession()

    if(!session){
      throw redirect({
        to: "/login"
      })
    }
  }
})