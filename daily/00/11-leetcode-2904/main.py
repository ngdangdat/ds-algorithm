"""
Find shortest beautiful string in provided binary

"""
from typing import Dict, List, Tuple


class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        one_poses = []
        for idx, c in enumerate(s):
            if c == '1':
                one_poses.append(idx)
        if len(one_poses) < k:
            return ""
        if k == 1 and len(one_poses) >= 1:
            return "1"
        left, right = 0, k - 1
        min_len = 10000000  # a random abnormally large number
        result_indexes_map: Dict[int, List[Tuple[int, int]]] = {}
        while left < right and right < len(one_poses):
            diff = one_poses[right] - one_poses[left] + 1
            if diff <= min_len:
                result_indexes_map.setdefault(diff, [])
                result_indexes_map[diff].append((left, right))
                min_len = diff
            left += 1
            right += 1
        if not result_indexes_map:
            return ""
        results = []
        result_indexes = result_indexes_map[min_len]
        for (left, right) in result_indexes:
            index_left, index_right = one_poses[left], one_poses[right]
            idx = index_left
            index_range = one_poses[left:right+1]
            result = ""
            while idx <= index_right:
                if idx in index_range:
                    result += "1"
                else:
                    result += "0"
                idx += 1
            results.append(result)
        return sorted(results)[0]


def main():
    sol = Solution()

    cases = [
        ("100011001", 3, "11001"),
        ("1011", 2, "11"),
        ("001110101101101111", 10, "10101101101111"),
    ]

    for (s, k, expected) in cases:
        res = sol.shortestBeautifulSubstring(s, k)
        output = "PASSED"
        if res != expected:
            output = "FAILED"
        print(f"{output} s={s} k={k} expected={expected} res={res}")


if __name__ == "__main__":
    main()
