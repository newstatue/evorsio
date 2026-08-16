export const PRODUCTION = "production"
export const DEVELOPMENT = "development"

export const NODE_ENV = process.env.NODE_ENV ?? PRODUCTION
export const IS_DEVELOPMENT = NODE_ENV === DEVELOPMENT

export const PORT = IS_DEVELOPMENT ? "" : requiredEnv("PORT")
export const DATABASE_URL = IS_DEVELOPMENT ? "": requiredEnv("DATABASE_URL")

export const WEB_URL = requiredEnv("WEB_URL")
export const BETTER_AUTH_URL = requiredEnv("BETTER_AUTH_URL")
export const BETTER_AUTH_SECRET = requiredEnv("BETTER_AUTH_SECRET")
export const GITHUB_CLIENT_ID = requiredEnv("GITHUB_CLIENT_ID")
export const GITHUB_CLIENT_SECRET = requiredEnv("GITHUB_CLIENT_SECRET")

export const DEEPSEEK_API_KEY = requiredEnv("DEEPSEEK_API_KEY")

function requiredEnv(name: string): string {
  const value = process.env[name]

  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`)
  }

  return value
}
