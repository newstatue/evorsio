import { Hono } from "hono"
import { authRoutes as auth } from "@/auth"
import { logger } from "hono/logger"

const api = new Hono()
api.get("/", (ctx) => ctx.text("Hello Hono!"))
api.route("/auth", auth)
const app = new Hono()
app.use(logger())

app.route("/api", api)

Bun.serve({
  fetch: app.fetch,
})

export type AppType = typeof app
export default app
