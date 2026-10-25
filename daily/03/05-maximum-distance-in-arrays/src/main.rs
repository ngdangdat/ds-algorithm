use std::cmp;

pub struct Solution;

impl Solution {
    pub fn max_distance(arrays: Vec<Vec<i32>>) -> i32 {
        let mut min_num = arrays[0][0];
        let mut max_num = arrays[0][arrays[0].len()-1];
        let mut res: i32 = 0;
        for idx in 1..arrays.len() {
            let arr = &arrays[idx];
            let curr_min = arr[0];
            let curr_max = arr[arr.len()-1];
            let mr = (curr_max - min_num).abs();
            let ml = (max_num - curr_min).abs();
            if res < mr {
                res = mr;
            }
            if res < ml {
                res = ml;
            }
            min_num = cmp::min(curr_min, min_num);
            max_num = cmp::max(curr_max, max_num);
        }
        // (max_num - min_num).abs()
        res
    }
}

fn run_case(arrays: &[&[i32]], expected: i32) {
    let input = arrays.iter().map(|array| array.to_vec()).collect();
    let actual = Solution::max_distance(input);
    assert_eq!(actual, expected);
}

fn main() {
    // The selected values must come from two different arrays.
    run_case(&[&[1, 2, 3], &[4, 5], &[1, 2, 3]], 4);

    // Single-element arrays.
    run_case(&[&[1], &[2]], 1);

    // Negative and positive values.
    run_case(&[&[-10, -5], &[-3, 0], &[2, 8]], 18);

    // The maximum may use the minimum of a later array.
    run_case(&[&[1, 4], &[0, 5]], 4);

    // All values are equal.
    run_case(&[&[1, 1], &[1, 1]], 0);

    // The smallest value can occur in a later array.
    run_case(&[&[-5, -2], &[-10, -8]], 8);

    println!("All manual tests passed.");
}
