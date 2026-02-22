# CLAUDE.md

This file provides guidance to Claude Code when working with code in this repository.

## Agent Commands

Use these slash commands to drive the development workflow:

| Command | Description |
|---------|-------------|
| `/discover` | Analyze project context (MANDATORY before /specify) |
| `/specify` | Full specification + planning + task breakdown |
| `/create` | Execute implementation from plan |
| `/verify` | Quality verification — tests, lint, security |
| `/debug` | Systematic root cause analysis |
| `/deploy` | Safe deployment with health checks and rollback |
| `/orchestrate` | Universal orchestrator for complex multi-agent missions |

## Workflow

```
/discover (MANDATORY) → /specify → /create → /verify → /deploy
```

For complex new projects, use `/orchestrate` directly.

## Rules and Conventions

See `@.claude/rules/MASTER.md` for project-wide conventions.

Memory Kernel protocol: `@.claude/rules/MEMORY_KERNEL.md`

## Memory Layout

Project memory lives in `memory/`:

- `memory/CONTEXT.md` — executive summary and overview
- `memory/SECURITY.md` — threats, controls, gaps
- `memory/QUALITY_GATES.md` — linters, tests, CI checks
- `memory/ACTION_ITEMS.md` — backlog prioritized by impact × effort
- `memory/HUB-DOCS.md` — documentation hub
- `memory/GLOSSARY.md` — shared terminology
- `memory/requirements/` — feature specifications
- `memory/sprint/` — sprint planning artifacts
