# Rust Single Runtime Decision

## Context

After completing Memory Kernel v1 and release hardening, the project decision changed from bridge/sidecar migration to full Rust cutover.

The new target is a single binary named `openkit` implemented in Rust. The Go CLI and runtime bridge are now considered transitional artifacts to be removed.

## Decision

1. OpenKit runtime becomes Rust-only.
2. Binary name is `openkit` (no `openkit-rs` sidecar model).
3. Release artifacts publish only `openkit_<OS>_<ARCH>` packages.
4. Installers download and install only the Rust binary.
5. Legacy Go command runtime is decommissioned after parity gates pass.

## Rationale

- Removes dual-runtime complexity.
- Eliminates bridge and sidecar operational failure modes.
- Simplifies install and release flows.
- Aligns product identity with one execution engine.

## Consequences

- Sprint-08 must prioritize command parity and release cutover.
- Existing bridge/sidecar docs and workflows must be removed.
- CI/release validation shifts from cross-runtime integration to Rust parity gates.

## Related

- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[sprint/Sprint-08/HUB-SPRINT-08.md]]
- [[SUPPORT_MATRIX.md]]
