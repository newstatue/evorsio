import { createFileRoute } from "@tanstack/react-router"
import DashboardLayout from "@/features/dashboard"

export const Route = createFileRoute("/dashboard")({
  component: DashboardLayout,
})
