from typing import List

# Phases: 0=increasing, 1=decreasing, 2=increasing
PHASES = [True, False, True]


class Solution:
    @staticmethod
    def _max_trionic_sum(turns: List[int], nums: List[int]) -> int:
        # Greedy extend left arm from peak backward
        left_sums = []
        curr_sum = 0
        for i in range(turns[1] - 1, turns[0] - 1, -1):
            curr_sum += nums[i]
            left_sums.append(curr_sum)

        # Greedy extend right arm from trough forward
        right_sums = []
        curr_sum = 0
        for i in range(turns[2] + 1, turns[3] + 1):
            curr_sum += nums[i]
            right_sums.append(curr_sum)

        # Mandatory core: peak to trough
        mid = sum(nums[turns[1]:turns[2] + 1])
        return max(left_sums) + mid + max(right_sums)

    def maxSumTrionic(self, nums: List[int]) -> int:
        phase = 0
        phase_counts = [0, 0, 0]
        results = []
        pi, ci = 0, 1
        # turns: [start_of_inc, peak, trough, end_of_inc]
        turns = [0, -1, -1, -1]
        turn_idx = 1

        def _try_record(end: int):
            if phase == 2 and phase_counts[2] > 0:
                turns[3] = end
                results.append(self._max_trionic_sum(turns, nums))

        def _reset(start: int):
            nonlocal phase, phase_counts, turns, turn_idx
            phase = 0
            phase_counts = [0, 0, 0]
            turns = [start, -1, -1, -1]
            turn_idx = 1

        def _rewind_to_trough():
            nonlocal pi, ci
            trough = turns[2]
            pi, ci = trough, trough + 1
            _reset(trough)

        while ci < len(nums):
            prev, curr = nums[pi], nums[ci]

            if prev == curr:
                _try_record(pi)
                _reset(ci)
                ci += 1
                pi += 1
                continue

            going_up = curr > prev

            if going_up == PHASES[phase]:
                # Continue current phase
                phase_counts[phase] += 1
            elif phase_counts[phase] > 0:
                # Direction changed with progress — transition
                phase += 1
                if phase >= len(PHASES):
                    # Full trionic found
                    turns[3] = pi
                    results.append(self._max_trionic_sum(turns, nums))
                    _rewind_to_trough()
                    continue
                else:
                    turns[turn_idx] = pi
                    turn_idx += 1
                    phase_counts[phase] += 1
            else:
                # Direction changed with no progress — restart
                _reset(ci)

            pi += 1
            ci += 1

        _try_record(len(nums) - 1)

        return max(results)


def main():
    sol = Solution()
    arr1 = [0, -2, -1, -3, 0, 2, -1]
    ans1 = sol.maxSumTrionic(arr1)
    print(f"arr1={arr1}, ans1={ans1}")

    arr2 = [1, 4, 2, 7]
    ans2 = sol.maxSumTrionic(arr2)
    print(f"arr2={arr2}, ans2={ans2}")

    arr3 = [2, 993, -791, -635, -569]
    ans3 = sol.maxSumTrionic(arr3)
    print(f"arr3={arr3}, ans3={ans3}")

    arr4 = [1, 4, 2, 2, 3, 1, 2]
    ans4 = sol.maxSumTrionic(arr4)
    print(f"arr4={arr4}, ans4={ans4}")

    arr5 = [-533, 224, -324, 251, 231, 479]
    ans5 = sol.maxSumTrionic(arr5)
    print(f"arr5={arr5}, ans5={ans5}")

    arr6 = [0, -2, -1, -3, 0, 0, 2, -1]
    ans6 = sol.maxSumTrionic(arr6)
    print(f"arr6={arr6}, ans6={ans6}")


main()
