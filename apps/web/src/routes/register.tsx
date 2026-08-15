import { createFileRoute } from "@tanstack/react-router"
import RegisterPage from "@/features/register"

export const Route = createFileRoute("/register")({
  component: RegisterPage,
})
