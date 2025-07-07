"""
find first and last index

[1, 2, 3, 4, 5, 6, 6, 7] 6
[5, 6]
"""


LOOP = 0


def find_el(nums, start, end, target):
    global LOOP
    LOOP += 1
    if end < start:
        return -1

    if end == start and nums[end] != target:
        return -1

    mid = start + int((end - start) / 2)
    if nums[mid] == target:
        return mid

    if target < nums[mid]:
        return find_el(nums, start, mid - 1, target)
    if target > nums[mid]:
        return find_el(nums, mid + 1, end, target)


if __name__ == "__main__":
    nums = [1, 2, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 9, 10]
    target = 6
    idx = find_el(nums, 0, len(nums) - 1, target)
    if idx == -1:
        print([-1, -1])

    left_idx, right_idx = idx, idx
    while left_idx - 1 >= 0 and nums[left_idx - 1] == target:
        left_idx = find_el(nums, 0, left_idx - 1, target)
    while right_idx + 1 <= len(nums) - 1 and nums[right_idx + 1] == target:
        right_idx = find_el(nums, right_idx + 1, len(nums) - 1, target)

    if left_idx == right_idx:
        print([left_idx, -1])
    else:
        print([left_idx, right_idx])
    print(LOOP, len(nums))
