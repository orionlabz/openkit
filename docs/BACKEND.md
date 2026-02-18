# BACKEND

## Endpoints

| Method | Route | Handler | Auth | Notes |
|---|---|---|---|---|
| N/A | not found | not found | not found | Project is a CLI application, not an HTTP server. |

## Jobs / Async

- `openkit upgrade --check` and `openkit upgrade` call GitHub release API synchronously.
- All command execution is process-local; there is no background worker subsystem.

## Backend Components

| Component | Location | Responsibility |
|---|---|---|
| Command dispatch | `rust-cli/src/main.rs` | Clap root command, subcommand routing, runtime entrypoint. |
| Agent operations | `rust-cli/src/main.rs` | `openkit <agent> sync|doctor|upgrade` baseline behaviors. |
| Memory kernel ops | `rust-cli/src/main.rs` | `memory init|doctor|capture|review` docs-first operations. |
| Self-update | `rust-cli/src/main.rs` | Fetch release metadata, download artifact, checksum verify, atomic replace. |
| Release pipeline | `.github/workflows/release.yml` | Multi-platform Rust build and asset publication. |

## Evidence

- `rust-cli/src/main.rs`: command definitions and runtime dispatch.
- `rust-cli/tests/command_contracts.rs`: command behavior contract tests.
- `.github/workflows/ci.yml`: Rust fmt/clippy/build/test checks.
- `.github/workflows/release.yml`: Rust release packaging and upload.

## Related

- [[CONTEXT.md]]
- [[API.md]]
- [[SECURITY.md]]
