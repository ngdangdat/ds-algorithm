---
title: "feat: Add /next and /checkout commands for daily LeetCode practice"
type: feat
status: active
date: 2026-05-28
origin: docs/brainstorms/2026-05-28-daily-practice-commands.md
---

# feat: Add /next and /checkout commands for daily LeetCode practice

## Summary

Two Pi slash commands that automate the daily LeetCode workflow: `/next` reads a state file and scaffolds a problem directory with Go boilerplate, `/checkout` runs tests, grades the solution per the strict CLAUDE.md criteria, regenerates a progress checklist, and pushes to git. Both commands are defined as Pi command markdown files interpreted by the AI agent, plus a JSON state file that persists position and completion data across sessions.

---

## Problem Frame

The learner works through a structured roadmap of ~100 LeetCode problems across 7 categories in `daily/02/README.md`. The manual workflow — scanning the roadmap to find the next problem, creating Go directories and module boilerplate, self-grading, and tracking progress — adds cognitive overhead. Answering "where was I?" after a gap requires mentally reconstructing completion state. The proposed commands eliminate this overhead: one command to start, one to finish.

---

## Requirements

**State management**
- R1. A state file tracks the current category and the problem identifier (slug or numeric ID) of the current problem. Position is determined by scanning the README from that identifier forward, so insertions and deletions in the roadmap don't cause silent desync.
- R2. The state file is stored in `daily/02/` and is human-readable.
- R3. Categories are traversed in the order they appear in `daily/02/README.md`: Arrays & Hashing → Two Pointers → Stack → Binary Search → Sliding Window → Linked List → Trees.

**`/next` command**
- R4. `/next` reads the state file and `daily/02/README.md` to determine the next uncompleted problem.
- R4a. `/next` is idempotent — if a problem directory already exists for the current position, `/next` reports the existing URL and exits without scaffolding.
- R5. `/next` creates the problem directory named `<NN>-leetcode-<problemID>`, where `<NN>` is zero-padded and sequential across the batch, and `<problemID>` is the numeric LeetCode problem ID. The README must include the numeric ID alongside each problem entry for extraction.
- R6. `/next` creates `main.go` and `main_test.go` in the directory, each containing only `package main`.
- R7. `/next` initializes a `go.mod` in the directory using `go mod init`, with a module path that distinguishes this problem.
- R8. `/next` outputs the LeetCode problem URL.
- R9. When the current category has no remaining problems, `/next` prompts the learner to advance to the next category instead of bootstrapping a problem.

**`/checkout` command**
- R10. `/checkout` reads the state file to identify which problem to grade.
- R11. `/checkout` runs `go test -v` in the problem directory and reports the test results.
- R12. `/checkout` grades the solution code using the strict evaluation criteria in CLAUDE.md.
- R13. `/checkout` regenerates `PROGRESS.md` from the state file after grading, reflecting the updated completion status and grade.
- R14. `/checkout` commits the solution and pushes to the remote.
- R15. After a successful checkout, `/checkout` advances the state file to the next position in the category.

**`PROGRESS.md` checklist**
- R16. `PROGRESS.md` is a generated artifact derived from the state file, containing a checkbox list of all problems from `daily/02/README.md` grouped by category with completion status and grades.
- R17. `PROGRESS.md` is created once during initial setup and updated on each `/checkout`.

**Origin actors:** A1 (Learner)
**Origin flows:** F1 (Start a new problem), F2 (Check out a solution)
**Origin acceptance examples:** AE1 (scaffold and output URL), AE2 (category exhaustion prompt), AE3 (grade, mark done, commit, advance), AE4 (failing tests → grade F, commit anyway)

---

## Scope Boundaries

- No LeetCode API integration — the learner visits the URL and solves on LeetCode.com.
- No test generation — the learner writes their own `main_test.go`.
- No out-of-order problem selection — the linear traversal is intentional for structured practice.
- No multi-language support — Go only.
- No automatic problem description scraping.

