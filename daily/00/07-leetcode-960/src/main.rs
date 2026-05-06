struct Solution;

impl Solution {
    pub fn min_deletion_size(strs: Vec<String>) -> i32 {
        todo!()
    }
}

fn main() {
    // Example 1
    let strs = vec!["babca".to_string(), "bbazb".to_string()];
    println!("Example 1: {}", Solution::min_deletion_size(strs)); // Expected: 3

    // Example 2
    let strs = vec!["edcba".to_string()];
    println!("Example 2: {}", Solution::min_deletion_size(strs)); // Expected: 4

    // Example 3
    let strs = vec!["ghi".to_string(), "def".to_string(), "abc".to_string()];
    println!("Example 3: {}", Solution::min_deletion_size(strs)); // Expected: 0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example_1() {
        let strs = vec!["babca".to_string(), "bbazb".to_string()];
        assert_eq!(Solution::min_deletion_size(strs), 3);
    }

    #[test]
    fn test_example_2() {
        let strs = vec!["edcba".to_string()];
        assert_eq!(Solution::min_deletion_size(strs), 4);
    }

    #[test]
    fn test_example_3() {
        let strs = vec!["ghi".to_string(), "def".to_string(), "abc".to_string()];
        assert_eq!(Solution::min_deletion_size(strs), 0);
    }
}
