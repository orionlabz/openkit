# RISK REGISTER

**Sprint**: Sprint-09
**Status**: Active

| ID | Risk | Probability | Impact | Mitigation |
|---|---|---|---|---|
| R09-001 | Rust `check` output drifts from expected UX | Medium | Medium | Keep human output close to Go and add JSON contract tests |
| R09-002 | Parity implementation introduces regressions in release flow | Medium | High | Keep CI + release checks mandatory on every merge |
| R09-003 | Go decommission attempted before command parity | Medium | High | Gate decommission on Sprint-09/10 parity checklist |

## Related

- [[sprint/Sprint-09/TASKS.md]]
- [[sprint/Sprint-09/BACKLOG.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
