
import { auth } from "@/auth/auth"
import { createMiddleware } from "hono/factory"
import { Env } from "@/env"

export const session=  createMiddleware<Env>(
  async (ctx, next) =>{
  const session = await auth.api.getSession({
    headers: ctx.req.raw.headers,
  })

  ctx.set("session", session);

  await next();
})