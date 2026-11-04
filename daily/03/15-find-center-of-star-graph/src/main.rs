use std::collections::HashMap;

pub struct Solution;

const N_UPPER: i32 = 100_000;

impl Solution {
    pub fn find_center(edges: Vec<Vec<i32>>) -> i32 {
        let mut degrees = [0; (N_UPPER+1) as usize];
        let mut max_n = -1;
        for edge in &edges {
            let (u, v) = (edge[0], edge[1]);
            degrees[u as usize] += 1;
            degrees[v as usize] += 1;
            if u > max_n {
                max_n = u;
            }
            if v > max_n {
                max_n = v;
            }
        }
        for n in 1..(max_n+1) {
            if degrees[n as usize] == 0 {
                continue;
            }
            if degrees[n as usize] == (edges.len() as i32) {
                return n;
            }
        }
        -1
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn finds_center_in_the_first_example() {
        assert_eq!(
            Solution::find_center(vec![vec![1, 2], vec![2, 3], vec![4, 2]]),
            2
        );
    }

    #[test]
    fn finds_center_in_the_second_example() {
        assert_eq!(
            Solution::find_center(vec![vec![1, 2], vec![5, 1], vec![1, 3], vec![1, 4]]),
            1
        );
    }

    #[test]
    fn handles_the_smallest_valid_star() {
        assert_eq!(Solution::find_center(vec![vec![1, 2], vec![2, 3]]), 2);
    }
}
