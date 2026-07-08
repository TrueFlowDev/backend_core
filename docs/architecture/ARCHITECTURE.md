# TrueFlow Backend — Architecture Documentation

## 1. Overview

TrueFlow's backend is built as a **modular monolith** using **Hexagonal Architecture** (Ports & Adapters), combined with
elements of **Domain-Driven Design** (entities, value objects, aggregates) and **Clean/Onion layering** (domain →
application → infrastructure/presentation).

The system runs as a single deployable process, but internally it is partitioned into independent, self-contained *
*modules** (bounded contexts). Each module owns its own domain model, business rules, and persistence, and exposes
itself to the rest of the system only through explicit interfaces. This gives the project two things at once:

- **The simplicity of a monolith** — one codebase, one build, one deployment, shared infrastructure (database, cache,
  HTTP server, logger).
- **The isolation of microservices** — each module can be reasoned about, tested, and eventually extracted into its own
  service without a rewrite, because it does not reach into another module's internals.

### Purpose of this architecture

- **Isolate business rules from technical detail.** Domain logic never depends on a web framework, an ORM, or any
  third-party library. Frameworks are treated as replaceable *plugins* to the business logic, never the other way
  around.
- **Make dependencies point inward.** Outer layers (HTTP, database) depend on inner layers (domain), never the reverse.
  The domain layer has zero knowledge of how it is invoked or how it is persisted.
- **Keep modules independent.** Modules communicate through explicitly exposed ports/interfaces, not by importing each
  other's internal packages. This bounds the blast radius of change and keeps the codebase safe to grow.
- **Make the system testable without infrastructure.** Because business logic depends only on interfaces (ports), it can
  be tested with in-memory fakes, with no database, cache, or HTTP server involved.

### How responsibility is divided

At the top level, the codebase is split into three concerns:

1. **`module/*`** — the business capabilities of the product (bounded contexts). This is where all product-specific
   logic lives.
2. **`platform`** — technical infrastructure that the whole application needs to boot and run (HTTP server, database
   connection, cache connection, config, logging), independent of any business concept.
3. **`shared`** — cross-cutting building blocks that multiple modules are allowed to depend on (base repository
   behavior, cross-module ports such as transaction management and logging contracts).

Everything is wired together at the composition root (application entry point) using **dependency injection** (Uber
`fx`). No module wires itself into existence; it declares what it provides and what it needs, and the composition root
assembles the graph.

---

## 2. Project Structure — Conceptual Areas

The project is organized into the following conceptual areas. New code must always be placed according to *what it is*,
not *where it is convenient*.

### 2.1 `module/<name>` — Business Modules

**Responsibility:** Implements one bounded context / business capability end-to-end — domain rules, use cases, HTTP
entry points, and persistence for that capability.

**Belongs here:**

- Everything needed to implement one business capability: entities, value objects, use cases, ports, adapters,
  controllers, and the module's own DI wiring.

**Must not contain:**

- Generic technical infrastructure unrelated to the module's business purpose (that belongs in `platform`).
- Direct imports of another module's internal packages (domain, application, infrastructure). Cross-module interaction
  must happen through a port that the other module explicitly exposes, or by depending on `shared`.

**Allowed dependencies:**

- Its own internal layers (domain ← application ← infrastructure/presentation, see §3).
- `shared` (cross-cutting ports and base infrastructure).
- `platform` types that represent pure configuration or cross-cutting infrastructure contracts (e.g. reading
  `*config.Config` values, depending on the `Logger`/`TxManager` ports).
- Third-party libraries, but **only from the `infrastructure` and `presentation` layers** — never from `domain`.

**Forbidden dependencies:**

- Another module's `domain`, `application`, or `infrastructure` packages.
- A module reaching "sideways" into another module to reuse a use case directly — such capability must be exposed as a
  port and injected, exactly like the `notification` module exposes an SMS-sending use case that `auth` consumes
  indirectly through configuration/DI, not by importing `auth`'s guts into `notification` or vice versa.

### 2.2 `platform` — Technical Infrastructure

**Responsibility:** Boots and operates the technical foundation the application needs to run, independent of any
business rule: HTTP server, database driver, cache client, configuration loading, logging implementation, and generic
HTTP middleware (request ID, structured logging, centralized error translation).

**Belongs here:**

