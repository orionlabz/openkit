# Rust Parity Matrix

## Context

This matrix tracks command-level migration progress from legacy Go runtime to Rust `openkit` runtime.

## Command Parity Status

| Command Surface | Rust Status | Notes |
|---|---|---|
| `openkit --version` | Complete | Version flag supported in Rust root command |
| `openkit check` | Complete | Human output + `--json` machine-readable mode |
| `openkit memory init|doctor|capture|review` | Complete | Implemented and covered by contract tests |
| `openkit init` | Baseline Complete | Rust command creates baseline project docs/structure |
| `openkit <agent> sync` | Baseline Complete | Creates agent config marker and supports dry-run/overwrite |
| `openkit <agent> doctor` | Baseline Complete | Reports config health with optional JSON output |
| `openkit <agent> upgrade` | Baseline Complete | Baseline sync with overwrite semantics |
| `openkit uninstall` | Not Implemented | To be implemented in Rust roadmap |
| `openkit upgrade` (self-update) | Not Implemented | To be implemented in Rust roadmap |

## Decommission Gate

Go runtime decommission can start only when all P0 and P1 command surfaces above are marked complete and validated in CI.

## Related

- [[sprint/Sprint-09/TASKS.md]]
- [[sprint/Sprint-09/BACKLOG.md]]
- [[requirements/memory-kernel-rust-cli/PLAN.md]]