### Deferred to Follow-Up Work

- Consistency nudges (streak counter, "last practiced N days ago") — future iteration
- Machine-readable problem manifest (JSON/YAML) to decouple `/next` from README parsing — future iteration
- Directory-scanning as an alternative to state file — future evaluation
- Grade-trend analysis ("weakest category based on recent grades") — future iteration

---

## Context & Research

### Relevant Code and Patterns

- **Existing Pi command:** `.claude/commands/setup-leetcode.md` — manual, uses positional args (`$1`, `$2`). The new commands are argument-free and state-driven.
- **Go scaffold pattern from `daily/01/`:** Module path `github.com/ngdangdat/ds-algorithm/daily/<batch>/<NN>-leetcode-<ID>`, `package main` in both `main.go` and `main_test.go`, table-driven tests with struct slices.
- **Existing partial scaffold:** `daily/02/00-leetcode-217/` has `main.go` only — no `go.mod`, no `main_test.go`. Must be repaired during bootstrap.
- **README table format:** `| Problem Name | Difficulty | [Link](URL) |` with H3 category headers like `## 🔢 Arrays & Hashing`. Needs a numeric ID column added.
- **Git remote:** `https://github.com/ngdangdat/ds-algorithm` (push-capable).
- **Go toolchain:** `go1.26.0 darwin/arm64` (use `go 1.26.0` in generated `go.mod`).
- **`.gitignore`:** Ignores `*.out` and `**/target/**`.

### Institutional Learnings

No prior learnings exist (`docs/solutions/` is empty). This is greenfield workflow automation.

---

## Key Technical Decisions

- **State file as single source of truth:** JSON at `daily/02/.state.json`. PROGRESS.md is generated from state, never updated independently. Eliminates dual-write sync risk. (see origin: Key Decisions)
- **State schema — full completion ledger, not a cursor:** The state file stores all problems with per-problem `status` and `grade` fields, plus a `currentProblem` pointer. This supports PROGRESS.md regeneration (R13) without needing the cursor-only model from the original R1.
- **Numeric IDs in README:** Add a `| # |` column to all 7 category tables rather than scraping LeetCode or maintaining a separate manifest. The README is the canonical data source; including IDs directly is the simplest path to fulfill R5.
- **Duplicate problems re-scaffolded per category:** Problems appearing in multiple categories (e.g., "Sliding Window Maximum" in 3 categories) get separate directories and state entries. This aligns with the intentional linear traversal and the scope boundary against out-of-order selection.
- **go.mod version `go 1.26.0`:** Matches the installed toolchain. The daily/01 batch used `go 1.25.0` but the installed version has advanced.
- **NN counter — global, initialized from existing directories:** The counter starts at `max(existing NN) + 1` to avoid colliding with `00-leetcode-217`.
- **Push is best-effort:** `/checkout` commits locally and attempts push. If push fails (no network), it warns and still advances state — the learner can push manually later. This prevents network dependency from blocking practice completion.
- **Empty test file → refuse to grade:** If `main_test.go` has no test functions, `/checkout` reports "no tests to run" and refuses to assign a grade. This is more useful than auto-grading F — it prompts the learner to write tests rather than accepting an unearned failing grade.

---

## Open Questions

### Resolved During Planning

- **`/next`-twice behavior:** Idempotent — detects existing directory and reports URL, exits. No state change. (R4a)
- **State file format:** JSON. Structured, human-readable, Pi can parse/write reliably.
- **Commit message format:** `✅ [Problem Name] (LeetCode #[ID]) - Grade: [grade]`
- **Partial scaffold repair on `/next`:** If directory exists but is missing `go.mod` or `main_test.go`, `/next` completes the missing files rather than skipping.

### Deferred to Implementation

