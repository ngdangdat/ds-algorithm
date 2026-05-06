from functools import lru_cache
from typing import Dict, List

_DEBUG = False

"""
You are given an integer array nums of length n and an integer numSlots such that 2 * numSlots >= n. There are numSlots slots numbered from 1 to numSlots.

You have to place all n integers into the slots such that each slot contains at most two numbers. The AND sum of a given placement is the sum of the bitwise AND of every number with its respective slot number.

For example, the AND sum of placing the numbers [1, 3] into slot 1 and [4, 6] into slot 2 is equal to (1 AND 1) + (3 AND 1) + (4 AND 2) + (6 AND 2) = 1 + 1 + 0 + 2 = 4.
Return the maximum possible AND sum of nums given numSlots slots.
"""

def log(msg) -> None:
    global _DEBUG
    if _DEBUG is False:
        return 
    print(msg)

class Solution:
    def maximumANDSum(self, nums: List[int], numSlots: int) -> int:
        @lru_cache(None)
        def solve(index, mask) -> int:
            res = 0
            if index == len(nums):
                return 0
            for slot in range(numSlots):
                b = 3 ** slot
                if mask // b % 3 < 2:
                    res = max(res, (nums[index] & (slot + 1)) + solve(index + 1, mask + b))
            return res
        result = solve(0, 0)
        return result


def main():
    global _DEBUG
    _DEBUG = True
    sol = Solution()
    cases = [
        ([1,2,3,4,5,6], 3, 9),
        ([1,3,10,4,7,1], 9, 24),
        ([14,7,9,8,2,4,11,1,9], 8, 40),
        ([8,13,3,15,3,15,2,15,5,7,6], 8, 40),
    ]

    for (nums, numSlots, expected) in cases:
        result = "PASSED"
        got = sol.maximumANDSum(nums, numSlots)
        if got != expected:
            result = "FAILED"
        log(f"{result} nums={nums} numSlots={numSlots} expected={expected} got={got}")

if __name__ == "__main__":
    main()
