---
description: Structured brainstorming for projects and features (OPTIONAL, use when scope unclear).
---

# /brainstorm - Structured Idea Exploration (OPTIONAL)

$ARGUMENTS

## Purpose

This command activates BRAINSTORM mode for structured idea exploration. Use when you need to explore options before committing to an implementation.

**IMPORTANT:** This command is OPTIONAL and does NOT replace `/context`.
- `/context` is MANDATORY and generates technical documentation
- `/brainstorm` is OPTIONAL and explores options when scope is unclear

**When to use /brainstorm:**
- Multiple valid approaches exist
- Scope is ambiguous
- User needs help deciding direction
- Comparing frameworks, architectures, or strategies

**Discovery Gate Flow:**
```
/context (MANDATORY) → /specify → /clarify → /plan → /tasks → /impl
    ↑
/brainstorm (OPTIONAL) ← use when scope unclear
```

## Behavior

When `/brainstorm` is triggered:

1. **Understand the goal**
   - Use the question tool to ask what problem we are solving.
   - Use the question tool to ask who the user is.
   - Use the question tool to ask what constraints exist.

2. **Generate options**
   - Provide at least 3 different approaches
   - Each with pros and cons
   - Consider unconventional solutions

3. **Compare and recommend**
   - Summarize tradeoffs
   - Give a recommendation with reasoning

---

## Output Format

```markdown
## Brainstorm: [Topic]

### Context
[Brief problem statement]

---

### Option A: [Name]
[Description]

 **Pros:**
- [benefit 1]
- [benefit 2]

 **Cons:**
- [drawback 1]

 **Effort:** Low | Medium | High

---

### Option B: [Name]
[Description]

 **Pros:**
- [benefit 1]

 **Cons:**
- [drawback 1]
- [drawback 2]

 **Effort:** Low | Medium | High

---

### Option C: [Name]
[Description]

 **Pros:**
- [benefit 1]

 **Cons:**
- [drawback 1]

 **Effort:** Low | Medium | High

---

## Recommendation

**Option [X]** because [reasoning].

Use the question tool to ask: What direction would you like to explore?
```

---

## Examples

```
/brainstorm authentication system
/brainstorm state management for complex form
/brainstorm database schema for social app
/brainstorm caching strategy
```

---

## Key Principles

- **No code** - this is about ideas, not implementation
- **Visual when helpful** - use diagrams for architecture
- **Honest tradeoffs** - don't hide complexity
- **Defer to user** - present options, let them decide

---

## TodoList Usage (MANDATORY for Complex Brainstorming)

When brainstorming complex projects with 3+ options or multiple decision points:

### Phase 0: TodoList Setup

1. **Read existing todolist:**
   ```javascript
   todoread()
   ```

2. **Create brainstorming todolist (using standard ID schema):**
   ```javascript
   todowrite({
     todos: [
       {
         id: "brainstorm-01-context",
         content: "Gather context and requirements",
         status: "in_progress",
         priority: "high"
       },
       {
         id: "brainstorm-02-options",
         content: "Generate 3+ viable options",
         status: "pending",
         priority: "high"
       },
       {
         id: "brainstorm-03-analysis",
         content: "Analyze trade-offs",
         status: "pending",
         priority: "medium"
       },
       {
         id: "brainstorm-04-recommend",
         content: "Provide recommendation",
         status: "pending",
         priority: "high"
       }
     ]
   })
   ```

### Phase 1: Context Gathering

- Use the question tool to ask what problem we are solving
- Use the question tool to ask who the user is
- Use the question tool to ask what constraints exist

**After context:**
- Update todolist: Mark "brainstorm-01-context" as `completed`
- Mark "brainstorm-02-options" as `in_progress`

### Phase 2: Generate Options

- Provide at least 3 different approaches
- Each with pros and cons
- Consider unconventional solutions

**After options:**
- Update todolist: Mark "brainstorm-02-options" as `completed`
- Mark "brainstorm-03-analysis" as `in_progress`

### Phase 3: Compare and Recommend

- Summarize tradeoffs
- Give a recommendation with reasoning

**After recommendation:**
- Update todolist: Mark "brainstorm-03-analysis" and "brainstorm-04-recommend" as `completed`

### Use Question Tool

Always use the question tool to ask: "What direction would you like to explore?" after presenting options
