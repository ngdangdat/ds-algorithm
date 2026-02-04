from typing import List

class Solution:
    @staticmethod
    def _get_next() -> int:
        return 0
    def maxSumTrionic(self, nums: List[int]) -> int:
        # inc = True, dec = False
        reqs= [True, False, True]
        currs = [0, 0, 0]
        reqi = 0
        sums = []
        pi, ci = 0, 1
        curr_sum = nums[0]
        while ci < len(nums):
            prev, curr = nums[pi], nums[ci]
            inc = True
            if prev > curr:
                inc = False
            print(f"reqi={reqi}, curr_sum={curr_sum}, prev={prev}, curr={curr}, inc={inc}, reqs[ri]={reqs[reqi]}")
            if prev == curr:
                currs = [0, 0, 0]
                reqi = 0
                curr_sum = curr
                ci += 1
                pi += 1
                continue
            if inc == reqs[reqi]:
                currs[reqi] += 1
                curr_sum += curr
            else:
                if currs[reqi] > 0:
                    # a turn is made here
                    print(f"reqi={reqi}, inc={inc}, turn")
                    reqi += 1
                    if reqi >= len(reqs):
                        # reset here
                        reqi = 0
                        currs = [0, 0, 0]
                        sums.append(curr_sum - curr)
                        curr_sum = curr
                    else:
                        curr_sum += curr
                        currs[reqi] += 1
                else:
                    reqi = 0
                    currs = [0, 0, 0]
                    sums.append(curr_sum - curr)
                    curr_sum = curr

            pi += 1
            ci += 1

        if reqi == 2 and currs[reqi] > 0:
            sums.append(curr_sum)

        return max(sums)


def main():
    sol = Solution()
    arr1 = [0,-2,-1,-3,0,2,-1]
    ans1 = sol.maxSumTrionic(arr1)
    print(f"arr1={arr1}, ans1={ans1}")

    arr2 = [1, 4, 2, 7]
    ans2 = sol.maxSumTrionic(arr2)
    print(f"arr2={arr2}, ans2={ans2}")


main()
