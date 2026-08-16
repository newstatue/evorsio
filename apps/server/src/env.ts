import { auth } from "@/auth/auth"

export type Env = {
  Variables: {
    session: typeof auth.$Infer.Session | null
  }
}