- Framework/driver setup (Gin engine, GORM/Postgres connection, Redis connection).
- Application-wide middleware that does not encode business rules (request ID propagation, access logging, translating
  errors into HTTP responses).
- Configuration loading and typed config structs.
- Generic, business-agnostic HTTP concerns (request validation-error formatting, Swagger registration).

**Must not contain:**

- Any business/domain rule, entity, or use case.
- Anything specific to a single module (e.g. a JWT *strategy specific to auth's login flow* belongs in the `auth`
  module's own infrastructure — `platform` only owns things every module can rely on, like the raw DB connection or the
  HTTP engine itself).

**Allowed dependencies:** third-party technical libraries (web framework, ORM, cache client, logging library) and
`shared` ports it must implement (e.g. `platform/database` implements the `TxManager` port declared in `shared`).

**Forbidden dependencies:** any `module/*` package.

### 2.3 `shared` — Cross-Cutting Contracts and Base Infrastructure

**Responsibility:** Provides the small set of ports and base implementations that *every* module is entitled to depend
on, because they represent concerns that cut across all bounded contexts (persistence access pattern, transaction
coordination, logging contract).

**Belongs here:**

- Cross-module **ports** (interfaces) such as a transaction manager contract or a logger contract, owned centrally
  because every module needs the same abstraction.
- A generic base repository type that gives modules a consistent way to obtain the correct database executor (plain
  connection vs. active transaction) without duplicating that logic in every module.
- The module's own composition wiring for the pieces above.

**Must not contain:**

- Business rules belonging to any specific module.
- Concrete infrastructure that is specific to one technology choice beyond the minimal glue needed to satisfy the shared
  port's contract (heavier technology-specific setup belongs in `platform`).

**Allowed dependencies:** `platform` (to obtain concrete engines/connections it wraps), third-party persistence
libraries strictly for the base repository glue.

**Forbidden dependencies:** any `module/*` package.

### 2.4 `pkg` — Generic, Standalone Utilities

**Responsibility:** Houses small, generic, self-contained utility code with no dependency on the application's domain or
infrastructure — the kind of code that could be copy-pasted into an unrelated project unchanged.

**Belongs here:** pure functions/algorithms with no business meaning (e.g. normalizing/validating a phone number string
format).

**Must not contain:** anything that references a domain entity, a port, a module, or the application's configuration.

**Allowed dependencies:** third-party libraries appropriate to the utility itself (e.g. a phone-number parsing library).
No dependency on any other internal package is allowed — `pkg` sits at the very bottom of the dependency graph.

### 2.5 `cmd` — Composition Root

**Responsibility:** The executable entry point. Assembles every module and the platform/shared infrastructure into a
single dependency-injection container and starts the application.

**Belongs here:** wiring of top-level modules into the DI container, application-level metadata (API documentation
annotations), and nothing else.

**Must not contain:** any business logic, any direct infrastructure setup — it only *composes* what modules and platform
already expose.

### 2.6 `migrations`

**Responsibility:** Versioned, ordered database schema changes. Each migration is an immutable, timestamped unit —
migrations are never edited after being merged, only added.

---

## 3. Module Architecture — The Pattern for New Modules

Every business module follows the same internal layering. This is the template to follow when creating a new module,
regardless of which business capability it implements.

### 3.1 Layers within a module

```
presentation/   → entry points (HTTP controllers, HTTP middleware specific to this module)
application/    → use cases (application services) that orchestrate a single business operation
domain/         → entities, value objects, and ports (interfaces) — the module's business rules
infrastructure/ → adapters that implement the module's domain ports using real technology
```

**`domain`**

- **Responsibility:** Defines *what the module knows to be true about the business* — entities, value objects,
  invariants, and the ports (interfaces) it needs from the outside world to do its job (a repository, a token provider,
  a password hasher, an ID generator, etc.).
- Entities encapsulate their own invariants: state is private, mutated only through explicit methods that enforce
  business rules, and constructed either via a validating constructor (`New...`) for brand-new instances or a restoring
  constructor (`Restore...`) for rehydrating from persistence without re-running creation-time validation.
- Value objects wrap primitive values (phone numbers, IDs, passwords, tokens) so that invalid states are
  unrepresentable — construction validates and normalizes, and once built, they are immutable.
- Ports are interfaces named after the capability they provide, from the domain's point of view (e.g. "something that
  can look up a user by phone"), not from the point of view of the technology that will implement them.
