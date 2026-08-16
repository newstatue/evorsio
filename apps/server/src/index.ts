import { Hono } from "hono"
import { authRoutes as auth } from "@/auth"
import { logger } from "hono/logger"
import { serveStatic } from "hono/bun"

const api = new Hono()
api.get("/", (ctx) => ctx.text("Hello Hono!"))
api.route("/auth", auth)
const app = new Hono()
app.use(logger())

app.use("/*", serveStatic({
  root: "../web/dist"
}))
app.get("/*",serveStatic({
  root: "../web/dist",
  path: "index.html",
}))
app.route("/api", api)

console.log("cwd:", process.cwd())

export type AppType = typeof app
export default app
