---
description: Quality verification - tests, lint, security, performance
---

# /verify - Quality Verification

$ARGUMENTS

## Overview

Run comprehensive quality checks on the codebase. This command unifies (old commands):
- `/test` (old) - Test execution
- `/checklist` (old) - Quality checklist
- `/analyze` (old) - Cross-artifact analysis

## If $ARGUMENTS is empty

Ask the user: "Which verification scope?"
- All (Full verification suite)
- Quick (Lint + Security only)
- Custom (Specific checks)

---

## Verification Protocol (by priority)

### P0 - Critical Checks

#### 1. Lint & Type Check
```bash
npm run lint
npx tsc --noEmit
```
- Status: [PASS/FAIL]

#### 2. Security Scan
Run `.claude/skills/vulnerability-scanner/scripts/security_scan.py .` if available.
- Status: [PASS/WARN/FAIL]
- Findings: [details]

### P1 - Quality Checks

#### 3. Unit Tests
```bash
npm test
# or
pytest
```
- Status: [PASS/FAIL]
- Coverage: [XX%]
- Failures: [if any]

#### 4. UX/Accessibility Audit (if frontend)
Run `.claude/skills/frontend-design/scripts/ux_audit.py .` if available.
- Status: [PASS/WARN]
- Issues: [if any]

### P2 - Performance

#### 5. Build Verification
```bash
npm run build
```
- Status: [PASS/FAIL]
- Warnings: [if any]

#### 6. Lighthouse Audit (requires running server)
If a server is detected at http://localhost:3000:
Run `.claude/skills/performance-profiling/scripts/lighthouse_audit.py http://localhost:3000` if available.
- Score: [XX/100]
- Web Vitals: [LCP, FID, CLS]

### P3 - E2E Tests

#### 7. Playwright E2E (requires server)
Run `.claude/skills/webapp-testing/scripts/playwright_runner.py http://localhost:3000 --screenshot` if available.
- Status: [PASS/FAIL]
- Screenshots: [path]

---

## Final Report

```markdown
## Verification Results Summary

| Check | Status | Details |
|-------|--------|---------|
| Lint | PASS/FAIL | ... |
| Type Check | PASS/FAIL | ... |
| Security | PASS/WARN/FAIL | ... |
| Unit Tests | PASS/FAIL | XX% coverage |
| UX Audit | PASS/WARN | ... |
| Build | PASS/FAIL | ... |
| Lighthouse | PASS/WARN | XX/100 |
| E2E Tests | PASS/FAIL | ... |

### Action Items
- [ ] [If there are failures, list required fixes]
```

---

## STOP Point

Ask the user: "Results: [PASS/FAIL]. Proceed to deploy (/deploy)?"
- Yes, proceed to /deploy
- Fix issues first

**IMPORTANT:** Do not mark checks as passing without actually running the commands!
