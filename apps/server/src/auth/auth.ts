import { betterAuth } from "better-auth/minimal"
import { drizzleAdapter } from "better-auth/adapters/drizzle"
import { db } from "@/db"
import { GITHUB_CLIENT_ID, GITHUB_CLIENT_SECRET, WEB_URL } from "@/config"

export const auth = betterAuth({
  database: drizzleAdapter(db, {
    provider: "pg",
  }),
  emailAndPassword: {
    enabled: true,
  },
  trustedOrigins: [WEB_URL],
  socialProviders: {
    github: {
      clientId: GITHUB_CLIENT_ID,
      clientSecret: GITHUB_CLIENT_SECRET,
    },
  },
})
