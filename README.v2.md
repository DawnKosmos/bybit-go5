# bybit-go5 v2 Plan

This document outlines the plan for v2: a reproducible code-generation pipeline from Bybit Docs v5 to Go code, preserving your existing client design patterns.

## Goals
- __Single source of truth__: autogenerate from `libs/bybit-docs/docs/v5`.
- __Preserve design__: same request/response modeling approach as v1.
  - GET requests use `url` tags for query encoding.
  - POST requests use `json` tags for bodies.
  - Optional fields use `omitempty` and pointer types where needed.
- __Coverage__: REST and WebSocket (public/private) endpoints.
- __Repeatable__: one command to update docs + regenerate code.
- __Quality__: typed enums, validation, tests with sample payloads.

## Architecture
- __Parser__: read Markdown in `libs/bybit-docs/docs/v5/**` to extract endpoints, params, enums, responses.
- __IR (Intermediate Representation)__: normalized schema of endpoints and types.
- __Generators__: Go code via templates for models, REST client methods, and WS topics/messages.
- __Overrides__: YAML/JSON patches to correct doc inconsistencies without hand-editing generated code.

## Repo Layout Additions
- `cmd/bybitgen/` — codegen CLI.
- `internal/docparse/` — Markdown parser utilities.
- `internal/ir/` — IR definitions and normalization.
- `internal/gen/{models,rest,ws}/` — text/template generators.
- `schemas/overrides/*.yaml` — manual corrections and enums.
- `testdata/bybit/` — sample JSON fixtures from docs.

Generated output will continue to populate:
- `models/` — request/response structs, preserving `url` vs `json` tags and `omitempty`.
- `client.go` — REST methods with `context.Context` support.
- `ws/` — topics, subscribe/unsubscribe payloads, router/handlers.

## Design Principles (kept from v1)
- __Tag usage__:
  - GET requests: struct fields tagged with ``url:"..."`` for query params.
  - POST requests: struct fields tagged with ``json:"..."`` for body payloads.
- __Optionality__: optional fields include `,omitempty`; where tri-state is needed, use pointer types.
- __Errors__: unify error handling using response `retCode/retMsg` (via an interface similar to `models.ReturnCode`).
- __Types__: favor strong types and enums; use decimal-safe types for prices/qty.

## Makefile Targets
- `make submodule-update` — fetch latest Bybit docs into `libs/bybit-docs`.
- `make generate` — run the code generator to update `models/`, `client.go`, and `ws/`.
- `make check` — `go vet`, lint (optional), and `go test ./...`.
- `make regen` — update docs + generate + test.

Example flow:
```bash
make submodule-update
make generate
make check
# commit and push
```

## Codegen Flow
1. __Parse__: scan `libs/bybit-docs/docs/v5/**` for endpoint pages.
2. __Extract__: endpoint path, method, tag/category; params (name, type, enum, required); response schemas; WS topics.
3. __Normalize__: map primitives to Go types; handle `string|number`, timestamps, decimals; apply `schemas/overrides`.
4. __Generate__:
   - Requests: types with `Validate()` and `ToQuery()` helpers where useful.
   - Responses: structs with `json` tags; helpers like `ToCandle()` retained.
   - Client methods: `GetKline(ctx, req) (Resp, error)` pattern; retries configurable.
   - WS: topic constants, subscription payloads, message dispatch helpers.
5. __Verify__: golden tests decode sample JSON; ensure backward-compatible helpers.

## Migration
- Module path: release as `v2` (e.g., `github.com/DawnKosmos/bybit-go5/v2`).
- Keep core API shapes; deprecate renamed items with shims where trivial.
- Document breaking changes and provide mapping tips.

## Open Decisions
- Decimal representation: `shopspring/decimal` vs `string`. Default proposal: `shopspring/decimal` for numeric precision.
- Transport abstraction: interface for HTTP client with retry/policies.
- Context-first: ensure all public calls accept `context.Context`.

## Next Steps
- Scaffold `cmd/bybitgen` and minimal pipeline.
- Implement parser + generation for one group (Market/Kline) end-to-end.
- Add overrides + golden tests.
- Expand coverage iteratively across REST/WS.
