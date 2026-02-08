# Acceptance Criteria: Remove Blueprint References

- No repository-tracked files contain blueprint alias references.
- No repository-tracked files contain the term `blueprints` in a way that implies a shipped feature/directory.
- References are replaced with the correct shipped locations (e.g. `internal/templates/base/` or `.opencode/`).
- The audit/check scripts do not mention "templates/blueprints" as if both exist; comments reflect reality.
- `python .opencode/scripts/checklist.py .` succeeds.
- `python .opencode/skills/vulnerability-scanner/scripts/security_scan.py .` succeeds.
