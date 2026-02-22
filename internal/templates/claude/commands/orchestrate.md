---
description: Universal orchestrator for complex multi-agent missions
---

# /orchestrate - Universal Orchestrator

$ARGUMENTS

## Overview

The orchestrator coordinates multiple specialized agents for complex tasks. This command unifies:
- `/engineer` (old) - Original orchestrator
- `/ui-ux` (old) - Design work
- For new applications, use `/orchestrate` directly

## If $ARGUMENTS is empty

Ask the user: "Describe the complex task you need executed."

---

## Mode Detection

### Router Mode (Simple Tasks)

**Triggers when:**
- Keywords: "specify", "clarify", "plan", "verify", "test", "debug", "deploy"
- Single domain tasks

**Action:** Redirect to appropriate command:
- "test..." → `/verify`
- "debug..." → `/debug`
- "deploy..." → `/deploy`
- "discover..." → `/discover`

### Orchestrator Mode (Complex Tasks)

**Triggers when:**
- Multiple domains (backend + frontend + database + security)
- Keywords: "build", "create", "full", "system", "platform"
- Large feature or new project

**Action:** Start multi-phase orchestration

---

## Orchestration Protocol

### Phase 1: Analysis & Planning

1. **Chain of Thought (MANDATORY):**
   - What did the user ask?
   - What is the implicit goal?
   - Which specialists are required?
   - Why is this complex enough?

2. **Create task list** with phases: analysis, discovery, specification, P0–P3, verification.

3. **Run `/discover`** if not already done

4. **Run `/specify`** for full specification

---

### Phase 2: Implementation

**Execution Order (P0 → P1 → P2 → P3):**

**P0 - Foundation:**
- Database setup (if DB needed)
- Security audit (always)

**STOP:** Ask "Foundation phase complete. Proceed to P1 (Backend)?"

**P1 - Core Backend:**
- API, business logic, models

**STOP:** Ask "Backend phase complete. Proceed to P2 (Frontend)?"

**P2 - UI/UX:**
- UI components (WEB projects)
- Mobile screens (MOBILE projects)

**STOP:** Ask "Frontend phase complete. Proceed to P3 (Polish)?"

**P3 - Polish:**
- Tests
- Performance optimization (if needed)

---

### Phase 3: Verification

Run `/verify` automatically

---

## STOP Point

Ask the user: "All phases executed. Mark project complete?"
- Yes, complete
- Review results
