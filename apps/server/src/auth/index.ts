import { Hono } from "hono"
import { auth } from "@/auth/auth"

export const authRoutes = new Hono()

authRoutes.on(["POST", "GET"], "/*", (ctx) => auth.handler(ctx.req.raw))
