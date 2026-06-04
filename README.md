# Introduction to Design Patterns with Go — Source Code

*For further enquiry please contact Akin Kaldiroglu at akin@kaldiroglu.dev*

**Created:** 2026-06-03
**Project:** Companion source code for the course *Introduction to Design Patterns with Go* — the Go port of the Java edition. It backs the slides with runnable examples.

## Expected benefits

- A small, runnable codebase students can read alongside the video.
- The course examples in real, idiomatic Go:
  1. **Three "before" smells** — a nested `if` access check (`complexifs`), a factory-heavy object graph (`complexobject`), and telescoping constructors (`db`).
  2. **The Proxy pattern** — the Citizen / Prime Minister exercise, in three refactoring stages (the *after* picture).

## Functional properties

This is a faithful, idiomatic-Go port of the author's Java repo. It covers **all four**
examples from the Java source — three "problem" (before) examples and the Proxy pattern in
three stages.

- **`problems/complexifs`** — `User` (many fields) + `UserProcessor.ProcessUser`, a deeply nested `if` tree returning access-code strings (e.g. `FULL_ACCESS_GRANTED`). Kept deliberately as the Big Ball of Mud the course teaches you to refactor.
- **`problems/complexobject`** — a `UserService` that assembles a `User` graph from a `NewUserRequest` through several factories. The factories are deliberately left as **stubs that return `nil`** — the "complex object construction" smell that motivates the **Factory / Builder** patterns.
- **`problems/db`** — `DatabaseConnection`, demonstrating the **telescoping constructor** smell that motivates the **Builder** pattern. `db.Run()` mirrors the Java `db/Main`'s four connection profiles.
- **`proxy/pm`** — the Gang of Four **Proxy** pattern, in three stages (`pm1`/`pm2`/`pm3`):
  - **pm1 (Solution 1):** `Citizen` → concrete `PM` (low cohesion — the PM sorts out *and* resolves).
  - **pm2 (Solution 2):** `Citizen` → `Proxy` → `PM` (the proxy sorts out and delegates).
  - **pm3 (Solution 3):** `PM` is an **interface** (the *Subject*); `RealPM` and `ProxyPM` implement it; `PMSecretary` is a **factory** (`getMePM()`) serving a `ProxyPM` typed as `PM`; `Citizen` depends only on `PM`.

## Architectural approach

- Idiomatic Go: structs with **exported fields** (Go doesn't use getters/setters), `New…` constructor functions, an **interface** for the Solution-3 `PM`, methods with PascalCase exported names and lowercase unexported helpers (`sortOut`, `resolve`, `delegate`).
- Module path **`dev.kaldiroglu/dp/intro`** honours the house root namespace; package paths mirror the Java packages (`proxy/pm/pm1`, `problems/complexifs`, …).

Two Go-specific adaptations (noted in the code):

- Go forbids the **import cycle** that Java tolerates between `complexobject` and its `user` subpackage, so the stub `UserFactory.CreateUserFromRequest` takes the request as `any`.
- Go has no nil `string`, so the Java nullable-`String` checks (`SuspensionReason`, `userType`, `department`) become `!= ""` — empty string stands in for Java's `null`.

## Layout

```
IntroductionToDesignPatternsWithGo/
├── go.mod                               (module dev.kaldiroglu/dp/intro)
├── main.go                              (runs the pm1/pm2/pm3 + db demos)
├── problems/
│   ├── complexifs/   { user.go, userprocessor.go }
│   ├── complexobject/ { newuserrequest.go, userservice.go,
│   │                    user/ { entities.go, factories.go } }
│   └── db/           { db.go }
└── proxy/pm/         { pm1/pm1.go, pm2/pm2.go, pm3/pm3.go }
```

> The `complexobject` example has no `Main` in the Java repo, so (like the original) it is library-only — compiled by `go build ./...` but not run by `main`.

This is a read-along companion to the course: there is no test package (the Java original has none either). The behaviour is shown by running the demos.

## Run it with

```sh
# from the repo root
go run .

# compile every package (incl. the library-only examples):
go build ./...
```
