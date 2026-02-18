# Sprint Goal: Sprint-08

Cut over OpenKit to a single Rust runtime and remove bridge/sidecar complexity from CLI, release, and installer flows.

## Success Indicators

- Rust binary is named and published as `openkit`.
- Release workflow publishes only Rust `openkit` artifacts.
- Installers download only Rust `openkit` artifacts.
- Legacy Go runtime command entrypoints are removed.

## Related

- [[sprint/Sprint-08/BACKLOG.md]]
- [[sprint/Sprint-08/TASKS.md]]
- [[requirements/memory-kernel-rust-cli/RUST_SINGLE_RUNTIME_DECISION.md]]
