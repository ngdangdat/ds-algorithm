from typing import List

class Solution:
    def isTrionic(self, nums: List[int]) -> bool:
        if len(nums) <= 3:
            return False
        requirements = [
            # inc = True, dec = False
            [True, 0],
            [False, 0],
            [True, 0],
        ]
        prev, curr = 0, 1
        req_index = 0
        while curr < len(nums):
            inc = True
            if nums[curr] < nums[prev]:
                inc = False
            elif nums[curr] == nums[prev]:
                return False
            if inc == requirements[req_index][0]:
                requirements[req_index][1] += 1
            else:
                if requirements[req_index][1] == 0:
                    return False
                req_index += 1
                if req_index >= len(requirements):
                    return False
                requirements[req_index][1] += 1
            prev += 1
            curr += 1

        for req in requirements:
            if req[1] == 0:
                return False
        return True


def main():
    """
    """
    sol = Solution()
    arr1 = [1,3,5,4,2,6]
    ans1 = sol.isTrionic(arr1)
    print(f"ans1=[{ans1}], arr=[{arr1}]")

    arr2 = [2,1,3]
    ans2 = sol.isTrionic(arr2)
    print(f"ans2=[{ans2}], arr=[{arr2}]")

    arr3 = [7,6,4,4]
    ans3 = sol.isTrionic(arr3)
    print(f"ans3=[{ans3}], arr=[{arr3}]")

    arr4 = [6,7,5,1]
    ans4 = sol.isTrionic(arr4)
    print(f"ans4=[{ans4}], arr4=[{arr4}]")


main()
