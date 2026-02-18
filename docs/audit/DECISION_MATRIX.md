# DECISION_MATRIX

## Decision Rules

- `keep`: accurate and required by current runtime/process.
- `rewrite`: topic is needed, but content is stale or conflicting.
- `archive`: historical artifact that no longer guides active execution.
- `remove`: redundant/no value and not required for traceability.

## Initial Matrix (Phase 1)

| Area | Target | Decision | Reason |
|---|---|---|---|
| Runtime architecture | `docs/BACKEND.md` | rewrite | References removed Go internal paths. |
| Quality process | `docs/QUALITY_GATES.md` | rewrite | Commands and CI references are pre-Rust cutover. |
| Terminology | `docs/GLOSSARY.md` | rewrite | Core product definition still says Go CLI. |
| Security evidence | `docs/SECURITY.md` | rewrite | Mixed evidence from removed code paths. |
| Old memory plugin sprint | `docs/sprint/Sprint-05/*` | archive | Valuable historical context, not active guidance. |
| Rust migration closure | `docs/sprint/Sprint-08/*`, `docs/sprint/Sprint-09/*` | keep | Relevant migration traceability and parity outcomes. |
| Legacy migration guide | `docs/MEMORY_LEGACY_MIGRATION.md` | rewrite | Keep only currently valid migration statements. |
| Prompt link integrity | `.opencode/prompts/mobile-developer.md` | rewrite | Uses broken `.internal` skill references. |
| Stack selection assumptions | `.opencode/skills/stack-selection/SKILL.md` | rewrite | Includes legacy OpenKit-internal Go references. |
| Internal memory templates | `internal/templates/memory/*` | rewrite/remove | Replace deprecated semantic-memory plugin files with docs-first Memory Kernel resources. |
| Canonical docs hub | `docs/HUB-DOCS.md` | keep | Valid index; now linked to audit hub. |

## Phase 2 Deliverables

- Updated `docs/BACKEND.md`, `docs/QUALITY_GATES.md`, `docs/GLOSSARY.md`.
- Fixed `.opencode/prompts/mobile-developer.md` references.
- Updated `.opencode/skills/stack-selection/SKILL.md` examples.
- Archival proposal for Sprint-05 plugin-era docs.

## Verification Checklist

- [ ] `obsidian-link-lint` reports zero broken links.
- [ ] No active doc references removed runtime paths (`internal/*`, `cmd/openkit`).
- [ ] No active doc references unsupported command surface.
- [ ] Skills/prompts reference existing paths only.

## Related

- [[audit/HUB-AUDIT.md]]
- [[audit/SWEEP_PLAN.md]]
- [[audit/BASELINE_AUDIT.md]]
- [[HUB-DOCS.md]]
