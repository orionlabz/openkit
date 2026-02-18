# INTERNAL_BASELINE

## Snapshot Date

- 2026-02-18

## Internal Scope Reality

- `internal/**/*.go`: 0 files
- `internal/` currently acts as template/content source (`internal/templates/**`), not active runtime implementation.

## Key Findings

### F1: Broken template references in mobile prompt template

- File: `internal/templates/base/prompts/mobile-developer.md`
- Issue: references to `../skills/.internal/...` paths that do not exist.
- Impact: generated prompt docs can point to invalid skill paths.

### F2: Legacy stack hints in stack-selection skill template

- File: `internal/templates/base/skills/stack-selection/SKILL.md`
- Issue: references include `backend/go.mod` and `internal/templates/base/` as required-style guidance.
- Impact: biases generated guidance toward outdated/irrelevant checks for current runtime direction.

### F3: Legacy memory command surface still present in semantic-memory rule template

- File: `internal/templates/memory/rules/SEMANTIC_MEMORY.md`
- Issue: command examples include unsupported commands like:
  - `openkit memory list`
  - `openkit memory search`
  - `openkit memory stats`
  - `openkit memory debug`
  - `openkit memory export`
- Impact: generated artifacts can instruct non-existent CLI behavior.

## Risk Ranking

| Priority | Item | Risk |
|---|---|---|
| P0 | Fix invalid `.internal` skill paths in mobile prompt template | Broken guidance in generated agent prompts |
| P0 | Replace legacy memory command examples in semantic-memory rule template | User confusion and invalid command execution |
| P1 | Update stack-selection template to runtime-neutral checks | Reduces stale assumptions in generated docs |

## Phase 1/2 Status (Executed)

- `internal/templates/base/prompts/mobile-developer.md`: fixed broken `.internal` paths.
- `internal/templates/base/prompts/backend-specialist.md`: replaced `go.mod` hint with `Cargo.toml` support.
- `internal/templates/base/skills/stack-selection/SKILL.md`: removed stale `backend/go.mod` and legacy internal reference guidance.
- `internal/templates/base/prompts/orchestrator.md`: replaced plugin-centric memory section with docs-first Memory Kernel workflow.
- `internal/templates/memory/rules/SEMANTIC_MEMORY.md`: rewritten as legacy notice + migration boundary.

## Phase 3 Status (Executed)

- Deprecated semantic-memory plugin implementation files removed from `internal/templates/memory/`:
  - `index.ts`, `package.json`
  - `lib/embeddings.ts`, `lib/memory.ts`, `lib/storage.ts`
  - `scripts/bridge.ts`, `scripts/verify_optimization.ts`
  - `rules/SEMANTIC_MEMORY.md`
- Replacement resources added for docs-first Memory Kernel:
  - `internal/templates/memory/config.yaml`
  - `internal/templates/memory/derivation.yaml`
  - `internal/templates/memory/queue.yaml`
  - `internal/templates/memory/rules/MEMORY_KERNEL.md`
  - `internal/templates/memory/README.md`

Post-fix validation snapshot (`internal/templates/**`):

- `.internal` broken path references: 0
- unsupported `openkit memory` subcommands (`list/search/stats/prune/export/config/debug`): 0
- stale Go-runtime prompt references (`go.mod`, `Go CLI`, `internal/cli/*`): 0
- plugin-only memory tool references in `internal/templates/memory/**`: 0

## Related

- [[audit/HUB-AUDIT.md]]
- [[audit/SWEEP_PLAN.md]]
- [[audit/BASELINE_AUDIT.md]]
