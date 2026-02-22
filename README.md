# OpenKit CLI

OpenKit is a Rust CLI for project bootstrap, agent pack sync, environment checks, upgrade/uninstall, and Memory Kernel maintenance. It supports multiple AI agent environments and generates a ready-to-use scaffold for each one.

## Install

### macOS / Linux / WSL

```bash
curl -fsSL https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.ps1 | iex
```

### Manual Download

Download from [latest release](https://github.com/orionlabz/openkit/releases/latest):

- `openkit_Darwin_x86_64.tar.gz`
- `openkit_Darwin_arm64.tar.gz`
- `openkit_Linux_x86_64.tar.gz`
- `openkit_Linux_arm64.tar.gz`
- `openkit_Windows_x86_64.zip`

## Quick Start

```bash
# Bootstrap a new project for Claude Code
openkit init my-project --agent claude

# Bootstrap for OpenCode
openkit init my-project --agent opencode

# Check environment dependencies
openkit check
```

After `init`, open the project directory in your AI agent and use the slash commands (`/discover`, `/specify`, `/create`, `/verify`, `/deploy`) to drive the development workflow.

## Supported Agents (`--agent`)

| Agent | Config dir | Root file |
|-------|-----------|-----------|
| `claude` | `.claude/` | `CLAUDE.md` |
| `opencode` | `.opencode/` | `opencode.json` |
| `codex` | `.codex/` | — |
| `cursor` | `.cursor/` | — |
| `gemini` | `.gemini/` | — |
| `antigravity` | `.antigravity/` | — |

Each agent scaffold includes: commands, rules, prompts, skills, and a Memory Kernel under `.openkit/`.

## Command Surface

```bash
# Environment checks
openkit check
openkit check --json

# Project bootstrap
openkit init [project-name] --agent claude
openkit init [project-name] --agent opencode --no-git
openkit init --overwrite --no-git

# Agent pack lifecycle
openkit sync --agent claude --overwrite
openkit sync --agent opencode --prune
openkit doctor --agent claude --json

# Binary lifecycle
openkit upgrade --check
openkit upgrade --dry-run
openkit upgrade
openkit uninstall --dry-run
openkit uninstall --yes

# Memory Kernel maintenance
openkit memory init
openkit memory doctor --json
openkit memory capture --session-id s01 --summary "Sprint work" --action check
openkit memory review --json
```

## What Gets Created

Running `openkit init my-project --agent claude` produces:

```
my-project/
├── CLAUDE.md                    # Project guidance (read automatically by Claude Code)
├── .claude/
│   ├── settings.json            # Claude Code permission settings
│   ├── commands/                # Slash commands: /discover /specify /create /verify /debug /deploy /orchestrate
│   ├── rules/                   # MASTER.md, MEMORY_KERNEL.md, DOCS_FILE_GLOSSARY.md
│   ├── prompts/                 # Agent prompts
│   └── skills/                  # Reusable skill packs
├── memory/                      # Project memory: context, security, quality gates, sprints
└── .openkit/                    # OpenKit internal state and Memory Kernel config
```

## Development Workflow

Once initialized, use slash commands inside your agent:

```
/discover   → analyze the codebase and generate a context pack (run this first)
/specify    → write spec, plan, and task breakdown for a feature
/create     → implement tasks from the spec
/verify     → run lint, tests, and security checks
/deploy     → deploy with pre-flight checklist and rollback procedure
/debug      → systematic 4-phase root cause analysis
/orchestrate → coordinate complex multi-agent missions
```

## Keeping Agent Packs Updated

```bash
# Pull latest templates without touching user files
openkit sync --agent claude --overwrite

# Also remove files no longer managed by OpenKit
openkit sync --agent claude --overwrite --prune
```

## Upgrade Behavior

- `openkit upgrade --check`: queries the latest release tag from GitHub.
- `openkit upgrade`: Linux/macOS downloads the release artifact, verifies SHA-256 against `checksums.txt`, and replaces the binary with rollback. Windows delegates to the PowerShell installer.
- `openkit upgrade --dry-run`: prints the planned update source/asset without changing binaries.

## From Source

```bash
cargo fmt --manifest-path rust-cli/Cargo.toml --all --check
cargo clippy --manifest-path rust-cli/Cargo.toml --all-targets -- -D warnings
cargo build --release --manifest-path rust-cli/Cargo.toml
cargo test --manifest-path rust-cli/Cargo.toml
```

## Platform Support

- macOS: `x86_64`, `arm64`
- Linux: `x86_64`, `arm64`
- Windows: `x86_64`
