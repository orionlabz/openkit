# BACKLOG

**Sprint**: Sprint-08
**Title**: Rust Single Runtime Cutover
**Status**: In Progress

## Stories

### S08-001: Rust Binary Canonicalization

- Description: Rename and publish Rust runtime as canonical `openkit` binary.
- Priority: P0
- Status: Completed
- Links: [[requirements/memory-kernel-rust-cli/RUST_SINGLE_RUNTIME_DECISION.md]]

### S08-002: Release Pipeline Cutover

- Description: Replace bridge/sidecar release model with Rust-only release artifacts.
- Priority: P0
- Status: Completed
- Links: [[requirements/memory-kernel-rust-cli/PLAN.md]]

### S08-003: Installer Simplification

- Description: Update install scripts to install only Rust `openkit` binary.
- Priority: P1
- Status: Completed
- Links: [[DEPRECATIONS.md]], [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]

### S08-004: Go Runtime Decommission

- Description: Remove bridge/legacy Go entrypoints once parity checks pass.
- Priority: P1
- Status: In Progress
- Links: [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]

## Related

- [[sprint/Sprint-08/TASKS.md]]
- [[sprint/Sprint-08/RISK_REGISTER.md]]
- [[sprint/Sprint-08/SPRINT_GOAL.md]]
