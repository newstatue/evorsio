import { Hono } from "hono"
import { Env } from "@/env"

export const user = new Hono<Env>()

user.get("/", async (ctx) => {
  const session = ctx.get("session")

  if (!session) {
    return ctx.json({message:"Unauthorized"},401)
  }

  return ctx.json({
     ...session.user,
  })
})
