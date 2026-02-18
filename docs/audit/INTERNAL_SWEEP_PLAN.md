# INTERNAL_SWEEP_PLAN

## Goal

Sanitize `internal/templates/**` so generated assets are fully consistent with current OpenKit CLI surface and architecture.

## Scope

- `internal/templates/base/prompts/**`
- `internal/templates/base/skills/**`
- `internal/templates/memory/**`

## Constraints

- Keep historical context only where explicitly marked as legacy.
- Active templates must never reference unsupported commands.
- Template paths must resolve inside generated package structure.

## Phases

### Phase 1: Template path integrity

Status: completed

- Fix invalid path prefixes (notably `.internal` path segments).
- Verify all relative links in templates resolve.

VERIFY:
- Zero unresolved local links in `internal/templates/**` markdown.

### Phase 2: Command surface normalization

Status: completed

- Replace legacy memory command examples with supported runtime set:
  - `openkit memory init`
  - `openkit memory doctor`
  - `openkit memory capture`
  - `openkit memory review`
- Remove `--memory` flag guidance from active templates.

VERIFY:
- No occurrences of unsupported command names in active templates.

### Phase 3: Runtime-era assumption cleanup

Status: completed

- Remove stale assumptions tied to removed internals or old stack examples.
- Keep stack checks generic and language-agnostic where possible.

VERIFY:
- No required guidance references removed runtime internals.

### Phase 4: Sync parity pass

Status: in progress

- Ensure equivalent content in `.opencode` and `internal/templates/base` stays aligned (or define explicit source-of-truth + generation direction).

VERIFY:
- Documented ownership matrix for duplicated artifacts.

### Phase 5: Regression gates

Status: pending

- Add grep-based checks in CI/local checklist to block reintroduction of legacy commands in active templates.

VERIFY:
- CI/local audit check fails on banned patterns and passes on clean state.

## Deliverables

- `INTERNAL_BASELINE.md` updated to post-fix state.
- Template patch set for prompts/skills/memory rules.
- Pattern-based guard list for recurring drift.

## Related

- [[audit/HUB-AUDIT.md]]
- [[audit/INTERNAL_BASELINE.md]]
- [[audit/DECISION_MATRIX.md]]
