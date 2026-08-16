import { Hono } from "hono"
import { authRoutes as auth } from "@/auth"
import { logger } from "hono/logger"
import { serveStatic } from "hono/bun"
import { cors } from "hono/cors"
import { WEB_URL } from "@/config"

const api = new Hono()
api.use("/*", cors({
  origin: WEB_URL,
  credentials: true,
}))

api.get("/", (ctx) => ctx.text("Hello Hono!"))
api.route("/auth", auth)

const app = new Hono()
app.use(logger())

app.route("/api", api)
app.use(
  "/*",
  serveStatic({
    root: "../web/dist",
  })
)
app.get(
  "/*",
  serveStatic({
    root: "../web/dist",
    path: "index.html",
  })
)


console.log("cwd:", process.cwd())

export type AppType = typeof app
export default app
