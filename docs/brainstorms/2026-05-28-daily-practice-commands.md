---
date: 2026-05-28
topic: daily-practice-commands
---

# Daily Practice Commands (/next, /checkout)

## Summary

Two Pi slash commands that automate the daily LeetCode practice flow: `/next` bootstraps a problem scaffold from the roadmap, and `/checkout` grades the solution, marks it done, and pushes.

---

## Problem Frame

The learner works through LeetCode problems daily using a structured roadmap (`daily/02/README.md`) organized into 7 categories with ~100 problems total. The current manual workflow — reading the roadmap to find the next problem, creating directories and Go module boilerplate by hand, self-grading, and manually updating a checklist — adds cognitive overhead before and after each problem. Answering "where was I?" after a gap of days requires scanning the roadmap and mentally reconstructing what's done. The scaffolding and grading steps are rote and automatable; automating them leaves only the core practice (reading the problem, writing the solution) as deliberate work.

---

## Actors

- A1. **Learner**: The developer practicing algorithm problems daily. Calls `/next` to start a problem and `/checkout` after solving it.

---

## Key Flows

- F1. **Start a new problem**
  - **Trigger:** Learner runs `/next`
  - **Actors:** A1
  - **Steps:**
    1. Pi reads the state file to find the current category and position.
    2. Pi reads `daily/02/README.md` to locate the next undone problem in that category.
    3. If the category is exhausted, Pi prompts the learner to advance to the next category.
    4. Pi creates the problem directory `daily/02/<NN>-leetcode-<problemID>/`.
    5. Pi writes `main.go` and `main_test.go` with `package main`.
    6. Pi runs `go mod init` in that directory with an appropriate module path.
    7. Pi outputs the LeetCode problem URL.
    8. Pi updates the state file to point to this problem.
  - **Outcome:** The problem directory is scaffolded and ready for coding. The learner opens the URL and begins solving.
  - **Covered by:** R1, R2, R3, R4, R5, R6, R7, R8, R9

- F2. **Check out a solution**
  - **Trigger:** Learner runs `/checkout`
  - **Actors:** A1
  - **Steps:**
    1. Pi reads the state file to identify the current problem.
    2. Pi runs `go test -v` in the problem directory.
    3. Pi reads the solution code and grades it using the CLAUDE.md evaluation criteria.
    4. Pi marks the problem as `[x] Done` in `PROGRESS.md`, appending the grade.
    5. Pi runs `git add`, commits with a structured message, and pushes.
    6. Pi advances the state file to the next position in the category.
  - **Outcome:** The problem is graded, tracked as complete, and the solution is pushed to the remote.
  - **Covered by:** R10, R11, R12, R13, R14, R15

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
- R7. `/next` initializes a `go.mod` in the directory using `go mod init`, with a module path that distinguishes this problem (e.g., `daily/02/<NN>-leetcode-<problemID>`).
- R8. `/next` outputs the LeetCode problem URL.
- R9. When the current category has no remaining problems, `/next` prompts the learner to advance to the next category instead of bootstrapping a problem.

**`/checkout` command**
- R10. `/checkout` reads the state file to identify which problem to grade.
- R11. `/checkout` runs `go test -v` in the problem directory and reports the test results.
- R12. `/checkout` grades the solution code using the strict evaluation criteria in CLAUDE.md (correctness, algorithm choice, complexity, code quality, edge cases).
- R13. `/checkout` regenerates `PROGRESS.md` from the state file after grading, reflecting the updated completion status and grade.
- R14. `/checkout` commits the solution with a structured message (e.g., `✅ [problem-name] - Grade: [grade]`) and pushes to the remote.
- R15. After a successful checkout, `/checkout` advances the state file to the next position in the category.

**`PROGRESS.md` checklist**
- R16. `PROGRESS.md` is a generated artifact derived from the state file, containing a checkbox list of all problems from `daily/02/README.md` grouped by category with completion status and grades.
- R17. `PROGRESS.md` is created once during initial setup and updated on each `/checkout`.

---

## Acceptance Examples

- AE1. **Covers R5, R6, R7, R8.** Given the state says Arrays & Hashing position 0 (Contains Duplicate, ID 217), when the learner runs `/next`, Pi creates `daily/02/00-leetcode-217/` with `main.go` (`package main`), `main_test.go` (`package main`), and a valid `go.mod`, then outputs `https://leetcode.com/problems/contains-duplicate/`. The state still points to position 0 (problem is in-progress, not done).

- AE2. **Covers R9.** Given the state says Sliding Window position 11 (last problem in category), when the learner runs `/next` after completing it, Pi prompts: "All Sliding Window problems are done. Advance to Linked List?" and does not scaffold a problem.

- AE3. **Covers R11, R12, R13, R14.** Given the state says Arrays & Hashing position 0 and the learner has written a solution in `daily/02/00-leetcode-217/`, when the learner runs `/checkout`, Pi runs `go test -v`, reports pass/fail, grades the code, marks `[x] Done — Grade: B` in PROGRESS.md, commits with `✅ Contains Duplicate - Grade: B`, pushes, and advances state to position 1.

- AE4. **Covers R11, R12.** Given the state says Arrays & Hashing position 0 and the solution has failing tests, when `/checkout` runs, Pi reports the test failures and assigns a grade of F, marks it done, and commits anyway so progress is captured.

