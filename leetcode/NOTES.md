# Leetcode algorithm notes

## 03 - Jump Game II
> https://leetcode.com/problems/jump-game-ii/

- We check in range between (i, i + nums[i]], take the one with max value of nums[i] + i, choose it for the next point
- Jump until i + nums[i] >= nums.size() - 1s -> jump_count

## 04 - H-index
> https://leetcode.com/problems/h-index/

- Sort
