import { betterAuth } from "better-auth";
import { jwt } from "better-auth/plugins";
import { Pool } from "pg";
import { env } from "./env";
import { admin } from "better-auth/plugins";

export const auth = betterAuth({
  plugins: [jwt(), admin()],
  database: new Pool({
    connectionString: env.DATABASE_URL,
  }),
  advanced: {
    generateId: false,
  },
  emailAndPassword: {
    enabled: true,
  },
  trustedOrigins: ["http://localhost:5173"],
});