- Domain-level error values are declared next to the concept they belong to, as sentinel values, so application and
  presentation code can branch on them predictably.
- **Depends on:** nothing outside the module except `shared` ports it needs to declare against, and standard library.
  Zero third-party framework, ORM, or transport library imports.

**`application`**

- **Responsibility:** Implements *use cases* — one exported type per business operation, with an `Execute` method that
  takes a plain input struct and returns a plain output struct. A use case orchestrates domain objects and ports to
  fulfill one, and only one, business intent (e.g. "log a user in", "register a user", "save a profile").
- Use cases depend only on domain ports (interfaces), constructed via dependency injection, never on concrete
  infrastructure types.
- Use cases contain orchestration logic (call this port, then that port, then construct this entity) — they do not
  reimplement business rules that belong on the entity/value object itself.
- **Depends on:** its own module's `domain`. May depend on `platform` for pure configuration values (e.g. a token TTL
  from config). Does not depend on `infrastructure` or `presentation`.

**`infrastructure`**

- **Responsibility:** Implements the module's domain ports against real technology — database repositories, external API
  clients, ID generators, password hashers, token providers.
- Each adapter type is paired with a compile-time assertion that it satisfies the corresponding domain port (
  `var _ port.X = (*Impl)(nil)`), so a broken contract fails at compile time, not at runtime.
- Where persistence is involved, the adapter translates between the domain entity and a generated/persistence-specific
  model through a dedicated **mapper**, so the domain entity's shape never leaks into the storage schema and vice versa.
- Code-generated persistence access (query builders/DAOs) lives in its own sub-package, generated from the database
  schema through a dedicated, build-excluded generator entry point — it is treated as derived, replaceable code, never
  hand-edited business logic.
- **Depends on:** its own module's `domain`, `shared` (base repository, transaction/logger ports), `platform` (config,
  raw driver types it needs to construct a client), third-party technical libraries.

**`presentation`**

- **Responsibility:** Exposes the module's use cases to the outside world — HTTP controllers (and any module-specific
  middleware, such as an authentication guard that depends on this module's own ports).
- A controller's only responsibilities are: parse and validate the transport-level request, translate it into the use
  case's input type, invoke exactly one use case, and translate the result (success or domain error) into a
  transport-level response. A controller must not contain business rules, must not talk to a repository or any adapter
  directly, and must not know about persistence models.
- Route registration is a separate, explicit function invoked at composition time, kept apart from the controller's
  business method so that "what routes exist" is visible in one place per module.
- **Depends on:** its own module's `application` (use cases) and `domain` (to reference sentinel errors for translation,
  when needed), plus `platform`'s generic HTTP building blocks (validation helpers, error types).

**`module.go` (module wiring)**

- Each module exposes its own dependency-injection module, declaring every constructor it provides (adapters bound to
  the ports they satisfy, use cases, controllers) and every route-registration function it invokes at startup. This file
  is the only place inside a module allowed to know about the DI framework's wiring mechanics beyond individual
  constructors.

### 3.2 Request flow

A typical inbound request flows strictly in one direction, through the layers described above:

```
HTTP request
  → presentation/http/controller   (bind + validate transport input)
  → application/usecase            (orchestrate one business operation)
  → domain (entity / value object) (enforce business invariants)
  → domain/port                    (abstract capability needed, e.g. repository)
  → infrastructure/adapter         (fulfills the port against real tech, e.g. Postgres via GORM)
  → database / external system
```

The response follows the same path in reverse: the adapter returns a domain entity or a domain-level error, the use case
maps it into an output struct or propagates the error, and the controller renders it as an HTTP response. At no point
does a layer skip over the one directly beneath it, and at no point does an inner layer (domain, application) import a
concrete implementation from an outer layer (infrastructure, presentation) — only the reverse.

Errors are propagated as values through every layer using a single structured error type; the outermost HTTP layer is
the only place where an error is finally translated into a transport-level status code and response body, based on the
error's own declared metadata rather than a per-controller `switch` statement.

---

## 4. Layer Rules Summary

