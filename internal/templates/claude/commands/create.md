---
description: Execute implementation from plan with multi-agent coordination
---

# /create - Implementation

$ARGUMENTS

## Overview

Execute implementation tasks from the specification. This command:
- Reads tasks from `memory/sprint/Sprint-XX/TASKS.md`
- Executes in priority order (P0 → P1 → P2 → P3)
- Coordinates multiple agents as needed

**IMPORTANT:** `/specify` MUST be complete before running this command.

**For new applications:** Use `/orchestrate` instead.

## If $ARGUMENTS is empty

Ask the user: "Which sprint to implement?" (Latest Sprint / Custom Path)

**Syntax:** `from memory/sprint/Sprint-XX/TASKS.md`

---

## Pre-Implementation Checklist

- [ ] `/discover` has been run
- [ ] `/specify` is complete
- [ ] Sprint is selected
- [ ] `TASKS.md` exists and is valid

**SDD Gate:** If spec is missing, STOP and direct to `/specify` first.

---

## Execution Order (P0 → P1 → P2 → P3)

### P0 - Foundation

- Run database setup (if DB needed)
- Run security audit (always)

**STOP:** Ask "Foundation complete. Proceed to P1 (Backend)?"

### P1 - Core Backend

- Implement API, business logic, models

**STOP:** Ask "Backend complete. Proceed to P2 (Frontend)?"

### P2 - UI/UX

- Implement UI components (WEB projects)
- Implement mobile screens (MOBILE projects)

**STOP:** Ask "Frontend complete. Proceed to P3 (Polish)?"

### P3 - Polish

- Write and run tests
- Optimize performance (if needed)

---

## Progress Updates

Mark tasks in `memory/sprint/Sprint-XX/TASKS.md`:

```markdown
- [x] task-01: [Name]  COMPLETE
- [ ] task-02: [Name]  IN PROGRESS
```

---

## Final STOP

Ask the user: "All phases complete. Run verification (/verify)?"
- Yes, run /verify
- Later