- Exact regex or parsing logic for README table extraction — Pi's natural language interpretation handles this; the command description should specify the table format contract clearly.
- Whether `git commit --allow-empty` is needed for the no-changes edge case — determined by actual usage patterns.

---

## High-Level Technical Design

The two commands share a common data flow: both read `.state.json` to determine the current problem, and both may write to it (advancing position). The README is read-only; the state file is the writable source of truth.

```
┌─────────────────────────────────────────────────────────┐
│                    daily/02/                             │
│                                                          │
│  README.md          .state.json        PROGRESS.md       │
│  (read-only)        (read/write)       (generated)       │
│  ┌──────────┐      ┌───────────┐      ┌───────────┐     │
│  │7 category│      │categories │      │checkbox   │     │
│  │tables w/ │ ←─── │+ problems│ ──── │list by    │     │
│  │numeric   │      │+ status  │      │category   │     │
│  │IDs       │      │+ grades  │      │+ grades   │     │
│  └──────────┘      │+ NN ctr  │      └───────────┘     │
│                     │+ current │                         │
│                     │  pointer │                         │
│                     └───────────┘                        │
│                          ▲                               │
│          ┌───────────────┴───────────────┐              │
│          │                               │              │
│     /next reads                     /checkout reads      │
│     /next writes                    /checkout writes     │
│     (scaffold +                     (grade +             │
│      update current)                 advance +           │
│                                      regenerate)         │
└─────────────────────────────────────────────────────────┘
```

### State file schema

```json
{
  "currentProblem": {
    "category": "arrays-and-hashing",
    "slug": "contains-duplicate"
  },
  "globalNN": 1,
  "categories": {
    "arrays-and-hashing": {
      "order": 0,
      "displayName": "Arrays & Hashing",
      "problems": [
        {
          "slug": "contains-duplicate",
          "numericId": 217,
          "name": "Contains Duplicate",
          "url": "https://leetcode.com/problems/contains-duplicate/",
          "nn": 0,
          "status": "in-progress",
          "grade": null
        }
      ]
    }
  }
}
```

- `status`: `"pending"` | `"in-progress"` | `"done"`
- `grade`: `null` | `"A"` | `"B"` | `"C"` | `"D"` | `"F"`
- Problem lookup: scan `categories[category].problems` for matching `slug`, advance to next `"pending"` entry.
- When no pending problems remain in a category, prompt to advance.

---

## Implementation Units

### U1. Update README with numeric problem IDs

**Goal:** Add a `#` column to all 7 category tables in `daily/02/README.md` so `/next` can extract numeric LeetCode problem IDs for directory naming.

**Requirements:** R5

**Dependencies:** None

**Files:**
- Modify: `daily/02/README.md`

**Approach:**
- Add `| # |` as the first column in each category's table header, separator, and all rows.
- Populate with the correct LeetCode problem number for each problem (e.g., Contains Duplicate → 217).
- Table format becomes: `| # | 🧩 Problem | ⚙️ Difficulty | 🔗 Link |`
- Row format becomes: `| 217 | Contains Duplicate | Easy | [Link](https://leetcode.com/problems/contains-duplicate/) |`

**Patterns to follow:**
- Existing README markdown table structure in `daily/02/README.md`

**Test scenarios:**
- Happy path: Each row in every category table has a numeric ID in the first column. All IDs match their corresponding LeetCode problem numbers.
- Edge case: Problems appearing in multiple categories (e.g., "3Sum" = 15, "Sliding Window Maximum" = 239) have the same numeric ID in every table they appear in.

**Verification:**
- `daily/02/README.md` has 7 category tables, each with a `#` column and correct numeric IDs. Spot-check 5-10 well-known problems against LeetCode to confirm accuracy.

---

### U2. Create `/next` command

**Goal:** Write `.claude/commands/next.md` — a Pi command that scaffolds the next LeetCode problem directory, updates state, and outputs the problem URL.

**Requirements:** R1, R2, R3, R4, R4a, R5, R6, R7, R8, R9

