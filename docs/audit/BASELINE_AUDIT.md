# BASELINE_AUDIT

## Snapshot Date

- 2026-02-18

## Scope Size

- `docs/**/*.md`: 146 files
- `.opencode/skills/**/*.md`: 108 files
- `.opencode/prompts/*.md`: 18 files
- `.opencode/commands/**/*.md`: 8 files
- `.opencode/rules/**/*.md`: 9 files

## Graph Health

- Obsidian link lint (docs): 0 broken links

## Drift Signals

### docs

- Legacy Go/runtime references: 50 files, 208 hits
- Legacy memory command/flag references (`openkit memory list/search/stats/...`, `--memory`): 17 files, 141 hits
- Minisign references in active docs: 1 file, 1 hit

### .opencode

- Legacy Go/runtime references: 4 files, 17 hits
- Legacy memory command references: 0 files
- Minisign references: 0 files

## High-Impact Inconsistencies (sample)

- `docs/BACKEND.md` points to removed Go paths (`internal/cli/*`, `internal/syncer/*`).
- `docs/QUALITY_GATES.md` still describes Go tooling (`golangci-lint`, Go-based CI/release notes).
- `docs/GLOSSARY.md` still defines OpenKit as a Go CLI.
- `docs/sprint/Sprint-05/**` contains detailed legacy plugin command docs not aligned with current runtime.
- `.opencode/prompts/mobile-developer.md` uses broken `.internal` skill paths.
- `.opencode/skills/stack-selection/SKILL.md` references `internal/templates/base/` and `backend/go.mod` checks.

## Immediate Priority Queue

1. Rewrite `docs/BACKEND.md` to Rust runtime reality.
2. Rewrite `docs/QUALITY_GATES.md` to Rust CI and current scripts.
3. Clean glossary/security/action-items references to removed Go internals.
4. Fix `.opencode/prompts/mobile-developer.md` path references.
5. Sanitize `.opencode/skills/stack-selection/SKILL.md` legacy repo assumptions.
6. Classify and archive legacy Sprint-05 memory-plugin docs.

## Related

- [[audit/HUB-AUDIT.md]]
- [[audit/SWEEP_PLAN.md]]
- [[audit/DECISION_MATRIX.md]]
- [[HUB-DOCS.md]]