| Layer            | May contain                                                                        | Must not contain                                                       | Depends on                                                     |
|------------------|------------------------------------------------------------------------------------|------------------------------------------------------------------------|----------------------------------------------------------------|
| `domain`         | Entities, value objects, ports, domain error sentinels                             | Framework code, SQL/ORM, HTTP concepts, third-party business libraries | Nothing but standard library and `shared` ports                |
| `application`    | Use cases orchestrating domain + ports                                             | Direct infrastructure calls, HTTP concepts, SQL                        | `domain` (own module), `platform` config values                |
| `infrastructure` | Adapters implementing domain ports, mappers, generated persistence code            | Business rules/invariants, transport concerns                          | `domain` (own module), `shared`, `platform`, third-party libs  |
| `presentation`   | Controllers, request/response DTOs, module-specific middleware, route registration | Business rules, direct persistence/adapter access                      | `application`, `domain` (errors only), `platform` HTTP helpers |

The single governing rule across all of the above: **dependencies point inward, never outward.** `domain` knows nothing
about anyone. `application` knows only `domain`. `infrastructure` and `presentation` know `application` and `domain`,
but neither knows the other.

---

## 5. Infrastructure Folders — Philosophy

**`platform`**
Exists so that "how the process runs" is completely decoupled from "what the business does." If the web framework, ORM,
or cache client is ever replaced, the change is contained to `platform` (plus each module's `infrastructure`, which
depends on `platform`'s concrete types) — no module's `domain` or `application` code is touched. Anything that answers "
what technology powers this application" belongs here. Anything that answers "what does the business need" does not.

**`shared`**
Exists to hold the *minimum* set of abstractions that must be common across all modules because duplicating them
per-module would either be impossible (a single transaction must span repositories from potentially different modules)
or wasteful (every module needs the same "give me the right DB executor" logic). `shared` is intentionally kept small: a
growing `shared` package is a warning sign that module boundaries are being used as a dumping ground instead of being
respected. If something is only used by one module, it belongs inside that module, not in `shared`.

**`pkg`**
Exists for code that has no opinion about this application at all — it would be exactly as useful in a completely
different project. This is the test for whether something belongs in `pkg`: if it needs to know about a domain entity, a
config struct, or a port, it does not belong here.

**`internal`**
The entire application lives under `internal` by Go convention, guaranteeing that nothing in this codebase can be
imported by another module outside this repository. This is the outermost enforcement of "this is a single,
self-contained deployable," reinforcing the modular-monolith boundary at the language level.

---

## 6. Development & Maintenance Rules

**Adding a new feature to an existing module**

1. Express the business rule on the entity or value object in `domain` first — if the rule is an invariant, it belongs
   there, not in the use case.
2. If the feature needs a new capability from the outside world (persistence, an external service), define a new port in
   `domain/port` describing that capability abstractly.
3. Implement the use case in `application`, orchestrating existing or new ports.
4. Implement the port in `infrastructure/adapter`, with a compile-time interface assertion, and a mapper if persistence
   is involved.
5. Expose the use case through a controller in `presentation`, registered explicitly in the module's route registration
   function.
6. Wire any new constructor into the module's `module.go`, annotating adapters with the port they satisfy.

**Creating a new module**
Follow the exact layering in §3.1 (`domain` → `application` → `infrastructure`/`presentation`), give the module its own
`module.go` exposing an `fx.Module`, and register that module at the composition root (`cmd/server`). A new module must
not import another module's internal packages; if it needs another module's capability, that capability must be exposed
as a port and satisfied by an adapter, the same way cross-cutting concerns like transactions and logging are exposed
through `shared`.

**Where business logic goes**
Always in `domain` (invariants owned by entities/value objects) or `application` (orchestration across ports). Never in
a controller, never in a repository adapter, never in `platform`.

**Where database code goes**
Only inside a module's `infrastructure` layer, behind a port defined in that module's `domain`. Generated query-builder
code is isolated in its own sub-package and regenerated from the schema, never hand-modified. Raw driver/connection
setup lives in `platform`; how a specific module *queries* the database lives in that module's `infrastructure`.

**Dependency management / avoiding coupling**

- A module depends on another module only through a port, never through a direct import of its internals.
- `domain` packages never import a third-party library that isn't standard-library-equivalent in spirit (i.e., they
  don't import web frameworks, ORMs, or SDKs).
- All object construction happens through dependency injection at the composition boundary (`module.go` files and
  `cmd/server`) — business code never instantiates its own infrastructure dependencies.
- Interfaces are declared by the consumer (`domain/port`), not by the implementer — the implementer (
  `infrastructure/adapter`) depends on the interface, not the other way around. This is what keeps the dependency arrow
  pointing inward.

---

## 7. Data Flow Summary

A standard write or read request moves through the system as:

```
Controller (presentation)
   → Use Case (application)
      → Entity / Value Object validation (domain)
      → Port invocation (domain interface)
         → Adapter (infrastructure) — implements the port
            → Mapper (entity ⇄ persistence model)
            → Generated query layer / external client
            → Database or external system
```

Cross-cutting concerns are threaded through this flow rather than bolted onto individual layers:

- **Transactionality** is obtained by asking a shared transaction port to run a function within a transaction; that
  function receives a context carrying the active transaction, and any repository built on the shared base-repository
  type automatically picks up that transaction from the context instead of using the default connection. This keeps
  transaction *boundaries* a use-case-level decision, while keeping transaction *plumbing* out of both the use case and
  the repository's business code.
- **Error propagation** uses one structured error type throughout every layer, carrying an internal diagnostic payload
  for logging and a client-safe payload for the API response, so the same error can be logged with full context and
  rendered to the client with only what it should see.

---

## 8. Design Patterns in Use

**Hexagonal Architecture (Ports & Adapters)**
*Why:* Keeps business logic independent of frameworks and infrastructure so either can change without the other
noticing.
*Problem solved:* Prevents technical churn (swapping a database driver, a web framework, a cache) from cascading into
business logic.
*Rule:* A port is defined by the layer that needs the capability (`domain`); it is implemented by the layer that has the
technology to provide it (`infrastructure`).

**Modular Monolith**
*Why:* Gets most of the boundary discipline of microservices (independent modules, explicit contracts) while keeping one
deployable, one transaction boundary, and much lower operational overhead.
*Problem solved:* Avoids both the "big ball of mud" failure mode of an undisciplined monolith and the premature
operational complexity of distributed services before the product needs it.
*Rule:* Modules never import each other's internal layers; cross-module capability sharing goes through explicitly
declared ports.

**Repository Pattern with a shared executor-resolution base**
*Why:* Gives every repository a single, consistent way to run queries against either the default connection or an active
transaction, without each repository re-implementing that decision.
*Problem solved:* Removes duplicated transaction-detection logic from every repository, and keeps repositories agnostic
to whether they're currently inside a transaction.
*Rule:* All repositories embed the shared base and resolve their executor through it rather than holding their own raw
connection.

**Unit of Work (transaction-scoped context propagation)**
*Why:* Lets a use case span multiple repository calls, potentially across the same module, and commit or roll them back
atomically.
*Problem solved:* Keeps transaction boundaries a business/use-case-level decision, not something a repository or
controller decides on its own.
*Rule:* Only `application`-layer code (use cases) decides transaction boundaries. Repositories never begin or commit a
transaction themselves — they only consult the context for one that may already be active.

**Mapper Pattern (entity ⇄ persistence model translation)**
*Why:* Keeps the domain entity's shape independent of the database schema/ORM model, so either can evolve without
forcing a change in the other.
*Problem solved:* Prevents ORM annotations, nullable database quirks, and schema-specific representations from leaking
into domain objects.
*Rule:* Conversion between an entity and its persistence model always goes through an explicit mapper function; nothing
else performs this translation inline.

**Value Objects**
*Why:* Makes invalid data structurally unrepresentable — once constructed, a value object is guaranteed valid.
*Problem solved:* Eliminates scattered, duplicated validation of primitives (phone numbers, IDs, passwords) throughout
the codebase.
*Rule:* A value object's only constructor path validates/normalizes its input; once built, it exposes no way to become
invalid.

**Structured, Layered Errors (diagnostic vs. client-facing payload)**
*Why:* One error value needs to serve two very different audiences: an operator who needs to debug it, and an API client
who needs a safe, structured response.
*Problem solved:* Prevents leaking internal detail to clients while preventing loss of debugging context internally;
avoids ad-hoc error-to-HTTP-status mapping scattered across controllers.
*Rule:* Domain and infrastructure code produces/wraps errors using the shared structured error type with the appropriate
code and metadata; only the centralized HTTP error-handling middleware decides how an error is finally rendered to the
client.

**Dependency Injection via a composition-root container**
*Why:* Keeps every layer ignorant of how its dependencies are constructed, which is what makes the
inward-pointing-dependency rule enforceable in practice.
*Problem solved:* Avoids global state, manual wiring duplication, and hidden dependencies between layers/modules.
*Rule:* Constructors declare exactly what they need as parameters; nothing reaches for a dependency through a global or
a service locator. Wiring happens exclusively in `module.go` files and the composition root.

**Code Generation for Persistence Access**
*Why:* Keeps repetitive, schema-derived query code out of hand-written business code, generated straight from the
database schema.
*Problem solved:* Removes an entire class of hand-written query bugs and keeps the query layer in sync with the schema
by construction.
*Rule:* Generated code is never hand-edited; the generator itself is excluded from normal builds and re-run explicitly
when the schema changes.

---

## 9. Architecture Decision Records

**ADR: Business logic must never live in a controller.**
Controllers are transport adapters. Putting business rules there would couple business logic to the HTTP framework and
to the shape of an HTTP request, making the rule untestable without spinning up a web server and impossible to reuse
from another entry point (a queue consumer, a CLI, gRPC). Business rules live in `domain`/`application`, which know
nothing about HTTP.

**ADR: Transaction management belongs at the use-case (application) level.**
A single business operation may need to touch more than one repository atomically. Only the use case knows the true
transactional boundary of a business operation; a repository only knows about its own table, and a controller only knows
about the transport request. Placing transaction control anywhere else would either be too narrow (a repository can't
coordinate across other repositories) or too broad/leaky (a controller managing transactions would blur transport and
business concerns).

**ADR: Dependencies always point inward — infrastructure and presentation depend on domain, never the reverse.**
This is what makes business rules independently testable and technology-agnostic. If `domain` depended on
`infrastructure`, replacing a database or web framework would force changes to business rules; instead, only
`infrastructure`/`presentation` change, because they are the ones depending on `domain`'s ports.

**ADR: Modules communicate only through explicitly exposed ports, never by importing each other's internals.**
This is what preserves the option of extracting a module into its own service later, and what keeps a change in one
business capability from silently breaking another. It is enforced structurally by keeping each module's `domain`,
`application`, and `infrastructure` unimported by any other module.

**ADR: Interfaces (ports) are owned by the consumer, not the implementer.**
Ports live in `domain/port` — the layer that needs the capability — and are implemented in `infrastructure/adapter`.
This ensures the domain layer only ever depends on abstractions it defines itself, keeping the dependency arrow inward
even when the concrete implementation changes entirely.

**ADR: A single structured error type is used across all layers.**
Using one error type end-to-end (rather than ad-hoc `error` values reinterpreted per layer) allows every layer to attach
diagnostic context as an error crosses it, while guaranteeing the outermost layer has everything it needs to produce
both a safe client response and a fully-detailed internal log entry from the exact same value.

---

## 10. Guide for New Developers

**How this project thinks**

TrueFlow's backend is organized around one question at every step: *"does this code express a business rule, or does it
operate a piece of technology?"* Business rules go inward (`domain`, `application`); technology operation goes outward (
`infrastructure`, `presentation`, `platform`). Everything else in this document is a consequence of consistently
answering that question the same way.