**Dependencies:** U1 (READMe needs numeric IDs)

**Files:**
- Create: `.claude/commands/next.md`

**Approach:**
- Command is a no-argument Pi slash command. The markdown file describes step-by-step instructions that Pi follows:
  1. **Bootstrap if needed:** If `daily/02/.state.json` doesn't exist, run the bootstrap flow (see U4).
  2. **Read state:** Parse `.state.json` to find `currentProblem` (category + slug).
  3. **Read README:** Parse `daily/02/README.md` to find the current category's table and the current problem's row.
  4. **Find next problem:** Scan forward from the current problem in the category table. Skip any with `status: "done"` in state. If none remain:
     - Prompt: "All [category name] problems are done. Advance to [next category]?" 
     - If yes: update state's `currentProblem` to first problem in next category, set status to `"in-progress"`, continue from step 3.
     - If no: exit with message listing remaining categories.
  5. **Construct directory name:** Read the problem's `nn` (from state) and numeric ID (from README `#` column). Directory = `daily/02/<NN>-leetcode-<ID>/` (NN zero-padded to 2 digits).
  6. **Idempotency check:** If directory already exists:
     - Check for missing files (no `go.mod`, no `main_test.go`). If any missing, create them (repair mode).
     - If all files present: report URL, exit. No state change.
  7. **Scaffold:** Create the directory. Write `main.go` and `main_test.go` with `package main`. Run `go mod init github.com/ngdangdat/ds-algorithm/daily/02/<NN>-leetcode-<ID>`. In `go.mod`, ensure `go 1.26.0` is set.
  8. **Update state:** Set the problem's `status` to `"in-progress"` in state. Update `currentProblem`. Increment `globalNN`. Write state file.
  9. **Output:** Print the LeetCode problem URL from the README. Print the directory path.

**Patterns to follow:**
- Existing Pi command format from `.claude/commands/setup-leetcode.md` (YAML frontmatter with `description`, then prose instructions)
- Go scaffold from `daily/01/00-leetcode-2942/` (module path convention, file names, package declaration)
- State file schema defined in High-Level Technical Design above

**Test scenarios:**
- Happy path — first run: No state file → bootstrap → scaffold `00-leetcode-217` → output URL. State shows `status: "in-progress"`.
- Happy path — subsequent run: State exists, next problem in category → scaffold `01-leetcode-242` → output URL. `globalNN` advances.
- Edge case — idempotent: Run `/next` again without `/checkout` → detect existing directory → report URL, exit, no state change.
- Edge case — partial repair: Directory exists but missing `go.mod` → create `go.mod`, report URL, exit.
- Edge case — category exhaustion: Last problem in category done → prompt to advance. Confirm → scaffold first problem in next category.
- Edge case — category exhaustion refused: Prompt → user says no → exit with remaining categories list.

**Verification:**
- Run `/next` from a clean state → directory `daily/02/00-leetcode-217/` created with `main.go`, `main_test.go`, `go.mod`. `.state.json` exists and points to the problem. `go test ./daily/02/00-leetcode-217/` compiles (no tests to run is expected).

---

### U3. Create `/checkout` command

**Goal:** Write `.claude/commands/checkout.md` — a Pi command that runs tests, grades the solution, regenerates PROGRESS.md, commits, pushes, and advances state.

**Requirements:** R10, R11, R12, R13, R14, R15, R16, R17

**Dependencies:** U2 (state file must exist from /next)

**Files:**
- Create: `.claude/commands/checkout.md`

