pub struct Solution;

impl Solution {
    pub fn find_degrees(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        matrix.iter().map(|block| block.into_iter().sum()).collect()
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn finds_degrees_in_a_complete_three_vertex_graph() {
        assert_eq!(
            Solution::find_degrees(vec![vec![0, 1, 1], vec![1, 0, 1], vec![1, 1, 0]]),
            vec![2, 2, 2]
        );
    }

    #[test]
    fn handles_an_isolated_vertex() {
        assert_eq!(
            Solution::find_degrees(vec![vec![0, 1, 0], vec![1, 0, 0], vec![0, 0, 0]]),
            vec![1, 1, 0]
        );
    }

    #[test]
    fn handles_a_single_vertex() {
        assert_eq!(Solution::find_degrees(vec![vec![0]]), vec![0]);
    }
}