Each business capability is a self-contained module that could, in principle, be lifted out into its own service —
because it never reaches into another module's internals, only into ports that module or `shared` explicitly exposes.

**Path to adding a new feature**

1. Identify which module the feature belongs to (or whether it needs a brand-new module).
2. Model the business rule on an entity or value object in that module's `domain`.
3. If new external capability is needed, declare a port for it in `domain/port`.
4. Write a use case in `application` that orchestrates the operation.
5. Implement the port in `infrastructure/adapter`, with a mapper if persistence is involved.
6. Expose the use case via a controller in `presentation`, and register its route explicitly.
7. Wire any new constructors in the module's `module.go`.

**Mistakes to avoid**

- Do not put business rules in a controller — controllers translate transport ↔ use case input/output and nothing else.
- Do not call a repository or adapter directly from a controller or use case — always go through a port.
- Do not import another module's `domain`, `application`, or `infrastructure` package directly — expose a port instead.
- Do not let `domain` import a web framework, ORM, or any third-party technology library.
- Do not manage transactions inside a repository — transaction boundaries belong to the use case that orchestrates the
  operation.
- Do not hand-edit generated persistence code — change the schema/generator and regenerate.
- Do not grow `shared` with anything used by only one module — keep it reserved for genuinely cross-cutting concerns.
- Do not construct dependencies manually inside business code — declare them as constructor parameters and let the
  composition root inject them.
