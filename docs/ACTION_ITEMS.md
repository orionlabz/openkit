# ACTION_ITEMS

| Priority | Item | Impact | Effort | Owner | Notes |
|---|---|---|---|---|---|
| P0 | Complete docs/internal sweep execution | High | Medium | Docs/Platform | Execute `docs/audit/SWEEP_PLAN.md` and close matrix actions. |
| P0 | Add secret scanning in CI | High | Low | Security | Missing in `.github/workflows/ci.yml`. |
| P1 | Add Rust dependency vulnerability scan | Medium | Low | Security | Add `cargo audit` (or equivalent) in CI. |
| P1 | Add coverage gate in CI | Medium | Medium | Testing | Add coverage reporting for `rust-cli` tests. |
| P1 | Align command docs with CLI command surface | Medium | Medium | Docs/Backend | Keep docs aligned to `rust-cli/src/main.rs` command set only. |
| P2 | Add structured audit logging for sync/upgrade | Medium | Medium | Backend | No centralized audit/correlation-id logging pattern. |
| P2 | Add docs drift check in CI | Medium | Medium | DevOps | Block PRs with stale command/path references in active docs. |

## Cross-Repo Impact

| Severity | Owner | Impact | Action |
|---|---|---|---|
| Medium | Product/Docs | Command naming drift can confuse downstream agent packs and user docs consumers. | Standardize on one naming path and update generated docs templates accordingly. |

## Blockers

- External scanners/tools (e.g., `cargo audit`, gitleaks) may require CI environment updates and policy approval.

## Related

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
