---
description: Grade the current LeetCode solution, mark done, and push
---

# /checkout — Grade and Commit Current Solution

Grade the solution for the current problem, regenerate the progress checklist, and push to git.

## Main flow

1. **Read state:** Parse `daily/02/.state.json`. Read `currentProblem.category` and `currentProblem.slug`.

   If `.state.json` does not exist: Report "No active problem. Run /next first." Exit.

2. **Locate problem in state:** Find the problem entry in `categories[currentProblem.category].problems` where `slug` matches `currentProblem.slug`.

   If not found: Report "Current problem not found in state. Check .state.json." Exit.

3. **Build directory path:** `daily/02/<NN>-leetcode-<ID>/` where NN is the problem's `nn` field (zero-padded to 2 digits) and ID is `numericId`.

   If directory does not exist: Report "Problem directory missing. Run /next to scaffold it." Exit.

4. **Run tests:** Execute `go test -v` in the problem directory. Capture all output.

   If output contains "no tests to run" or "no test files": Report "No test functions found. Write tests in `main_test.go` before running /checkout." Exit without grading or state change.

5. **Grade the solution:**
   - Read `main.go` from the problem directory.
   - Apply the strict evaluation criteria from CLAUDE.md:
     - **Correctness:** Does it produce correct results? (Test pass/fail is evidence but not the whole story)
     - **Algorithm Choice:** Is the optimal algorithm used?
     - **Time/Space Complexity:** Are complexities optimal?
     - **Code Quality:** Is it clean, readable, maintainable?
     - **Edge Cases:** Are all edge cases handled?
     - **Input Preservation:** Are inputs preserved when required?
   - Grading scale: A (perfect) / B (good, minor issues) / C (acceptable, notable issues) / D (poor, major flaws) / F (doesn't work or fundamentally wrong).
   - If tests fail: grade is automatically F (per acceptance example AE4).
   - If tests pass: assign grade A-F based on code quality and algorithmic merit.

6. **Report the grade** with a brief justification listing which criteria drove the grade. Example: "Grade: B — correct algorithm (hash map) with O(n) time, but redundant map lookup on line 12 and missing edge case for empty input."

7. **Update state:**
   - Set the problem's `status` to `"done"` and `grade` to the assigned letter.
   - Find the next `"pending"` problem in the same category. Set `currentProblem` to it. If none remain in this category, leave `currentProblem` pointing to the first problem in the next category (by `order`).
   - Write `daily/02/.state.json`.

8. **Regenerate PROGRESS.md:**
   - Read `daily/02/.state.json`.
   - For each category (ordered by `order` field), render a markdown table:
     ```
     ## [emoji] [displayName]

     | # | Status | Problem | Grade |
     |---|--------|---------|-------|
     | [numericId] | ✅/🔄/⬜ | [name] | [grade or —] |
     ```
   - ✅ for `"done"`, 🔄 for `"in-progress"`, ⬜ for `"pending"`.
   - Write to `daily/02/PROGRESS.md`.

9. **Commit and push:**
   - Run `git add daily/02/`.
   - Run `git commit -m "✅ [Problem Name] (LeetCode #[numericId]) - Grade: [grade]"`.
     - Example: `✅ Contains Duplicate (LeetCode #217) - Grade: B`
   - Run `git push`.
   - If push fails: Report "⚠️ Push failed — commit is local. Run `git push` manually when network is available." Do NOT revert state — the problem is still marked done.

10. **Summary output:**
    - Test result (pass/fail) with key output lines
    - Grade and justification
    - Commit hash (short)
    - Next problem: "[name] ([url])" or "Category complete — run /next to advance"

## Error handling

- No state file → "No active problem. Run /next first."
- Directory missing → "Problem directory missing. Run /next to scaffold it."
- No tests → "No test functions found. Write tests in main_test.go before running /checkout."
- Git push failure → warn, continue (state already advanced)
