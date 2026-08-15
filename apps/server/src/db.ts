import { PGlite } from "@electric-sql/pglite"
import { drizzle } from "drizzle-orm/pglite"
import * as schema from "@/auth/auth-schema"

const client = new PGlite("./data")
export const db = drizzle(client, { schema })
