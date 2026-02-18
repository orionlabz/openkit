# SWEEP_PLAN

## Goal

Achieve maximum consistency between current OpenKit architecture (Rust-only runtime + docs-first permanent memory) and all internal documentation assets.

## Principles

- Keep only what is true for the current runtime.
- Prefer one canonical source per subject.
- Archive historical context; do not mix history with operational docs.
- Remove references that imply unsupported commands/paths.

## Workstreams

### W1: Inventory and classification

- INPUT: `docs/**`, `.opencode/**`, `rust-cli/src/main.rs`.
- OUTPUT: File-by-file status (`keep`, `rewrite`, `archive`, `remove`).
- VERIFY: Every markdown file belongs to exactly one status.

### W2: Runtime alignment

- INPUT: `rust-cli/src/main.rs`, `README.md`, `docs/API.md`.
- OUTPUT: Canonical command surface map.
- VERIFY: Every documented command exists in runtime; no unsupported command examples.

### W3: Legacy reference cleanup

- INPUT: references to `internal/`, `cmd/openkit`, Go-specific tooling, legacy memory plugin flows.
- OUTPUT: stale references removed or moved to archive docs.
- VERIFY: no stale references remain in active docs.

### W4: Skills/prompts/commands/rules alignment

- INPUT: `.opencode/skills/**`, `.opencode/prompts/**`, `.opencode/commands/**`, `.opencode/rules/**`.
- OUTPUT: corrected paths and instructions aligned with current repo layout.
- VERIFY: no broken local references and no invalid execution paths.

### W5: Documentation graph integrity

- INPUT: all docs hubs and feature/sprint docs.
- OUTPUT: complete wikilinks + `## Related` sections for active docs.
- VERIFY: Obsidian link lint passes with zero broken links.

### W6: Governance and quality gates

- INPUT: CI workflows + local verification scripts.
- OUTPUT: repeatable audit checklist for each release.
- VERIFY: checklist executable in CI/local without manual interpretation.

## Execution Phases

1. Baseline snapshot
2. Classification matrix
3. High-impact rewrites (API, quality gates, backend/security truth)
4. Skills/prompts path fixes
5. Archive/remove pass
6. Final validation and report

## Stop Criteria

- Zero broken links (`obsidian-link-lint`).
- Zero stale command/path references in active docs.
- Runtime/API docs fully aligned with `rust-cli/src/main.rs`.
- One source of truth per critical topic.

## Related

- [[audit/HUB-AUDIT.md]]
- [[audit/BASELINE_AUDIT.md]]
- [[audit/DECISION_MATRIX.md]]
- [[API.md]]
