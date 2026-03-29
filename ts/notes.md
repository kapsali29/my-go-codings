# TypeScript for Backend Development — A Strong Start

> A beginner-friendly, practical guide covering TypeScript fundamentals, Node.js patterns, REST APIs, databases, and CLI tools.

---

## Table of Contents

1. [What is TypeScript and Why Use It?](#1-what-is-typescript-and-why-use-it)
2. [Setting Up Your Environment](#2-setting-up-your-environment)
3. [TypeScript Fundamentals](#3-typescript-fundamentals)
4. [General Node.js Patterns in TypeScript](#4-general-nodejs-patterns-in-typescript)
5. [Building a REST API with Express](#5-building-a-rest-api-with-express)
6. [Databases with Prisma](#6-databases-with-prisma)
7. [CLI Tools & Scripting](#7-cli-tools--scripting)
8. [What to Learn Next](#8-what-to-learn-next)

---

## 1. What is TypeScript and Why Use It?

JavaScript is a flexible, dynamic language — which is great for quick scripts but painful in larger projects. You can accidentally pass a number where a string was expected, call a method that doesn't exist, or mistype a property name and only find out at runtime (often in production).

**TypeScript** is a superset of JavaScript that adds a *type system*. This means:

- You declare what kind of data a variable holds (`string`, `number`, an object with specific fields, etc.)
- TypeScript checks your code *before* you run it and tells you about mistakes
- Your editor (VS Code especially) gets powerful autocomplete, inline docs, and refactoring tools

```
JavaScript:  write → run → crash → fix
TypeScript:  write → type error caught → fix → run confidently
```

TypeScript compiles down to plain JavaScript, so it runs anywhere Node.js runs. You write `.ts` files, compile them to `.js`, and Node executes the `.js`. In modern setups, tools like `tsx` let you run `.ts` files directly without a separate compile step.

---

## 2. Setting Up Your Environment

### What you need

- **Node.js** (v18 or later) — download from nodejs.org
- **VS Code** — the best editor for TypeScript by far
- **npm** — comes bundled with Node.js

### Create your first project

```bash
mkdir ts-backend && cd ts-backend
npm init -y
```

### Install TypeScript and a dev runner

```bash
npm install --save-dev typescript tsx @types/node
```

- `typescript` — the TypeScript compiler (`tsc`)
- `tsx` — runs `.ts` files directly during development (no manual compile step)
- `@types/node` — TypeScript type definitions for Node.js built-ins (`fs`, `path`, `process`, etc.)

### Create the TypeScript config file

```bash
npx tsc --init
```

This creates `tsconfig.json`. Replace its contents with this practical starter config:

```json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "CommonJS",
    "rootDir": "./src",
    "outDir": "./dist",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules"]
}
```

**Key settings explained:**
- `strict: true` — enables all strict type checks. Leave this on — it's the whole point
- `rootDir` — your TypeScript source files live in `src/`
- `outDir` — compiled JavaScript goes into `dist/`
- `esModuleInterop` — makes importing CommonJS modules (like Express) cleaner

### Your project structure

```
ts-backend/
├── src/
│   └── index.ts       ← your code goes here
├── dist/              ← compiled output (auto-generated)
├── tsconfig.json
└── package.json
```

### Add scripts to package.json

```json
{
  "scripts": {
    "dev": "tsx watch src/index.ts",
    "build": "tsc",
    "start": "node dist/index.js"
  }
}
```

- `npm run dev` — runs your app in watch mode (restarts on file change)
- `npm run build` — compiles to JavaScript
- `npm start` — runs the compiled output

### Your first TypeScript file

Create `src/index.ts`:

```typescript
const message: string = "TypeScript backend is running!";
console.log(message);
```

Run it:

```bash
npm run dev
```

---

## 3. TypeScript Fundamentals

### 3.1 Basic Types

TypeScript has a handful of core types you'll use constantly:

```typescript
// Primitive types
let name: string = "Alice";
let age: number = 30;
let isActive: boolean = true;

// Arrays
let scores: number[] = [90, 85, 78];
let tags: string[] = ["node", "typescript", "api"];

// A value that could be one of several types (union type)
let id: string | number = "abc123";
id = 42; // also valid
```

You don't always need to write the type — TypeScript can **infer** it:

```typescript
let city = "Zurich";  // TypeScript knows this is a string
city = 99;            // Error: Type 'number' is not assignable to type 'string'
```

### 3.2 Objects and Interfaces

In backend work, you're constantly dealing with structured data — users, orders, products. You define the shape of objects using `interface` or `type`:

```typescript
interface User {
  id: number;
  name: string;
  email: string;
  isAdmin: boolean;
}

// TypeScript will catch if you miss a field or mistype a property name
const user: User = {
  id: 1,
  name: "Alice",
  email: "alice@example.com",
  isAdmin: false,
};

console.log(user.email); // "alice@example.com"
console.log(user.phone); // Error: Property 'phone' does not exist on type 'User'
```

#### Optional properties

Use `?` for fields that may or may not be present:

```typescript
interface Product {
  id: number;
  name: string;
  description?: string; // optional
  price: number;
}
```

#### The difference between `interface` and `type`

Both are used to describe object shapes. As a beginner, they're largely interchangeable. A simple rule of thumb:
- Use `interface` for objects and class shapes
- Use `type` for unions, primitives, or when you need more complex type expressions

```typescript
// type alias — good for unions
type Status = "pending" | "active" | "inactive";

let userStatus: Status = "active";
userStatus = "banned"; // Error: not a valid Status
```

### 3.3 Functions

Always type your function parameters and return value:

```typescript
function add(a: number, b: number): number {
  return a + b;
}

function greet(name: string): string {
  return `Hello, ${name}!`;
}

// Void return type — function returns nothing
function logMessage(message: string): void {
  console.log(message);
}
```

#### Arrow functions

```typescript
const multiply = (a: number, b: number): number => a * b;
```

#### Optional and default parameters

```typescript
function createUser(name: string, role: string = "user", age?: number): User {
  // age may be undefined if not passed
  return { id: Date.now(), name, email: "", isAdmin: role === "admin" };
}
```

### 3.4 Enums

Enums give names to a set of related constants — great for status codes, roles, etc.:

```typescript
enum UserRole {
  Admin = "ADMIN",
  Editor = "EDITOR",
  Viewer = "VIEWER",
}

interface User {
  id: number;
  name: string;
  role: UserRole;
}

const user: User = {
  id: 1,
  name: "Alice",
  role: UserRole.Admin,
};

if (user.role === UserRole.Admin) {
  console.log("Full access granted");
}
```

### 3.5 Generics — Write Reusable Code

Generics let you write functions that work with any type while still being type-safe. You'll encounter these constantly in libraries and in your own utilities:

```typescript
// Without generics — only works for numbers
function getFirstNumber(arr: number[]): number {
  return arr[0];
}

// With generics — works for any type
function getFirst<T>(arr: T[]): T {
  return arr[0];
}

const firstNumber = getFirst<number>([10, 20, 30]);  // 10
const firstName = getFirst<string>(["Alice", "Bob"]); // "Alice"
```

A very common real-world use: wrapping API responses in a consistent structure:

```typescript
interface ApiResponse<T> {
  data: T;
  success: boolean;
  message: string;
}

// Now you can type any response body
type UserResponse = ApiResponse<User>;
type UsersResponse = ApiResponse<User[]>;
```

### 3.6 Type Assertions and `unknown`

Sometimes you receive data from outside your system (JSON from an API, user input) and TypeScript doesn't know its type. Use `unknown` — it's the safe version of `any`:

```typescript
// BAD — 'any' disables type checking entirely
function parseData(input: any) {
  return input.name.toUpperCase(); // TypeScript won't catch if .name doesn't exist
}

// GOOD — 'unknown' forces you to check before using
function parseData(input: unknown): string {
  if (typeof input === "object" && input !== null && "name" in input) {
    return (input as { name: string }).name.toUpperCase();
  }
  throw new Error("Invalid input");
}
```

---

## 4. General Node.js Patterns in TypeScript

### 4.1 Working with the File System

```typescript
import * as fs from "fs/promises"; // async file system
import * as path from "path";

async function readConfig(filename: string): Promise<object> {
  const filePath = path.join(__dirname, filename);
  const raw = await fs.readFile(filePath, "utf-8");
  return JSON.parse(raw);
}

async function writeLog(message: string): Promise<void> {
  const logPath = path.join(__dirname, "app.log");
  const entry = `[${new Date().toISOString()}] ${message}\n`;
  await fs.appendFile(logPath, entry);
}
```

### 4.2 Async/Await

All serious Node.js backend code is asynchronous. TypeScript makes this clean:

```typescript
// A function that returns a Promise uses Promise<T> as its return type
async function fetchUser(id: number): Promise<User | null> {
  try {
    // imagine this calls a database
    const user = await database.findUserById(id);
    return user;
  } catch (error) {
    console.error("Failed to fetch user:", error);
    return null;
  }
}
```

### 4.3 Environment Variables

Never hardcode secrets. Use `process.env` and validate your config at startup:

```bash
npm install dotenv
```

Create a `.env` file (never commit this to git):
```
PORT=3000
DATABASE_URL=postgresql://user:pass@localhost:5432/mydb
JWT_SECRET=supersecretkey
```

Create `src/config.ts`:

```typescript
import * as dotenv from "dotenv";
dotenv.config();

interface Config {
  port: number;
  databaseUrl: string;
  jwtSecret: string;
}

function loadConfig(): Config {
  const port = process.env.PORT;
  const databaseUrl = process.env.DATABASE_URL;
  const jwtSecret = process.env.JWT_SECRET;

  if (!databaseUrl) throw new Error("DATABASE_URL is required");
  if (!jwtSecret) throw new Error("JWT_SECRET is required");

  return {
    port: port ? parseInt(port, 10) : 3000,
    databaseUrl,
    jwtSecret,
  };
}

export const config = loadConfig();
```

Now import `config` anywhere in your app — TypeScript knows exactly what fields it has.

### 4.4 Error Handling Patterns

Define custom error classes so you can handle different errors differently:

```typescript
class AppError extends Error {
  constructor(
    public message: string,
    public statusCode: number = 500
  ) {
    super(message);
    this.name = "AppError";
  }
}

class NotFoundError extends AppError {
  constructor(resource: string) {
    super(`${resource} not found`, 404);
  }
}

class ValidationError extends AppError {
  constructor(message: string) {
    super(message, 400);
  }
}

// Usage
function getUser(id: number): User {
  const user = db.find(id);
  if (!user) throw new NotFoundError("User");
  return user;
}
```

---

## 5. Building a REST API with Express

### 5.1 Install Express with TypeScript support

```bash
npm install express
npm install --save-dev @types/express
```

`@types/express` gives TypeScript the type definitions for Express — so it knows the shape of `Request`, `Response`, `NextFunction`, etc.

### 5.2 A typed Express server

```typescript
// src/index.ts
import express, { Request, Response, NextFunction } from "express";
import { config } from "./config";

const app = express();
app.use(express.json()); // parse JSON request bodies

app.listen(config.port, () => {
  console.log(`Server running on port ${config.port}`);
});
```

### 5.3 Typed request bodies and params

This is where TypeScript really shines in Express — you can type exactly what you expect to receive:

```typescript
interface CreateUserBody {
  name: string;
  email: string;
  role?: string;
}

interface UserParams {
  id: string;
}

// POST /users
app.post("/users", (req: Request<{}, {}, CreateUserBody>, res: Response) => {
  const { name, email, role } = req.body;
  // TypeScript knows these fields exist and their types
  const newUser: User = {
    id: Date.now(),
    name,
    email,
    isAdmin: role === "admin",
  };
  res.status(201).json(newUser);
});

// GET /users/:id
app.get("/users/:id", (req: Request<UserParams>, res: Response) => {
  const userId = parseInt(req.params.id, 10);
  // fetch from DB, etc.
  res.json({ id: userId });
});
```

### 5.4 Organising routes — controllers and routers

Don't put everything in `index.ts`. Split by resource:

```typescript
// src/routes/users.ts
import { Router, Request, Response } from "express";

const router = Router();

// In-memory store for this example
const users: User[] = [];

router.get("/", (req: Request, res: Response) => {
  res.json(users);
});

router.post("/", (req: Request<{}, {}, CreateUserBody>, res: Response) => {
  const user: User = {
    id: Date.now(),
    name: req.body.name,
    email: req.body.email,
    isAdmin: false,
  };
  users.push(user);
  res.status(201).json(user);
});

router.get("/:id", (req: Request<UserParams>, res: Response) => {
  const user = users.find((u) => u.id === parseInt(req.params.id));
  if (!user) return res.status(404).json({ error: "User not found" });
  res.json(user);
});

export default router;
```

```typescript
// src/index.ts
import express from "express";
import usersRouter from "./routes/users";

const app = express();
app.use(express.json());
app.use("/users", usersRouter);

app.listen(3000, () => console.log("Listening on port 3000"));
```

### 5.5 Error handling middleware

A single typed error handler for your whole app:

```typescript
// src/middleware/errorHandler.ts
import { Request, Response, NextFunction } from "express";
import { AppError } from "../errors";

export function errorHandler(
  err: Error,
  req: Request,
  res: Response,
  next: NextFunction
): void {
  if (err instanceof AppError) {
    res.status(err.statusCode).json({
      success: false,
      message: err.message,
    });
    return;
  }

  console.error("Unexpected error:", err);
  res.status(500).json({
    success: false,
    message: "Internal server error",
  });
}
```

```typescript
// Register it last in index.ts — after all routes
app.use(errorHandler);
```

---

## 6. Databases with Prisma

Prisma is the best ORM for TypeScript. It generates types directly from your database schema, meaning your database and TypeScript types are always in sync.

### 6.1 Setup

```bash
npm install prisma @prisma/client
npx prisma init
```

This creates a `prisma/schema.prisma` file and adds `DATABASE_URL` to your `.env`.

### 6.2 Define your schema

```prisma
// prisma/schema.prisma

generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "postgresql"   // or "sqlite" for local dev
  url      = env("DATABASE_URL")
}

model User {
  id        Int      @id @default(autoincrement())
  name      String
  email     String   @unique
  role      String   @default("user")
  posts     Post[]
  createdAt DateTime @default(now())
}

model Post {
  id        Int      @id @default(autoincrement())
  title     String
  content   String?
  published Boolean  @default(false)
  author    User     @relation(fields: [authorId], references: [id])
  authorId  Int
  createdAt DateTime @default(now())
}
```

### 6.3 Run migrations

```bash
npx prisma migrate dev --name init
```

This creates the database tables AND generates TypeScript types for your models automatically.

### 6.4 Using Prisma Client

```typescript
// src/db.ts
import { PrismaClient } from "@prisma/client";

export const prisma = new PrismaClient();
```

```typescript
// src/services/userService.ts
import { prisma } from "../db";
import { User } from "@prisma/client"; // auto-generated type!

export async function getAllUsers(): Promise<User[]> {
  return prisma.user.findMany();
}

export async function getUserById(id: number): Promise<User | null> {
  return prisma.user.findUnique({ where: { id } });
}

export async function createUser(data: {
  name: string;
  email: string;
}): Promise<User> {
  return prisma.user.create({ data });
}

export async function getUserWithPosts(id: number) {
  return prisma.user.findUnique({
    where: { id },
    include: { posts: true }, // join posts automatically
  });
}

export async function deleteUser(id: number): Promise<User> {
  return prisma.user.delete({ where: { id } });
}
```

Notice you never had to write the `User` type yourself — Prisma generates it from your schema. Change the schema → run `npx prisma generate` → types update everywhere.

### 6.5 Wire it into your routes

```typescript
// src/routes/users.ts
import { Router, Request, Response, NextFunction } from "express";
import * as userService from "../services/userService";

const router = Router();

router.get("/", async (req: Request, res: Response, next: NextFunction) => {
  try {
    const users = await userService.getAllUsers();
    res.json(users);
  } catch (error) {
    next(error); // passes to your error handler middleware
  }
});

router.post("/", async (req: Request, res: Response, next: NextFunction) => {
  try {
    const user = await userService.createUser(req.body);
    res.status(201).json(user);
  } catch (error) {
    next(error);
  }
});

export default router;
```

---

## 7. CLI Tools & Scripting

TypeScript is excellent for writing CLI tools — scripts you run from the terminal to automate tasks.

### 7.1 A simple CLI script

```typescript
// src/scripts/seed.ts
import { prisma } from "../db";

async function seed() {
  console.log("Seeding database...");

  await prisma.user.createMany({
    data: [
      { name: "Alice", email: "alice@example.com" },
      { name: "Bob", email: "bob@example.com" },
    ],
    skipDuplicates: true,
  });

  console.log("Done!");
  await prisma.$disconnect();
}

seed().catch((err) => {
  console.error(err);
  process.exit(1);
});
```

Run it:
```bash
npx tsx src/scripts/seed.ts
```

### 7.2 Parsing command-line arguments

```bash
npm install commander
```

```typescript
// src/scripts/cli.ts
import { Command } from "commander";

const program = new Command();

program
  .name("myapp")
  .description("CLI for managing the app")
  .version("1.0.0");

program
  .command("create-user")
  .description("Create a new user")
  .requiredOption("-n, --name <name>", "User's name")
  .requiredOption("-e, --email <email>", "User's email")
  .action(async (options: { name: string; email: string }) => {
    console.log(`Creating user: ${options.name} (${options.email})`);
    // await userService.createUser(options);
    console.log("User created successfully");
  });

program
  .command("list-users")
  .description("List all users")
  .action(async () => {
    console.log("Fetching users...");
    // const users = await userService.getAllUsers();
    // users.forEach(u => console.log(`- ${u.name} (${u.email})`));
  });

program.parse(process.argv);
```

Run it:
```bash
npx tsx src/scripts/cli.ts create-user --name Alice --email alice@example.com
npx tsx src/scripts/cli.ts list-users
```

### 7.3 Useful scripting patterns

Reading a JSON file and processing it:

```typescript
// src/scripts/importUsers.ts
import * as fs from "fs/promises";
import * as path from "path";
import { prisma } from "../db";

interface RawUser {
  name: string;
  email: string;
}

async function importUsersFromFile(filename: string): Promise<void> {
  const filePath = path.resolve(filename);
  const content = await fs.readFile(filePath, "utf-8");
  const users: RawUser[] = JSON.parse(content);

  console.log(`Importing ${users.length} users...`);

  for (const user of users) {
    await prisma.user.upsert({
      where: { email: user.email },
      update: { name: user.name },
      create: user,
    });
    console.log(`  ✓ ${user.email}`);
  }

  console.log("Import complete");
  await prisma.$disconnect();
}

const file = process.argv[2];
if (!file) {
  console.error("Usage: tsx importUsers.ts <filename.json>");
  process.exit(1);
}

importUsersFromFile(file).catch(console.error);
```

---

## 8. What to Learn Next

You now have a solid foundation. Here's a clear progression path:

### Immediate next steps

- **Validation** — Use `zod` to validate request bodies at runtime and infer TypeScript types from your schemas automatically. This pairs perfectly with Express.
- **Authentication** — Learn JWT-based auth with the `jsonwebtoken` package. Type the JWT payload with a custom interface.
- **Testing** — Use `vitest` (faster, native TypeScript support) to write unit and integration tests for your services and routes.

### Growing your skills

- **Fastify** — A faster, more TypeScript-native alternative to Express with built-in schema validation
- **Decorators & NestJS** — NestJS is a full framework built entirely on TypeScript with a structured, Angular-like architecture. Great for larger projects.
- **Advanced Types** — Explore `Partial<T>`, `Pick<T, K>`, `Omit<T, K>`, `Record<K, V>` — TypeScript's built-in utility types that you'll use every day

### Reference resources

- **TypeScript Handbook** — typescriptlang.org/docs/handbook — the official, well-written reference
- **Prisma Docs** — prisma.io/docs — excellent, practical documentation
- **Total TypeScript** (Matt Pocock) — one of the best free+paid TypeScript learning resources online

---

*Build something real as soon as possible — even a small API with two or three routes connected to a database will teach you more than any tutorial.*