**Approach:**
- Command is a no-argument Pi slash command. Step-by-step instructions:
  1. **Read state:** Parse `.state.json`. If no state file, error: "No active problem. Run /next first."
  2. **Locate problem:** From `currentProblem`, find the problem entry and its directory path (`daily/02/<NN>-leetcode-<ID>/`).
  3. **Verify directory:** If directory doesn't exist, error: "Problem directory missing. Run /next to scaffold it."
  4. **Run tests:** Execute `go test -v` in the problem directory. Capture output.
  5. **Check for no tests:** If output shows "no tests to run" or "no test files", report: "No test functions found. Write tests in main_test.go before running /checkout." Exit without grading or state change.
  6. **Grade the code:**
     - Read `main.go` from the problem directory.
     - Apply the CLAUDE.md evaluation criteria: correctness, algorithm choice, time/space complexity, code quality, edge cases, input preservation.
     - If tests pass: assign grade A-F based on code quality and algorithmic merit.
     - If tests fail: assign grade F (per AE4).
     - Output the grade with a brief justification (which criteria drove the grade).
  7. **Update state:** Set problem's `status` to `"done"`, `grade` to the assigned letter. Find the next pending problem in the same category and set `currentProblem` to it. Write state file.
  8. **Regenerate PROGRESS.md:** For each category and problem in state, render a markdown table with checkboxes. Done problems get `[x]` with grade; in-progress gets `[ ]` with "in progress"; pending gets `[ ]`. Write to `daily/02/PROGRESS.md`.
  9. **Commit and push:**
     - `git add daily/02/` (covers the solution files, state, and PROGRESS.md)
     - `git commit -m "✅ [Problem Name] (LeetCode #[ID]) - Grade: [grade]"`
     - `git push`
     - If push fails: warn "Push failed — commit is local. Run 'git push' manually when network is available." Still advance state (the commit is safe locally).

**Patterns to follow:**
- CLAUDE.md grading rubric (A–F scale, 6 assessment areas)
- State file schema from High-Level Technical Design
- Git commit conventions: present-tense, structured prefix

**Test scenarios:**
- Happy path — passing: State points to a problem with completed solution and passing tests → grade assigned (A/B/C/D based on code quality) → PROGRESS.md generated → committed and pushed → state advanced to next problem.
- Happy path — failing tests: Tests fail → grade F → PROGRESS.md shows F → committed → state advanced.
- Edge case — no state file: Error message, exit.
- Edge case — directory missing: Error message suggesting `/next`.
- Edge case — empty test file: "No test functions found" message, exit without grading.
- Edge case — push failure: Commit succeeds, push fails → warning message, state still advances.
- Integration: After `/checkout`, running `/next` scaffolds the next problem in the same category.

**Verification:**
- Set up a problem with passing tests → run `/checkout` → verify PROGRESS.md shows `[x] Done — Grade: [X]`, git log shows the commit, state file has advanced.

---

### U4. Bootstrap state file and repair existing scaffold

**Goal:** Initialize `.state.json` from the README roadmap, detect and preserve existing work, and repair the partial scaffold at `daily/02/00-leetcode-217/`.

**Requirements:** R1, R2, R16, R17

**Dependencies:** U1 (READMe with numeric IDs), U2 (bootstrap is triggered by /next)

**Files:**
- Create: `daily/02/.state.json`
- Create: `daily/02/PROGRESS.md`
- Create: `daily/02/00-leetcode-217/go.mod`
- Create: `daily/02/00-leetcode-217/main_test.go`

**Approach:**
- Bootstrap runs on first invocation of `/next` (detected by missing `.state.json`).
  1. **Parse README:** Extract all 7 category names, display names, and problem rows (numeric ID, name, slug, URL) from `daily/02/README.md`.
  2. **Detect existing work:** Scan `daily/02/` for existing `NN-leetcode-<ID>/` directories. For `00-leetcode-217`, mark it as `"in-progress"` rather than `"pending"`.
  3. **Initialize NN counter:** Set `globalNN` to `max(existing NN) + 1`. If only `00-leetcode-217` exists, `globalNN = 1`.
  4. **Build state:** Populate all categories and problems with status `"pending"`, except the detected existing problem set to `"in-progress"`. Set `currentProblem` to the first in-progress problem, or the first problem in the first category if none.
  5. **Write `.state.json`.**
  6. **Generate initial PROGRESS.md:** All problems as unchecked, except existing work.
  7. **Repair `daily/02/00-leetcode-217/`:**
     - Create `go.mod` with `module github.com/ngdangdat/ds-algorithm/daily/02/00-leetcode-217` and `go 1.26.0`.
     - Create `main_test.go` with `package main` (empty test file, learner writes tests).
     - Do NOT touch existing `main.go` — the solution is preserved.

