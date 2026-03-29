# TypeScript Backend — Level 2

> Covers: Zod validation · Vitest testing · Fastify · Utility Types · Project Structure · Debugging with Breakpoints

---

## Table of Contents

1. [Project Structure — The Right Way](#1-project-structure--the-right-way)
2. [Utility Types — Partial, Pick, Omit, Record](#2-utility-types--partialt-pickt-k-omitt-k-recordk-v)
3. [Zod — Runtime Validation + TypeScript Types](#3-zod--runtime-validation--typescript-types)
4. [Fastify — A TypeScript-Native API Framework](#4-fastify--a-typescript-native-api-framework)
5. [Vitest — Testing Your Backend](#5-vitest--testing-your-backend)
6. [Debugging with Breakpoints](#6-debugging-with-breakpoints)
7. [Putting It All Together](#7-putting-it-all-together)

---

## 1. Project Structure — The Right Way

Good structure is not about following one "correct" pattern — it's about making your codebase easy to navigate, test, and grow. Here is a proven layout for a TypeScript backend that scales well.

### The folder structure

```
my-api/
├── src/
│   ├── index.ts               ← entry point: starts the server
│   ├── app.ts                 ← creates and configures the app (no listen() call)
│   ├── config.ts              ← loads and validates env variables
│   │
│   ├── routes/                ← HTTP layer: defines endpoints, calls services
│   │   ├── users.ts
│   │   └── posts.ts
│   │
│   ├── services/              ← business logic: the core of your app
│   │   ├── userService.ts
│   │   └── postService.ts
│   │
│   ├── repositories/          ← data access: all database queries live here
│   │   ├── userRepository.ts
│   │   └── postRepository.ts
│   │
│   ├── schemas/               ← Zod schemas for validation and type inference
│   │   ├── userSchema.ts
│   │   └── postSchema.ts
│   │
│   ├── middleware/            ← reusable middleware (auth, logging, errors)
│   │   ├── authenticate.ts
│   │   └── errorHandler.ts
│   │
│   ├── errors/                ← custom error classes
│   │   └── index.ts
│   │
│   └── types/                 ← shared TypeScript types used across the app
│       └── index.ts
│
├── tests/
│   ├── unit/                  ← tests for services (no HTTP, no DB)
│   │   └── userService.test.ts
│   └── integration/           ← tests for routes (real HTTP requests)
│       └── users.test.ts
│
├── prisma/
│   └── schema.prisma
├── .env
├── .env.example               ← safe to commit — shows required variables
├── tsconfig.json
├── vitest.config.ts
└── package.json
```

### Why this structure works

**The key principle is layered separation:**

```
Request → Route → Service → Repository → Database
                     ↑
               Business logic
               lives only here
```

- **Routes** handle HTTP concerns only: parse the request, call a service, send the response. They should have almost no logic.
- **Services** contain your business logic: "a user can only have one active session", "posts by banned users are hidden". No database code here.
- **Repositories** contain all database queries. If you ever switch from Prisma to something else, you only touch this folder.
- **Schemas** define Zod validation schemas. Types are inferred from these — you never duplicate type definitions.

### Why `app.ts` and `index.ts` are separate

```typescript
// src/app.ts — builds the app, registers routes, no listen() call
import Fastify from "fastify";
import { usersRoutes } from "./routes/users";

export function buildApp() {
  const app = Fastify({ logger: true });
  app.register(usersRoutes, { prefix: "/users" });
  return app;
}
```

```typescript
// src/index.ts — only starts the server
import { buildApp } from "./app";
import { config } from "./config";

const app = buildApp();

app.listen({ port: config.port }, (err) => {
  if (err) {
    app.log.error(err);
    process.exit(1);
  }
});
```

The reason: your tests import `buildApp()` directly and call routes without starting a real server. If `listen()` was inside `app.ts`, every test would try to bind a port.

---

## 2. Utility Types — Partial\<T\>, Pick\<T, K\>, Omit\<T, K\>, Record\<K, V\>

TypeScript ships with a set of built-in "utility types" that transform existing types into new ones. These save you from duplicating type definitions and are used heavily in real-world backend code.

Start with a base type for all examples:

```typescript
interface User {
  id: number;
  name: string;
  email: string;
  password: string;
  role: string;
  createdAt: Date;
}
```

---

### `Partial<T>` — make all fields optional

Useful for **update operations** where the caller only sends the fields they want to change:

```typescript
type UserUpdate = Partial<User>;
// Equivalent to:
// {
//   id?: number;
//   name?: string;
//   email?: string;
//   password?: string;
//   role?: string;
//   createdAt?: Date;
// }

async function updateUser(id: number, data: Partial<User>): Promise<User> {
  // data can have any combination of User fields
  return prisma.user.update({ where: { id }, data });
}

// Valid — only sending what needs to change
await updateUser(1, { name: "Alice Updated" });
await updateUser(1, { email: "new@email.com", role: "admin" });
```

---

### `Pick<T, K>` — keep only certain fields

Useful when you want a **subset** of a type — for example, what you expose in an API response:

```typescript
type PublicUser = Pick<User, "id" | "name" | "email">;
// {
//   id: number;
//   name: string;
//   email: string;
// }

// Password and role are never accidentally leaked
function formatUserResponse(user: User): PublicUser {
  return {
    id: user.id,
    name: user.name,
    email: user.email,
  };
}
```

---

### `Omit<T, K>` — remove certain fields

The opposite of Pick — keep everything **except** specified fields:

```typescript
type UserWithoutPassword = Omit<User, "password">;
// All User fields except password

type CreateUserInput = Omit<User, "id" | "createdAt">;
// {
//   name: string;
//   email: string;
//   password: string;
//   role: string;
// }
// id and createdAt are generated by the DB, so the caller shouldn't provide them

async function createUser(data: CreateUserInput): Promise<User> {
  return prisma.user.create({ data });
}
```

---

### `Record<K, V>` — typed key-value maps

Creates an object type where keys are of type `K` and values are of type `V`:

```typescript
// A map of role name → list of allowed actions
type RolePermissions = Record<string, string[]>;

const permissions: RolePermissions = {
  admin:  ["read", "write", "delete"],
  editor: ["read", "write"],
  viewer: ["read"],
};

// More specific — restrict keys to a union of known strings
type Role = "admin" | "editor" | "viewer";
type StrictRolePermissions = Record<Role, string[]>;

// Now TypeScript enforces all three roles must be present
const strictPermissions: StrictRolePermissions = {
  admin:  ["read", "write", "delete"],
  editor: ["read", "write"],
  viewer: ["read"],
  // guest: ["read"]  ← Error: 'guest' is not assignable to type 'Role'
};
```

A very common pattern — caching database results in memory:

```typescript
// Map user IDs to user objects
const userCache: Record<number, User> = {};

function cacheUser(user: User): void {
  userCache[user.id] = user;
}

function getCachedUser(id: number): User | undefined {
  return userCache[id];
}
```

---

### Combining utility types

Real-world code chains these together:

```typescript
// For a PATCH /users/:id endpoint:
// - Must not allow changing id or createdAt (Omit)
// - All fields are optional since it's a partial update (Partial)
type PatchUserBody = Partial<Omit<User, "id" | "createdAt">>;

// For a safe public response — no password, no internal role
type SafeUserResponse = Omit<User, "password" | "role">;
```

---

## 3. Zod — Runtime Validation + TypeScript Types

TypeScript types only exist at compile time — they are erased when your code runs. This means when a request comes in from the network, TypeScript cannot guarantee the body has the right shape. Zod solves this by validating data **at runtime** and lets you derive TypeScript types from the same schema, so you never write the same shape twice.

### Install

```bash
npm install zod
```

### 3.1 Defining schemas and inferring types

```typescript
// src/schemas/userSchema.ts
import { z } from "zod";

// Define the schema — this is runtime validation
export const CreateUserSchema = z.object({
  name: z.string().min(2, "Name must be at least 2 characters"),
  email: z.string().email("Invalid email address"),
  password: z.string().min(8, "Password must be at least 8 characters"),
  role: z.enum(["admin", "editor", "viewer"]).default("viewer"),
});

export const UpdateUserSchema = z.object({
  name: z.string().min(2).optional(),
  email: z.string().email().optional(),
  role: z.enum(["admin", "editor", "viewer"]).optional(),
});

// Infer TypeScript types from the schemas — no duplication!
export type CreateUserInput = z.infer<typeof CreateUserSchema>;
export type UpdateUserInput = z.infer<typeof UpdateUserSchema>;

// CreateUserInput is now equivalent to:
// {
//   name: string;
//   email: string;
//   password: string;
//   role: "admin" | "editor" | "viewer";
// }
```

### 3.2 Parsing and handling errors

```typescript
import { CreateUserSchema } from "../schemas/userSchema";

const rawBody = {
  name: "Al",          // too short
  email: "not-an-email",
  password: "secret",  // too short
};

// .parse() — throws if invalid
try {
  const validated = CreateUserSchema.parse(rawBody);
} catch (error) {
  // error.errors is an array of field-level issues
  console.log(error.errors);
  // [
  //   { path: ["name"],     message: "Name must be at least 2 characters" },
  //   { path: ["email"],    message: "Invalid email address" },
  //   { path: ["password"], message: "Password must be at least 8 characters" },
  // ]
}

// .safeParse() — returns { success, data } or { success, error } — no try/catch needed
const result = CreateUserSchema.safeParse(rawBody);

if (!result.success) {
  console.log(result.error.errors); // structured error array
} else {
  const user = result.data; // fully typed as CreateUserInput
}
```

### 3.3 Useful Zod validators for backend work

```typescript
import { z } from "zod";

const PostSchema = z.object({
  title:     z.string().min(1).max(200),
  content:   z.string().optional(),
  published: z.boolean().default(false),
  tags:      z.array(z.string()).max(10),
  authorId:  z.number().int().positive(),
});

// Numeric route params come in as strings — coerce converts them
const IdParamSchema = z.object({
  id: z.coerce.number().int().positive("ID must be a positive integer"),
});

// Query string validation
const PaginationSchema = z.object({
  page:  z.coerce.number().int().min(1).default(1),
  limit: z.coerce.number().int().min(1).max(100).default(20),
});

type PaginationQuery = z.infer<typeof PaginationSchema>;
// { page: number; limit: number }
```

### 3.4 Transforming data with Zod

Zod can clean and transform data as part of parsing:

```typescript
const UserSchema = z.object({
  name:  z.string().trim(),                        // strips whitespace
  email: z.string().email().toLowerCase(),          // normalises email
  age:   z.number().int().min(0).max(150).nullable(),
});
```

---

## 4. Fastify — A TypeScript-Native API Framework

Fastify is faster than Express, has built-in TypeScript support, and uses a schema-first approach. It is the better choice for new TypeScript projects.

### Install

```bash
npm install fastify
npm install --save-dev @fastify/type-provider-zod
npm install @fastify/type-provider-zod zod
```

### 4.1 Building the app

```typescript
// src/app.ts
import Fastify from "fastify";
import {
  serializerCompiler,
  validatorCompiler,
  ZodTypeProvider,
} from "fastify-type-provider-zod";
import { usersRoutes } from "./routes/users";
import { errorHandler } from "./middleware/errorHandler";

export function buildApp() {
  const app = Fastify({ logger: true });

  // Wire up Zod as the validator/serializer
  app.setValidatorCompiler(validatorCompiler);
  app.setSerializerCompiler(serializerCompiler);

  // Global error handler
  app.setErrorHandler(errorHandler);

  // Register route modules with a URL prefix
  app.withTypeProvider<ZodTypeProvider>().register(usersRoutes, {
    prefix: "/users",
  });

  return app;
}
```

### 4.2 Typed routes with Zod schemas

This is where Fastify + Zod shines. You define the schema once and get:
- Automatic request validation (400 if invalid)
- Full TypeScript types on `request.body`, `request.params`, `request.query`
- Auto-generated API documentation (if using Swagger plugin)

```typescript
// src/routes/users.ts
import { FastifyPluginAsyncZod } from "fastify-type-provider-zod";
import { z } from "zod";
import { CreateUserSchema, UpdateUserSchema } from "../schemas/userSchema";
import * as userService from "../services/userService";

const IdParam = z.object({
  id: z.coerce.number().int().positive(),
});

export const usersRoutes: FastifyPluginAsyncZod = async (app) => {

  // GET /users
  app.get("/", {
    schema: {
      response: { 200: z.array(CreateUserSchema) },
    },
  }, async (request, reply) => {
    const users = await userService.getAllUsers();
    return users;
  });

  // GET /users/:id
  app.get("/:id", {
    schema: {
      params: IdParam,
    },
  }, async (request, reply) => {
    // request.params.id is typed as number — guaranteed
    const user = await userService.getUserById(request.params.id);
    if (!user) return reply.status(404).send({ error: "User not found" });
    return user;
  });

  // POST /users
  app.post("/", {
    schema: {
      body: CreateUserSchema,
    },
  }, async (request, reply) => {
    // request.body is fully typed as CreateUserInput
    const user = await userService.createUser(request.body);
    return reply.status(201).send(user);
  });

  // PATCH /users/:id
  app.patch("/:id", {
    schema: {
      params: IdParam,
      body: UpdateUserSchema,
    },
  }, async (request, reply) => {
    const user = await userService.updateUser(
      request.params.id,
      request.body
    );
    return user;
  });

  // DELETE /users/:id
  app.delete("/:id", {
    schema: { params: IdParam },
  }, async (request, reply) => {
    await userService.deleteUser(request.params.id);
    return reply.status(204).send();
  });
};
```

### 4.3 Error handler middleware

```typescript
// src/middleware/errorHandler.ts
import { FastifyError, FastifyRequest, FastifyReply } from "fastify";

export function errorHandler(
  error: FastifyError,
  request: FastifyRequest,
  reply: FastifyReply
): void {
  // Zod/Fastify validation errors
  if (error.statusCode === 400) {
    reply.status(400).send({
      success: false,
      error: "Validation failed",
      details: error.message,
    });
    return;
  }

  // Log unexpected errors
  request.log.error(error);

  reply.status(error.statusCode ?? 500).send({
    success: false,
    error: "Internal server error",
  });
}
```

---

## 5. Vitest — Testing Your Backend

Vitest is a modern test runner with native TypeScript support. It uses the same syntax as Jest, so it's easy to pick up.

### Install

```bash
npm install --save-dev vitest
```

### Configure — `vitest.config.ts`

```typescript
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    globals: true,           // no need to import describe/it/expect
    environment: "node",
    coverage: {
      provider: "v8",
      reporter: ["text", "html"],
    },
  },
});
```

Add scripts to `package.json`:

```json
{
  "scripts": {
    "test":          "vitest run",
    "test:watch":    "vitest",
    "test:coverage": "vitest run --coverage"
  }
}
```

---

### 5.1 Unit tests — testing services in isolation

Unit tests test a single function without a database or HTTP server. You **mock** the repository so the service thinks it's hitting a database but isn't.

```typescript
// tests/unit/userService.test.ts
import { describe, it, expect, vi, beforeEach } from "vitest";
import * as userRepository from "../../src/repositories/userRepository";
import * as userService from "../../src/services/userService";

// Mock the entire repository module
vi.mock("../../src/repositories/userRepository");

describe("userService", () => {
  beforeEach(() => {
    vi.clearAllMocks(); // reset mocks before each test
  });

  describe("getUserById", () => {
    it("returns a user when found", async () => {
      const mockUser = { id: 1, name: "Alice", email: "alice@example.com" };

      // Tell the mock what to return
      vi.mocked(userRepository.findById).mockResolvedValue(mockUser);

      const result = await userService.getUserById(1);

      expect(result).toEqual(mockUser);
      expect(userRepository.findById).toHaveBeenCalledWith(1);
      expect(userRepository.findById).toHaveBeenCalledTimes(1);
    });

    it("throws NotFoundError when user does not exist", async () => {
      vi.mocked(userRepository.findById).mockResolvedValue(null);

      await expect(userService.getUserById(999)).rejects.toThrow("User not found");
    });
  });

  describe("createUser", () => {
    it("hashes the password before saving", async () => {
      const input = { name: "Bob", email: "bob@example.com", password: "plaintext" };
      const savedUser = { id: 2, ...input, password: "hashed_password", role: "viewer" };

      vi.mocked(userRepository.create).mockResolvedValue(savedUser);

      const result = await userService.createUser(input);

      // The saved password should NOT be the plain text original
      expect(userRepository.create).toHaveBeenCalledWith(
        expect.objectContaining({
          password: expect.not.stringContaining("plaintext"),
        })
      );
    });
  });
});
```

---

### 5.2 Integration tests — testing routes end to end

Integration tests spin up your real Fastify app (using `buildApp()`) and send real HTTP requests to it. No mocking — you test the whole stack except the database, which you can use a test database for or mock at the repository level.

```typescript
// tests/integration/users.test.ts
import { describe, it, expect, beforeAll, afterAll } from "vitest";
import { buildApp } from "../../src/app";
import type { FastifyInstance } from "fastify";

let app: FastifyInstance;

beforeAll(async () => {
  app = buildApp();
  await app.ready(); // waits for all plugins to load
});

afterAll(async () => {
  await app.close();
});

describe("POST /users", () => {
  it("creates a user with valid input", async () => {
    const response = await app.inject({
      method: "POST",
      url: "/users",
      payload: {
        name: "Alice",
        email: "alice@example.com",
        password: "securepassword",
      },
    });

    expect(response.statusCode).toBe(201);

    const body = response.json();
    expect(body.name).toBe("Alice");
    expect(body.email).toBe("alice@example.com");
    expect(body.password).toBeUndefined(); // never returned
  });

  it("returns 400 for invalid email", async () => {
    const response = await app.inject({
      method: "POST",
      url: "/users",
      payload: {
        name: "Alice",
        email: "not-an-email",
        password: "securepassword",
      },
    });

    expect(response.statusCode).toBe(400);
  });

  it("returns 400 when required fields are missing", async () => {
    const response = await app.inject({
      method: "POST",
      url: "/users",
      payload: { name: "Alice" }, // missing email and password
    });

    expect(response.statusCode).toBe(400);
  });
});

describe("GET /users/:id", () => {
  it("returns 404 for non-existent user", async () => {
    const response = await app.inject({
      method: "GET",
      url: "/users/99999",
    });

    expect(response.statusCode).toBe(404);
  });
});
```

---

### 5.3 Testing Zod schemas directly

Zod schemas are pure functions — they're easy to test on their own:

```typescript
// tests/unit/userSchema.test.ts
import { describe, it, expect } from "vitest";
import { CreateUserSchema } from "../../src/schemas/userSchema";

describe("CreateUserSchema", () => {
  it("accepts valid input", () => {
    const result = CreateUserSchema.safeParse({
      name: "Alice",
      email: "alice@example.com",
      password: "securepassword",
    });
    expect(result.success).toBe(true);
  });

  it("rejects short names", () => {
    const result = CreateUserSchema.safeParse({
      name: "A",
      email: "alice@example.com",
      password: "securepassword",
    });
    expect(result.success).toBe(false);
    expect(result.error?.issues[0].path).toContain("name");
  });

  it("defaults role to viewer", () => {
    const result = CreateUserSchema.safeParse({
      name: "Alice",
      email: "alice@example.com",
      password: "securepassword",
    });
    expect(result.success).toBe(true);
    expect(result.data?.role).toBe("viewer");
  });
});
```

---

## 6. Debugging with Breakpoints

A breakpoint pauses your running program at a specific line so you can inspect variables, step through code line by line, and understand exactly what's happening. No more `console.log` everywhere.

### 6.1 Setup in VS Code

Create the file `.vscode/launch.json` in your project root:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug: Dev Server",
      "type": "node",
      "request": "launch",
      "runtimeExecutable": "node",
      "runtimeArgs": ["--import", "tsx/esm"],
      "args": ["src/index.ts"],
      "sourceMaps": true,
      "cwd": "${workspaceFolder}",
      "env": { "NODE_ENV": "development" }
    },
    {
      "name": "Debug: Current Test File",
      "type": "node",
      "request": "launch",
      "runtimeExecutable": "${workspaceFolder}/node_modules/.bin/vitest",
      "args": ["run", "${relativeFile}"],
      "sourceMaps": true,
      "cwd": "${workspaceFolder}",
      "smartStep": true
    },
    {
      "name": "Debug: All Tests",
      "type": "node",
      "request": "launch",
      "runtimeExecutable": "${workspaceFolder}/node_modules/.bin/vitest",
      "args": ["run"],
      "sourceMaps": true,
      "cwd": "${workspaceFolder}"
    }
  ]
}
```

### 6.2 Setting breakpoints

1. Open any `.ts` file in VS Code
2. Click in the **gutter** (the margin to the left of the line numbers) — a red dot appears
3. Press `F5` (or go to **Run → Start Debugging**) and select a configuration
4. Your program runs until it hits the breakpoint and pauses

### 6.3 The Debug toolbar — what each button does

When paused at a breakpoint, a toolbar appears at the top of VS Code:

```
▶  Continue (F5)         — resume running until the next breakpoint
⤵  Step Over (F10)       — execute the current line, move to next (doesn't enter function calls)
⤷  Step Into (F11)       — step inside the function being called on this line
⤴  Step Out (Shift+F11)  — finish the current function, pause in the caller
⟳  Restart              — restart the whole debug session
■  Stop                  — end the session
```

### 6.4 The Debug panel — what you can see while paused

**Variables panel** — all variables in the current scope, their types and values. You can expand objects to see nested fields. You can even edit variable values live.

**Watch panel** — add any expression (e.g. `user.email`, `users.length`, `request.body`) and it evaluates in real time as you step through.

**Call Stack panel** — shows you how you got here. If you're inside `getUserById`, you can see it was called by `usersRoute`, which was called by Fastify's request handler. Click any frame to jump to that code.

**Debug Console** — a REPL that runs in the paused context. Type any expression and it evaluates with the current variables in scope.

### 6.5 A practical debugging walkthrough

Say a route is returning 400 unexpectedly. Here's how to investigate:

```typescript
// src/routes/users.ts

app.post("/", {
  schema: { body: CreateUserSchema },
}, async (request, reply) => {
  // 1. Set a breakpoint on this line ←
  const user = await userService.createUser(request.body);
  return reply.status(201).send(user);
});
```

```typescript
// src/services/userService.ts

export async function createUser(input: CreateUserInput): Promise<User> {
  // 2. Set another breakpoint here ←
  const hashed = await hashPassword(input.password);
  return userRepository.create({ ...input, password: hashed });
}
```

**Steps:**
1. Start debug session with `F5`
2. Send a POST request (use curl, Postman, or your integration test)
3. Execution pauses at breakpoint 1 — inspect `request.body` in the Variables panel to confirm the data arrived correctly
4. Press F10 (Step Over) or F11 (Step Into) — execution moves to breakpoint 2
5. Inspect `input` — does it look right?
6. Step through `hashPassword` — check the output
7. Step into `userRepository.create` — see exactly what's being sent to Prisma

### 6.6 Conditional breakpoints

Right-click a breakpoint dot → **Edit Breakpoint** → add a condition. The debugger only pauses when the condition is true:

```
request.params.id === 42
users.length > 100
error.statusCode !== 500
```

This is invaluable when debugging loops or routes that get called many times.

### 6.7 Logpoints — breakpoints that don't stop

Right-click a gutter → **Add Logpoint**. Enter a message with `{}` for expressions:

```
User created: {user.id} — {user.email}
Request body: {JSON.stringify(request.body)}
```

The message is printed to the Debug Console without pausing execution. A non-intrusive `console.log` that you never forget to remove.

---

## 7. Putting It All Together

Here is how all the pieces connect in a real request lifecycle:

```
POST /users  { name: "Alice", email: "alice@example.com", password: "secret" }

  │
  ▼
Fastify receives request
  │
  ▼
Zod validates request.body against CreateUserSchema
  ├─ Invalid → 400 response with field errors (automatic, no code needed)
  └─ Valid   → request.body is typed as CreateUserInput
  │
  ▼
Route handler calls userService.createUser(request.body)
  │
  ▼
Service applies business logic (hash password, assign default role)
  │
  ▼
Repository calls prisma.user.create(data)
  │
  ▼
Prisma inserts row, returns typed User object
  │
  ▼
Service returns User to route
  │
  ▼
Route sends 201 response with user (password field omitted)
```

### The stack summary

| Layer        | Tool          | Responsibility                          |
|--------------|---------------|-----------------------------------------|
| Framework    | Fastify        | HTTP, routing, request lifecycle        |
| Validation   | Zod            | Schema definition, type inference, parsing |
| Business     | Service layer  | Rules, logic, orchestration             |
| Data         | Prisma         | Database queries, generated types       |
| Testing      | Vitest         | Unit and integration tests              |
| Debugging    | VS Code + tsx  | Breakpoints, variable inspection        |

### Recommended learning order

1. Get the project structure in place before writing any code
2. Define your Zod schemas first — types flow from them
3. Write the repository layer (database queries)
4. Write the service layer with unit tests alongside it
5. Wire up Fastify routes last and write integration tests
6. Use breakpoints as soon as something doesn't behave as expected — don't guess

---

*The jump from tutorial code to production code is mostly about this: validation at the boundary, clear separation of concerns, and tests that let you refactor without fear.*
