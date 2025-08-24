import { serve } from "@hono/node-server";
import { Hono } from "hono";
import { auth } from "../auth";
import { cors } from "hono/cors";
import { logger } from "hono/logger";

const app = new Hono();

app.use(logger());
app.use(
  cors({
    origin: "http://localhost:5173",
    allowHeaders: [
      "Content-Type",
      "Grpc-Timeout",
      "X-Grpc-Web",
      "X-User-Agent",
    ],
    allowMethods: ["POST", "GET", "OPTIONS"],
    exposeHeaders: ["Content-Length"],
    maxAge: 600,
    credentials: true,
  }),
);

app.on(["POST", "GET"], "/api/auth/*", (c) => {
  return auth.handler(c.req.raw);
});

const getToken = async (headers: Headers) => {
  try {
    const { token } = await auth.api.getToken({ headers });
    return token;
  } catch (error) {
    return null;
  }
};

app.on(["POST"], "*", async (c) => {
  const headers = c.req.raw.headers;

  const token = await getToken(headers);

  if (token) {
    headers.append("Authorization", `Bearer ${token}`);
  }

  const res = await fetch(`http://localhost:3001${c.req.path}`, {
    method: c.req.method,
    headers,
    body: await c.req.arrayBuffer(),
  });

  return new Response(res.body, {
    status: res.status,
    statusText: res.statusText,
    headers: res.headers,
  });
});

serve(
  {
    fetch: app.fetch,
    port: 3000,
  },
  (info) => {
    console.log(`Server is running on http://localhost:${info.port}`);
  },
);
