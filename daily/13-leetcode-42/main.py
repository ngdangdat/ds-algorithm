"""
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
"""
from typing import List

_DEBUG = False


def debug(msg) -> None:
    global _DEBUG
    if not _DEBUG:
        return
    print(msg)


class Solution:
    def _calculate(self, left: int, right: int, height: List[int]) -> int:
        return 0
    def trap(self, height: List[int]) -> int:
        if len(height) <= 2:
            return 0
        l, r = 0, 1
        while l < r and r < len(height) - 1:
        

        return 0


def main():
    global _DEBUG
    _DEBUG = True
    cases = [
        ([0,1,0,2,1,0,1,3,2,1,2,1], 6),
    ]
    sol = Solution()
    for heights, expected in cases:
        got = sol.trap(heights)
        res = "PASSED"
        if got != expected:
            res = "FAILED"
        debug(f"{res} heights={heights} got={got} expected={expected}")


if __name__ == "__main__":
    main()
