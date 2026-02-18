# RISK REGISTER

**Sprint**: Sprint-08
**Status**: Active

| ID | Risk | Probability | Impact | Mitigation |
|---|---|---|---|---|
| R08-001 | Release pipeline drift during Rust cutover | Medium | High | Add CI guards and release dry-run checks |
| R08-002 | Installer breakage on one platform | Medium | High | Add multi-platform install smoke tests |
| R08-003 | Command parity gap after Go removal | Medium | High | Gate decommission on golden parity tests |

## Related

- [[sprint/Sprint-08/TASKS.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
