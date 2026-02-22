---
description: Safe deployment with verification, health checks, and rollback capability
---

**DEPLOY MODE ACTIVATED**

**Target:** $ARGUMENTS (staging, production, preview, verify)

**If $ARGUMENTS is empty:**
Ask the user: "Which environment do you want to deploy to?"
- Staging (Pre-production testing)
- Production (Live environment)
- Preview (Temporary preview)
- Verify (Check current deployment)

---

## PRE-DEPLOYMENT CHECKLIST (MANDATORY)

**Before ANY deploy, verify:**

### Code Quality
- [ ] All tests passing
- [ ] Lint checks clean
- [ ] Type checks passing
- [ ] No console errors

### Security
- [ ] Security scan clean
- [ ] No secrets in code
- [ ] Dependencies updated

### Documentation
- [ ] Changelog updated
- [ ] README current
- [ ] API docs updated

### Backup & Rollback
- [ ] Database backup confirmed
- [ ] Previous version tagged
- [ ] Rollback plan ready

---

## Deploy Targets

### STAGING
**Purpose:** Pre-production testing

**Steps:**
1. Build the application
2. Deploy to staging server
3. Smoke tests
4. Functional verification
5. Notify the team

**Typical command:**
```bash
npm run build
# [staging deploy command]
```

### PRODUCTION
**Purpose:** Live environment

**MANDATORY STOP:**
> "DEPLOY TO PRODUCTION
>
> Pre-deployment checklist:
> - [ ] Staging tests passed
> - [ ] Team approval
> - [ ] Low-traffic window
>
> **TYPE EXACTLY 'deploy production confirm' TO PROCEED**
>
> In case of failure, rollback will be available in [X minutes]."

**Steps:**
1. Explicit approval obtained
2. Database backup
3. Tag current version
4. Deploy during low-traffic window
5. Monitor metrics
6. Health checks
7. Smoke tests in production

### PREVIEW
**Purpose:** Preview of feature branches

**Steps:**
1. Build feature branch
2. Deploy to temporary preview URL
3. Share link with stakeholders
4. Auto-destroy after X days

### VERIFY
**Purpose:** Verify current deployment

**Steps:**
1. Check health endpoints
2. Verify metrics
3. Smoke tests
4. Report status

---

## Deployment Methods

**Recommended: Blue-Green Deployment**
```
1. Deploy to Green (idle)
2. Test Green
3. Switch traffic: Blue → Green
4. Monitor
5. Keep Blue for quick rollback
```

**Alternatives:**
- **Rolling:** One node at a time
- **Canary:** 5% → 25% → 100% traffic

---

## VERIFICATION STEPS

### Health Checks
```bash
# Critical endpoints
curl -f http://localhost:3000/health
curl -f http://localhost:3000/api/health

# Metrics
curl http://localhost:3000/metrics
```

### Monitoring
Observe for 10-30 minutes:
- Error rate (should stay < baseline)
- Response time (should stay < baseline)
- CPU/Memory usage
- Database connections
- User complaints/feedback

---

## ROLLBACK PROCEDURE

**If deployment fails:**

### Immediate (0-5 min)
1. Stop deployment
2. Switch traffic to previous version
3. Assess impact

### Short term (5-30 min)
1. Investigate root cause
2. Fix critical issues
3. Test the fix
4. Decide: retry or abort

### Long term (30+ min)
1. Full incident response
2. Post-mortem
3. Preventive actions

---

## DEPLOYMENT REPORT

Document after deploy:

```markdown
# Deployment Report: [Version] → [Environment]

## Timestamp
[Date/time]

## Changes
- [List of main changes]

## Verification
- [ ] Health checks pass
- [ ] Smoke tests pass
- [ ] Metrics normal
- [ ] No user impact

## Issues
[None / List]

## Rollback Ready
[Previous version: X.X.X]
[Rollback command available]
```

---

## BEST PRACTICES

1. **NEVER deploy on Friday** (unless emergency)
2. **ALWAYS have rollback ready**
3. **Deploy in low-traffic windows**
4. **Monitor for at least 30 minutes post-deploy**
5. **Keep deploys small and frequent**
6. **Automate everything you can**
7. **Test on staging first**
8. **Communicate with the team**

---

## STOP POINTS

**1. Before Production:**
> Ask the user: "Ready for production? Confirm explicitly."

**2. During Deploy:**
> Ask the user: "Deployment in progress. Metrics normal?"

**3. After Deploy:**
> Ask the user: "Deployment complete! Health checks passed. Confirm success?"
