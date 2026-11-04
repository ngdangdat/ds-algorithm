pub struct Solution;

impl Solution {
    /// Returns the index of any peak element in `nums`.
    ///
    /// A peak is strictly greater than its neighboring elements. Treat values
    /// outside the array as negative infinity.
    ///
    /// Required complexity: O(log n) time and O(1) extra space.
    fn find_peak_index(nums: &Vec<i32>, l: usize, r: usize) -> Option<usize> {
        if l < r {
            let m = l + (r-l)/2;
            if nums[m] < nums[m+1] {
                return Solution::find_peak_index(nums, m+1, r);
            } else {
                return Solution::find_peak_index(nums, l, m);
            }
        }
        if l == r {
            return Some(l);
        }
        None
    }

    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        let index = Solution::find_peak_index(&nums, 0, nums.len()-1);
        match index {
            Some(n) => return n as i32,
            None => return -1,
        }
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    fn is_peak(nums: &[i32], index: i32) -> bool {
        if index < 0 || index as usize >= nums.len() {
            return false;
        }

        let i = index as usize;
        let left = i.checked_sub(1).map_or(i32::MIN, |j| nums[j]);
        let right = nums.get(i + 1).copied().unwrap_or(i32::MIN);
        nums[i] > left && nums[i] > right
    }

    #[test]
    fn accepts_a_peak_from_the_first_example() {
        let nums = vec![1, 2, 3, 1];
        let result = Solution::find_peak_element(nums.clone());
        assert!(is_peak(&nums, result));
    }

    #[test]
    fn accepts_either_peak_from_the_second_example() {
        let nums = vec![1, 2, 1, 3, 5, 6, 4];
        let result = Solution::find_peak_element(nums.clone());
        assert!(is_peak(&nums, result));
    }

    #[test]
    fn handles_a_single_element() {
        let nums = vec![7];
        let result = Solution::find_peak_element(nums.clone());
        assert!(is_peak(&nums, result));
    }

    #[test]
    fn handles_a_two_element_array() {
        let nums = vec![1, 2];
        let result = Solution::find_peak_element(nums.clone());
        assert!(is_peak(&nums, result));
    }
}
