---
argument-hint: [leetcode-url] [programming language]
description: Setup Leetcode problem in a specific language
---

Help setup Leetcode problem at $1 using programming language $2, default to Rust
- Folder to create in: `daily/[next-index]-leetcode-[leetcode-problem-index]`
    - `next-index`: increasing index number
    - `leetcode-problem-index`: should be fetched from the URL
- Check existing folder for the structure of the language
- If user requests a new programming language, make sure these points
  - Easy to run
  - Run instruction is in place
  - To be created
    - Test cases
    - Main entrypoint
