---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# DOCUMENTATION FILE GLOSSARY

Canonical documentation filenames for all OpenKit projects.

## Naming Convention

- Documentation files MUST use canonical uppercase naming.
- Artifact files use uppercase snake case (for example `PROBLEM_STATEMENT.md`, `TASKS.md`, `TECH_STACK.md`).
- Hub files use the `HUB-<RESOURCE>.md` convention (for example `HUB-DOCS.md`, `HUB-SPRINT-XX.md`).

## Core Docs (always available)

- `docs/HUB-DOCS.md`
- `docs/GLOSSARY.md`
- `docs/CONTEXT.md`
- `docs/SECURITY.md`
- `docs/QUALITY_GATES.md`
- `docs/ACTION_ITEMS.md`
- `docs/ARCHITECTURE.md`
- `docs/COMMANDS.md`
- `docs/SKILLS.md`
- `docs/WORKFLOW.md`

## Contextual Docs (create only when applicable)

- `docs/FRONTEND.md`
- `docs/BACKEND.md`
- `docs/API.md`
- `docs/DATABASE.md`
- `docs/CHANGELOG.md`
- `docs/MIGRATION_CHECKLIST.md`

## Requirements Docs

- `docs/requirements/HUB-REQUIREMENTS.md`
- `docs/requirements/<feature>/HUB-<FEATURE>.md`
- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/DATA_CONTRACTS.md`
- `docs/requirements/<feature>/RISKS.md`
- `docs/requirements/<feature>/PLAN.md`
- `docs/requirements/<feature>/RESEARCH.md` (optional)
- `docs/requirements/<feature>/QUICKSTART.md` (optional)
- `docs/requirements/<feature>/ANALYSIS.md` (optional)
- `docs/requirements/<feature>/TECH_STACK.md` (optional)
- `docs/requirements/<feature>/CHECKLIST.md` (optional)

## Sprint Docs

- `docs/sprint/HUB-SPRINTS.md`
- `docs/sprint/Sprint-XX/HUB-SPRINT-XX.md`
- `docs/sprint/Sprint-XX/SPRINT_GOAL.md`
- `docs/sprint/Sprint-XX/BACKLOG.md`
- `docs/sprint/Sprint-XX/TASKS.md`
- `docs/sprint/Sprint-XX/RISK_REGISTER.md`

## ADR and Runbooks

- `docs/adr/HUB-ADR.md`
- `docs/adr/ADR_0001_<TITLE>.md`
- `docs/runbooks/HUB-RUNBOOKS.md`
- `docs/runbooks/<SERVICE>.md`

## Rules

- If a command, agent, or skill references a documentation filename, it MUST use this glossary.
- If a new canonical docs file is introduced, update this file in the same change.
