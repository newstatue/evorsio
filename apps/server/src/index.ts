import { Hono } from "hono"
import { auth } from "@/auth"
import { logger } from "hono/logger"
import { serveStatic } from "hono/bun"
import { cors } from "hono/cors"
import { WEB_URL } from "@/config"
import { session } from "@/auth/auth-middleware"
import { user } from "@/user"

const api = new Hono()
api.use("/*", cors({
  origin: WEB_URL,
  credentials: true,
}))
api.use("/*", session)

api.get("/", (ctx) => ctx.text("Hello Hono!"))
api.route("/auth", auth)
api.route("/user",user)

const app = new Hono()
app.use(logger())

app.route("/api", api)
app.get("/api/*",(ctx) => ctx.text("Not Found"))

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
