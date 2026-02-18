# TASKS

**Sprint**: Sprint-08
**Title**: Rust Single Runtime Cutover
**Status**: In Progress
**Total Tasks**: 8

## Task Breakdown

### Task 1: Rename Rust runtime to canonical `openkit`

**ID**: S08-T001
**Story**: S08-001
**Priority**: P0

**VERIFY**:
- [x] Rust crate binary name is `openkit`
- [x] Rust tests execute against `openkit`

### Task 2: Remove bridge/sidecar references from docs policy

**ID**: S08-T002
**Story**: S08-001
**Priority**: P0

**VERIFY**:
- [x] Decision artifact created for single-runtime cutover
- [x] Requirement plan and criteria updated

### Task 3: Create Sprint-08 execution artifacts

**ID**: S08-T003
**Story**: S08-001
**Priority**: P0

**VERIFY**:
- [x] Sprint hub, backlog, tasks, risk register created

### Task 4: Cut over release pipeline to Rust-only `openkit` artifacts

**ID**: S08-T004
**Story**: S08-002
**Priority**: P0

**VERIFY**:
- [x] Release workflow no longer publishes sidecar runtime assets
- [x] Release workflow publishes only canonical `openkit` artifacts

### Task 5: Simplify install scripts to single runtime

**ID**: S08-T005
**Story**: S08-003
**Priority**: P1

**VERIFY**:
- [x] `install.sh` installs only Rust `openkit`
- [x] `install.ps1` installs only Rust `openkit`

### Task 6: Decommission Go memory runtime bridge

**ID**: S08-T006
**Story**: S08-004
**Priority**: P1

**VERIFY**:
- [x] Bridge files removed from `internal/cli/`
- [x] No bridge env vars remain in docs/help

### Task 7: Execute parity and regression suite

**ID**: S08-T007
**Story**: S08-004
**Priority**: P1

**VERIFY**:
- [x] Rust contract tests pass
- [x] CI and lint gates pass post-cutover

### Task 8: Publish sprint exit report

**ID**: S08-T008
**Story**: S08-004
**Priority**: P1

**VERIFY**:
- [x] Acceptance criteria mapping completed
- [x] Open risks documented for next sprint

## Related

- [[sprint/Sprint-08/HUB-SPRINT-08.md]]
- [[sprint/Sprint-08/BACKLOG.md]]
- [[sprint/Sprint-08/EXIT_REPORT.md]]
- [[requirements/memory-kernel-rust-cli/RUST_SINGLE_RUNTIME_DECISION.md]]
