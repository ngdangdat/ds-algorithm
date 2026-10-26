pub struct Solution;

impl Solution {
    fn merge(nums: &mut Vec<i32>, ll: usize, lr: usize, rl: usize, rr: usize) {
        let mut li = ll;
        let mut ri = rl;
        let mut insert_index = 0;
        let mut tmp_vec = vec![-1; rr-ll+1];
        while li <= lr && ri <= rr {
            let mut to_be_inserted = nums[li];
            if nums[li] <= nums[ri] {
                li += 1;
            } else {
                to_be_inserted = nums[ri];
                ri += 1;
            }
            tmp_vec[insert_index] = to_be_inserted;
            insert_index += 1;
        }
        while li <= lr {
            tmp_vec[insert_index] = nums[li];
            insert_index += 1;
            li += 1;
        }
        while ri <= rr {
            tmp_vec[insert_index] = nums[ri];
            insert_index += 1;
            ri += 1;
        }
        for (index, n) in tmp_vec.iter().enumerate() {
            nums[ll+index] = *n;
        }
    }
    fn merge_sort(nums: &mut Vec<i32>, l: usize, r: usize) {
        if r <= l { return; }
        if r > l && r - l == 1 {
            if nums[l] > nums[r] {
                // swap without a tmp variable
                nums[r] = nums[l] + nums[r];
                nums[l] = nums[r] - nums[l];
                nums[r] = nums[r] - nums[l];
            }
            return;
        }
        let mid = l + (r-l)/2;
        Solution::merge_sort(nums, l, mid-1);
        Solution::merge_sort(nums, mid, r);
        Solution::merge(nums, l, mid-1, mid, r);
    }
    pub fn wiggle_sort(nums: &mut Vec<i32>) {
        Solution::merge_sort(nums, 0, nums.len()-1);
        let (mut left, mut right) = (2, 1);
        while left < nums.len() && right < nums.len() {
            if left == right {
                break;
            }
            let tmp = nums[left];
            nums[left] = nums[right];
            nums[right] = tmp;
            if right < nums.len() {
                right = right + 2;
            }
            if left < nums.len() {
                left = right + 1;
            }
        }
    }
}

fn assert_wiggle_sort(input: &[i32]) {
    let mut nums = input.to_vec();
    let mut expected_values = input.to_vec();

    Solution::wiggle_sort(&mut nums);

    expected_values.sort_unstable();
    let mut actual_values = nums.clone();
    actual_values.sort_unstable();

    assert_eq!(actual_values, expected_values, "output must preserve all values");

    for index in 1..nums.len() {
        if index % 2 == 1 {
            assert!(
                nums[index - 1] <= nums[index],
                "expected nums[{}] <= nums[{}], got {:?}",
                index - 1,
                index,
                nums
            );
        } else {
            assert!(
                nums[index - 1] >= nums[index],
                "expected nums[{}] >= nums[{}], got {:?}",
                index - 1,
                index,
                nums
            );
        }
    }
}

fn main() {
    // Example 1: multiple valid outputs are accepted.
    assert_wiggle_sort(&[3, 5, 2, 1, 6, 4]);

    // Example 2: duplicate values.
    assert_wiggle_sort(&[6, 6, 5, 6, 3, 8]);

    // Already wiggle-sorted input.
    assert_wiggle_sort(&[1, 4, 2, 5, 3]);

    // Two values.
    assert_wiggle_sort(&[2, 1]);

    // Single value.
    assert_wiggle_sort(&[7]);

    // Negative values and duplicates.
    assert_wiggle_sort(&[-3, -1, -2, -2, 0]);

    println!("All manual tests passed.");
}
