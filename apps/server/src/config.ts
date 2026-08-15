export const PORT = process.env.PORT ?? "8080"

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
