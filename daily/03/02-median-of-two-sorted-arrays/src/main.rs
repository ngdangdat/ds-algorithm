pub struct Solution;

impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let (sz1, sz2) = (nums1.len(), nums2.len());
        let arr_len = sz1 + sz2;
        let (mut cur1, mut cur2) = (0, 0);
        let mut arr: Vec<i32> = vec![i32::MAX; (arr_len / 2) + 1];
        let mut curarr = 0;
        while (cur1 < sz1) || (cur2 < sz2) {
            let mut shifted = false;
            if cur1 < sz1 && nums1[cur1] < arr[curarr] {
                arr[curarr] = nums1[cur1];
                cur1 += 1;
                shifted = true;
            }
            if cur2 < sz2 && nums2[cur2] < arr[curarr] {
                if shifted {
                    cur1 -= 1;
                }
                arr[curarr] = nums2[cur2];
                cur2 += 1;
            }
            curarr += 1;
            if curarr > arr_len/2 {
                break;
            }
        }
        if arr_len % 2 == 0 {
            return (f64::from(arr[arr_len/2]) + f64::from(arr[arr_len/2-1])) / 2_f64;
        }
        arr[arr_len/2] as f64
    }
}

fn run_case(nums1: &[i32], nums2: &[i32], expected: f64) {
    let actual = Solution::find_median_sorted_arrays(nums1.to_vec(), nums2.to_vec());
    assert_eq!(actual, expected);
}

fn main() {
    // Problem: find the median of two sorted arrays without merging them.
    run_case(&[1, 3], &[2], 2.0);
    run_case(&[1, 2], &[3, 4], 2.5);

    println!("All manual tests passed.");
}
