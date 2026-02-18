# SECURITY

## Threats

- Supply-chain tampering during binary upgrade flow.
- Over-privileged agent execution (`bash`, `edit`, `write`) in misconfigured usage.
- Drift/conflict bypass risk when users force `--overwrite`.
- Missing CI security checks allows vulnerable dependencies or leaked secrets.

## Controls

- SHA256 verification for downloaded upgrade artifacts.
- Runtime command contracts covered by Rust tests.
- CI lint/test/build pipeline on `main` PR/push.

## Gaps

- No signature verification (Cosign/GPG/minisign) in self-update path.
- No Rust dependency vulnerability scanner in CI.
- No secret scanning in CI.
- No dedicated structured audit log/correlation ID pattern for security operations.

## Prioritized Actions

| Priority | Action | Impact | Effort | Owner | Notes |
|---|---|---|---|---|---|
| P0 | Add signed release verification | High | Medium | Security/DevOps | Validate signatures in upgrade before install. |
| P0 | Add secret scanning in CI | High | Low | Security | Add gitleaks or equivalent in PR pipeline. |
| P1 | Add dependency scanning in CI | Medium | Low | Security | Add `cargo audit` (or equivalent) gate. |
| P1 | Add security-specific static analysis | Medium | Medium | Security | Add Rust-focused security linting/static checks. |
| P2 | Add audit-oriented structured logs | Medium | Medium | Backend | Capture sync/upgrade actions and outcomes. |

## Evidence

- `rust-cli/src/main.rs`: `run_self_update_unix()` validates SHA256 checksums from `checksums.txt`.
- `rust-cli/tests/command_contracts.rs`: validates command-level behavior.
- `.github/workflows/ci.yml`: includes lint/test/build only; no security or secret scan jobs.

## Related

- [[CONTEXT.md]]
- [[QUALITY_GATES.md]]
- [[ACTION_ITEMS.md]]
