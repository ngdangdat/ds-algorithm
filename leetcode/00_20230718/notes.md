# Notes
> https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/

## Problem understanding
- obstacles, [i] -> ith obstacle
- Length of longest obstacle course
  - Choose 0 and i inclusive
  - Must include ith obstacle in the course
  - Taller or same height as the one right before it 
- Return array of length of longest obstacle course up to i

## Approaches

### Approach 1

- Iterate, counter i
    - If current index == 0: ans[i]
    - If obs[i] <= obs[i-1]: traverse back until (obs[k] < obs[i] || k == 0) -> ans[i] = ans[k] + 1
    - If obs[i] > obs[i-1]: ans[i] = ans[i - 1] + 1

- O notation: O(n^2)

### Approach 2 -- Leetcode
- Idea greedy approach + binary search
- Create a temporary vector called `lis` Longest Increase Subsequence (LIS)
- Iterate
  - Find the rightmost index (named `idx`) to insert current element to LIS using upper_bound (an upgrade of binary search)
    - If `idx == lis.size()`: insert num to the last
    - If `idx < lis.size()`: insert num at `idx`
- `ans[i] = idx + 1`

## Timer
- 1st: 15
- 2nd: 16