**Patterns to follow:**
- State file schema from High-Level Technical Design
- Go scaffold from `daily/01/00-leetcode-2942/go.mod`
- README table parsing: category headers `## [emoji] [Name]`, table rows `| [ID] | [Name] | [Difficulty] | [Link](URL) |`

**Test scenarios:**
- Happy path — clean bootstrap: No `.state.json`, no problem directories → bootstrap creates state with all problems `"pending"`, `currentProblem` = first problem in Arrays & Hashing, `globalNN = 0`.
- Edge case — existing work detected: `00-leetcode-217/` exists with `main.go` → that problem is `"in-progress"`, `currentProblem` points to it, `globalNN = 1`. Missing `go.mod` and `main_test.go` are created.
- Edge case — multiple existing directories: `globalNN` = max existing + 1.
- Edge case — PROGRESS.md: Generated with correct checkbox states matching `.state.json`.

**Verification:**
- Delete `.state.json` if it exists → run `/next` → `.state.json` created with correct schema, `00-leetcode-217` marked in-progress, `go.mod` and `main_test.go` present in that directory. `PROGRESS.md` exists with all problems listed.

---

## System-Wide Impact

- **Interaction graph:** The two commands share `.state.json` as their coordination point. No other system components read or write state. The README is read-only by both commands.
- **Error propagation:** `/next` failures (missing README, parse errors) surface immediately with clear messages. `/checkout` failures are handled per the command logic — push failure warns but doesn't block; missing state/directory errors with guidance.
- **State lifecycle risks:** Concurrent edits to `.state.json` (manual + command) are possible but unlikely in a single-developer context. The state file schema uses problem slugs as stable identifiers, so README reordering doesn't break references.
- **API surface parity:** N/A — no API surface.
- **Integration coverage:** The full cycle (`/next` → write code → `/checkout` → `/next` again) is the critical integration path. Cold start and category transitions are the highest-risk integration points.
- **Unchanged invariants:** Existing problem directories outside `daily/02/` are not affected. The older `daily/00/` and `daily/01/` batches are untouched. The `categories/` directory layout is unchanged. CLAUDE.md behavior (strict teacher, no implementation) is unchanged — `/checkout` only reads and grades, never modifies solution code.

---

## Risks & Dependencies

| Risk | Mitigation |
|------|------------|
| README table parsing is fragile — a formatting change breaks `/next` | Command description includes the exact table format contract. Deferred: machine-readable manifest as a future hardening. |
| Pi interprets command instructions differently than intended | Commands use concrete, step-by-step prose with exact file paths and command strings. Test by actually invoking the commands. |
| State file grows large with ~100 problems | JSON is lightweight (~10KB for full ledger). Human-readable for debugging. No performance concern. |
| Learner manually edits state file and breaks JSON | Pi validates JSON on read and reports parse errors with the file path for manual fix. |

---

## Sources & References

- **Origin document:** [docs/brainstorms/2026-05-28-daily-practice-commands.md](docs/brainstorms/2026-05-28-daily-practice-commands.md)
- Related code: `.claude/commands/setup-leetcode.md` (existing Pi command pattern), `daily/01/00-leetcode-2942/` (Go scaffold reference), `daily/02/README.md` (roadmap data source), `CLAUDE.md` (grading criteria)
- State file schema: defined in High-Level Technical Design above