---

## Success Criteria

- The learner can go from deciding to practice to having a scaffolded problem ready in under 5 seconds.
- The learner finishes a problem, runs one command, and gets a grade, a checked-off entry, and a pushed commit — no manual bookkeeping.
- The state never gets stuck — `/next` and `/checkout` always know where the learner is in the roadmap.
- Practice consistency is maintained — the learner can resume after any gap without friction from reconstructing position or boilerplate.

---

## Scope Boundaries

- No LeetCode API integration — the learner visits the URL and solves on LeetCode.com.
- No test generation — the learner writes their own `main_test.go`.
- No out-of-order problem selection — the linear traversal is intentional for structured practice.
- No multi-language support — Go only, matching the repo's current direction.
- No automatic problem description scraping.

---

## Key Decisions

- **State file as single source of truth**: The state file is the authoritative record of completion. `PROGRESS.md` is a human-readable rendering generated from state on each `/checkout`, eliminating dual-write synchronization risk.
- **Go exclusively**: The repo already uses Go in `daily/01/`, and `CLAUDE.md` references `go test -v` as the test command.
- **Scaffold on `/next`, not on `/checkout`**: The learner needs the scaffold before they can write code, so bootstrap happens at the start of the problem cycle.

---

## Dependencies / Assumptions

- `daily/02/README.md` is the canonical roadmap and will not change mid-batch (additions/removals would desync the state).
- The remote Git repository is configured and pushable from the working directory.
- Go toolchain is installed and `go mod init` / `go test` work from the command line.

---

## Outstanding Questions

### Resolve Before Planning

- None.

### Deferred to Planning

- [Affects R4, R9][Edge case] Resolved: `/next` is idempotent — if the current problem directory already exists, `/next` reports the existing URL and exits without scaffolding. No state change.
- [Affects R14][Technical] Should the commit message include the LeetCode problem number, the grade, or both? Exact format TBD.
- [Affects R1][Technical] State file format: JSON, TOML, or simple key=value lines? Planning should pick the simplest format Pi can reliably parse.

---

## Deferred / Open Questions

### From 2026-05-28 review

- **"Prevent skipping days" goal has no requirements directly serving it** — Requirements (P2, product-lens, confidence 75)

  Even with the goal reframed to "reduce cognitive overhead," no requirement creates positive reinforcement for consistency — no streak counting, no gentle nudge after N inactive days, no pacing mechanism. The system bets entirely on removing barriers rather than creating pull. Either add a lightweight consistency mechanism or accept that consistency is measured by the learner's own habits and the tool only removes obstacles.

  <!-- dedup-key: section="requirements" title="prevent skipping days goal has no requirements directly serving it" evidence="even with the goal reframed to reduce cognitive overhead no requirement creates positive reinforcement for consistency" -->

- **Directory-scanning as alternative to state file** — Key Decisions (P2, adversarial, confidence 75)

  The current design uses a state file to track position. An alternative — scanning `daily/02/` for existing problem directories to infer completion — eliminates the state file entirely, is self-healing (adapts to manual changes), and handles cold start naturally. Worth evaluating during planning; the identifier-based tracking in R1 already addresses the core desync concern, making this an optimization rather than a redesign.

  <!-- dedup-key: section="key decisions" title="directoryscanning as alternative to state file" evidence="an alternative scanning daily02 for existing problem directories to infer completion eliminates the state file entirely" -->

- **Three failure modes ship correctly and still fail** — Problem Frame, Success Criteria (P2, product-lens, confidence 75)

  Three realistic scenarios where the system works as specified but still breaks the learner's flow: (1) README drifts when the learner edits their own roadmap, desyncing state silently; (2) grading discouragement when consistent C/D/F grades make `/checkout` aversive; (3) commit-timing tension when the learner wants to solve on one device and commit later. Each warrants a defensive mitigation or explicit scope acknowledgment.

  <!-- dedup-key: section="problem frame success criteria" title="three failure modes ship correctly and still fail" evidence="three realistic scenarios where the system works as specified but still breaks the learners flow" -->

- **README parsing fragility** — Requirements (P2, adversarial, confidence 75)

  The doc treats `daily/02/README.md` as a structured data source, but it's a human-authored markdown document with narrative prose, resource links, and duplicate problems across categories. A single formatting change breaks `/next`. Planning should either specify the exact table format contract or consider generating a machine-readable problem manifest (JSON/YAML) from the README as the canonical data source for `/next`.

  <!-- dedup-key: section="requirements" title="readme parsing fragility" evidence="the doc treats daily02readmemd as a structured data source but its a humanauthored markdown document with narrative prose" -->

- **Simpler MVP path** — Key Decisions (P2, product-lens, confidence 75)

  A minimal version delivering the core value: state tracking + print-next + pass/fail grading, deferring full Go module scaffolding, AI grading, git push, and PROGRESS.md. The highest-value component is answering "where am I?" — the full automation can be layered on once the core flow is validated. Worth considering as an incremental build strategy.

  <!-- dedup-key: section="key decisions" title="simpler mvp path" evidence="a minimal version delivering the core value state tracking plus printnext plus passfail grading deferring full go module scaffolding" -->
