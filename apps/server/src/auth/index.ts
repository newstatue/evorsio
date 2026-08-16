import { Hono } from "hono"
import { auth as authInstance } from "@/auth/auth"

export const auth = new Hono()

auth.on(["POST", "GET"], "/*", (ctx) => authInstance.handler(ctx.req.raw))
