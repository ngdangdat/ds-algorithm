struct Solution;

impl Solution {
    pub fn min_deletion_size(strs: Vec<String>) -> i32 {
        let strlen = strs[0].len();
        let mut total = 0;
        for i in 0..strlen {
            for j in 0..strs.len()-1 {
                if strs[j].as_bytes()[i] > strs[j+1].as_bytes()[i] {
                    total += 1;
                    break;
                }
            }
        }
        total
    }
}

fn main() {
    // Example 1
    let strs = vec!["cba".to_string(), "daf".to_string(), "ghi".to_string()];
    println!("Example 1: {}", Solution::min_deletion_size(strs)); // Expected: 1

    // Example 2
    let strs = vec!["a".to_string(), "b".to_string()];
    println!("Example 2: {}", Solution::min_deletion_size(strs)); // Expected: 0

    // Example 3
    let strs = vec!["zyx".to_string(), "wvu".to_string(), "tsr".to_string()];
    println!("Example 3: {}", Solution::min_deletion_size(strs)); // Expected: 3
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example_1() {
        let strs = vec!["cba".to_string(), "daf".to_string(), "ghi".to_string()];
        assert_eq!(Solution::min_deletion_size(strs), 1);
    }

    #[test]
    fn test_example_2() {
        let strs = vec!["a".to_string(), "b".to_string()];
        assert_eq!(Solution::min_deletion_size(strs), 0);
    }

    #[test]
    fn test_example_3() {
        let strs = vec!["zyx".to_string(), "wvu".to_string(), "tsr".to_string()];
        assert_eq!(Solution::min_deletion_size(strs), 3);
    }
}
