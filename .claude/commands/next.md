---
description: Scaffold the next LeetCode problem from the daily roadmap
---

# /next — Start Next LeetCode Problem

Scaffold the next problem from the daily/02 roadmap. No arguments — reads state from `.state.json`.

## Bootstrap (first run only)

If `daily/02/.state.json` does not exist, initialize the system:

1. Read `daily/02/README.md` — for each category section (## heading with emoji), parse the table rows. Table format: `| # | Problem Name | Difficulty | [Link](URL) |`. Extract the numeric ID (# column), problem name, and LeetCode URL. Derive the slug from the URL path (last segment without trailing slash).
2. Scan `daily/02/` for existing `NN-leetcode-<ID>/` directories. For any found, mark that problem as `"in-progress"` in state.
3. Set `globalNN` to `max(existing NN) + 1`. If none exist, `globalNN = 0`.
4. Create `daily/02/.state.json` with the schema below.
5. Create `daily/02/PROGRESS.md` — a markdown checklist of all problems grouped by category. Use ⬜ for pending, 🔄 for in-progress, ✅ for done.
6. If a problem directory exists but is missing `go.mod` or `main_test.go`, create the missing files (do not touch existing `main.go`).

**State file schema (daily/02/.state.json):**
```json
{
  "currentProblem": {"category": "<key>", "slug": "<slug>"},
  "globalNN": <number>,
  "categories": {
    "<category-key>": {
      "order": <number>,
      "displayName": "<Display Name>",
      "problems": [
        {
          "slug": "<problem-slug>",
          "numericId": <id>,
          "name": "<Problem Name>",
          "url": "<leetcode-url>",
          "nn": <global-number>,
          "status": "pending|in-progress|done",
          "grade": null|"A"|"B"|"C"|"D"|"F"
        }
      ]
    }
  }
}
```

## Main flow (every run)

1. **Read state:** Parse `daily/02/.state.json`. Read `currentProblem.category` and `currentProblem.slug`.

2. **Find current position in README:** Read `daily/02/README.md`. Find the category section matching the current category name. Find the table row whose URL slug matches `currentProblem.slug`.

3. **Advance to next problem:** Find the next table row in the same category after the current one. Skip rows where the problem's `status` in state is `"done"`.

   **Category exhausted:** If no remaining rows in this category:
   - Read the next category from `categories` (by `order` field).
   - Prompt the user: "All [current category displayName] problems are done. Advance to [next category displayName]?" Use the `ask_user` tool with two options: "Yes, advance" and "No, stay here".
   - If "Yes": Set `currentProblem` to the first problem in the next category. Set that problem's `status` to `"in-progress"`. Continue to step 4.
   - If "No": Report: "No remaining problems in [category]. Remaining categories: [list]." Exit.
   - If no more categories: Report "🎉 All 7 categories complete! No more problems." Exit.

4. **Build directory path:** Read the problem's `nn` (from state) and `numericId` (from README # column). Directory = `daily/02/<NN>-leetcode-<ID>/` where NN is zero-padded to 2 digits (e.g., `00`, `01`).

5. **Idempotency check:** If the directory already exists:
   - Check if `go.mod` is missing → create it with `module github.com/ngdangdat/ds-algorithm/daily/02/<NN>-leetcode-<ID>` and `go 1.26.0` (run `go mod init` then edit go version).
   - Check if `main_test.go` is missing → create it with `package main`.
   - If all files present: Report "Problem already scaffolded at [directory] → [URL]". Exit. No state change.

6. **Scaffold new directory:**
   - Create the directory `daily/02/<NN>-leetcode-<ID>/`.
   - Write `main.go` with content: `package main`
   - Write `main_test.go` with content: `package main`
   - Run `go mod init github.com/ngdangdat/ds-algorithm/daily/02/<NN>-leetcode-<ID>` in the directory.
   - Edit `go.mod` to ensure `go 1.26.0` line is present.

7. **Update state:**
   - Set this problem's `status` to `"in-progress"`.
   - Set `currentProblem` to `{"category": "<key>", "slug": "<slug>"}`.
   - If this is a new problem (not pre-existing), increment `globalNN` by 1.
   - Write `daily/02/.state.json`.

8. **Output:**
   - Print the LeetCode URL from the README.
   - Print: "Scaffolded: `daily/02/<NN>-leetcode-<ID>/`"

## Category order

Categories in traversal order (matching `order` field in state):
1. Arrays & Hashing
2. Two Pointers
3. Stack
4. Binary Search
5. Sliding Window
6. Linked List
7. Trees
