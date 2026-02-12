# Docs Migration Checklist

## Scope

Migration of `docs/` to OpenKit Obsidian-compatible standard.

## Progress

- [x] Audited current `docs/` structure and naming.
- [x] Re-audited `docs/` for OpenKit Obsidian compliance (2026-02-11).
- [x] Renamed non-canonical files to uppercase canonical names where applicable.
- [x] Created mandatory hubs:
  - [x] `docs/README.md` (already existed, normalized)
  - [x] `docs/requirements/README.md`
  - [x] `docs/sprint/README.md`
  - [x] `docs/requirements/<feature>/README.md` for all features
  - [x] `docs/sprint/Sprint-XX/README.md` for all Sprint-01..06
- [x] Added `docs/adr/README.md` hub.
- [x] Added `docs/content-protocol/README.md` hub.
- [x] Added `docs/agent-compat/agents/README.md` hub.
- [x] Normalized `## Related` sections across root docs, requirements, sprint artifacts, ADR, and protocol docs.
- [x] Converted key internal path references to Obsidian wikilinks.
- [x] Ran Obsidian link lint and fixed broken links.
- [x] Removed stray nested hub file `docs/docs/README.md` (unused and out of hub structure).

## Rename Map

- `docs/requirements/cursor-codex-targets/analysis.md` -> `docs/requirements/cursor-codex-targets/ANALYSIS.md`
- `docs/agent-compat/agents/claude.md` -> `docs/agent-compat/agents/CLAUDE.md`
- `docs/agent-compat/agents/codex.md` -> `docs/agent-compat/agents/CODEX.md`
- `docs/agent-compat/agents/cursor.md` -> `docs/agent-compat/agents/CURSOR.md`
- `docs/agent-compat/agents/gemini.md` -> `docs/agent-compat/agents/GEMINI.md`
- `docs/agent-compat/agents/opencode.md` -> `docs/agent-compat/agents/OPENCODE.md`

## Risks / Notes

- Legacy text references in code blocks/backticks remain in historical notes when they are not navigational links.
- Some files include historical paths under `cli/...`; these were preserved as context unless they were clear internal doc references.
- No content semantics were rewritten; migration remained structural (hubs, links, and file placement).

## Related

- [[docs/README.md]]
- [[docs/requirements/README.md]]
- [[docs/sprint/README.md]]
- [[.opencode/rules/OBSIDIAN_LINKING.md]]
- [[.opencode/rules/DOCS_FILE_GLOSSARY.md]]
