import { drizzle as drizzlePgSQL } from "drizzle-orm/node-postgres"
import { DATABASE_URL, IS_DEVELOPMENT } from "@/config"

export const db = IS_DEVELOPMENT ? await development() : drizzlePgSQL(DATABASE_URL)

async function development() {
  const [{ PGlite }, { drizzle }] = await Promise.all([import("@electric-sql/pglite"), import("drizzle-orm/pglite")])

  return drizzle(new PGlite("./data"))
}
