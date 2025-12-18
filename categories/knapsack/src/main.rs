// =============================================================================
// KNAPSACK PROBLEMS - Progressive Learning Path
// =============================================================================
//
// Problem 1: 0/1 Knapsack (Foundation)
// Problem 2: Unbounded Knapsack
// Problem 3: Subset Sum
//
// Start with Problem 1 and work your way up. Each builds on the previous.
// =============================================================================

use std::{cmp, collections::HashMap};

// =============================================================================
// PROBLEM 1: 0/1 Knapsack
// =============================================================================
//
// You have a knapsack with capacity W. There are n items, each with:
// - weight[i]: how much space it takes
// - value[i]: how much it's worth
//
// Each item can only be taken ONCE (0/1 = take or don't take).
// Return the maximum total value that fits in the knapsack.
//
// Example:
//   capacity = 7
//   weights = [1, 3, 4, 5]
//   values  = [1, 4, 5, 7]
//   Answer: 9 (take items with weights 3 and 4, values 4 and 5)
//
// Hint: Think about what decision you make for each item.
//       dp[i][w] = best value using first i items with capacity w
//
pub fn knapsack_01(capacity: i32, weights: Vec<i32>, values: Vec<i32>) -> i32 {
    // TODO: Implement 0/1 knapsack
    // Your code here
    let mut knapsack: HashMap<i32, HashMap<i32, i32>> = HashMap::new();
    for i in 0..weights.len() {
        knapsack.insert(i as i32, HashMap::new());
        for w in weights[i]..=capacity {
            if i == 0 {
                if let Some(bag) = knapsack.get_mut(&(i as i32)) { let _ = bag.insert(w, values[i]); }
                continue;
            }
            let prev_val = knapsack.get(&((i-1) as i32)).and_then(|pbag| pbag.get(&(w-weights[i]))).copied().unwrap_or_default();
            if let Some(bag) = knapsack.get_mut(&(i as i32)) {
                let mut val = values[i];
                if w - weights[i] > 0 {
                    val = cmp::max(val, val + prev_val);
                }

                bag.insert(w, val);
            }
        }
        println!("{:?}", knapsack);
    }
    0
}

// =============================================================================
// PROBLEM 2: Unbounded Knapsack
// =============================================================================
//
// Same as 0/1 knapsack, but you can take each item UNLIMITED times.
//
// Example:
//   capacity = 8
//   weights = [1, 3, 4]
//   values  = [1, 4, 5]
//   Answer: 11 (take weight=3 twice and weight=1 twice: 4+4+1+1+1 = 11)
//   Wait, let me recalculate: weight=3+3+1+1=8, value=4+4+1+1=10
//   Or: weight=4+4=8, value=5+5=10
//   Or: weight=4+3+1=8, value=5+4+1=10
//   Actually: weight=3+3+1+1=8, value=4+4+1+1=10
//   Hmm, answer should be 10 for this example.
//
// Better Example:
//   capacity = 10
//   weights = [2, 3, 5]
//   values  = [6, 8, 12]
//   Answer: 30 (take weight=2 five times, value=6*5=30)
//
// Hint: How does the recurrence change when you can reuse items?
//
pub fn knapsack_unbounded(capacity: i32, weights: Vec<i32>, values: Vec<i32>) -> i32 {
    // TODO: Implement unbounded knapsack
    // Your code here
    0
}

// =============================================================================
// PROBLEM 3: Subset Sum
// =============================================================================
//
// Given a set of integers and a target sum, determine if there exists
// a subset that adds up EXACTLY to the target.
//
// This is a special case of 0/1 knapsack where:
// - capacity = target
// - weight[i] = value[i] = nums[i]
// - We ask: can we achieve exactly capacity?
//
// Example:
//   nums = [3, 34, 4, 12, 5, 2]
//   target = 9
//   Answer: true (3 + 4 + 2 = 9)
//
// Example:
//   nums = [3, 34, 4, 12, 5, 2]
//   target = 30
//   Answer: false (no subset sums to 30)
//
// Hint: dp[i][s] = can we make sum s using first i elements?
//
pub fn subset_sum(nums: Vec<i32>, target: i32) -> bool {
    // TODO: Implement subset sum
    // Your code here
    false
}

fn main() {
    println!("=== Knapsack Problems ===\n");

    // Problem 1: 0/1 Knapsack
    println!("Problem 1: 0/1 Knapsack");
    let result = knapsack_01(7, vec![1, 3, 4, 5], vec![1, 4, 5, 7]);
    println!("  capacity=7, weights=[1,3,4,5], values=[1,4,5,7]");
    println!("  Your answer: {}, Expected: 9\n", result);

    // Problem 2: Unbounded Knapsack
    println!("Problem 2: Unbounded Knapsack");
    let result = knapsack_unbounded(10, vec![2, 3, 5], vec![6, 8, 12]);
    println!("  capacity=10, weights=[2,3,5], values=[6,8,12]");
    println!("  Your answer: {}, Expected: 30\n", result);

    // Problem 3: Subset Sum
    println!("Problem 3: Subset Sum");
    let result = subset_sum(vec![3, 34, 4, 12, 5, 2], 9);
    println!("  nums=[3,34,4,12,5,2], target=9");
    println!("  Your answer: {}, Expected: true\n", result);
}

#[cfg(test)]
mod tests {
    use super::*;

    // =========================================================================
    // Problem 1: 0/1 Knapsack Tests
    // =========================================================================

    #[test]
    fn test_01_basic() {
        assert_eq!(knapsack_01(7, vec![1, 3, 4, 5], vec![1, 4, 5, 7]), 9);
    }

    #[test]
    fn test_01_exact_fit() {
        // All items fit exactly
        assert_eq!(knapsack_01(6, vec![1, 2, 3], vec![6, 10, 12]), 28);
    }

    #[test]
    fn test_01_single_item() {
        assert_eq!(knapsack_01(5, vec![5], vec![10]), 10);
    }

    #[test]
    fn test_01_no_fit() {
        // No item fits
        assert_eq!(knapsack_01(1, vec![2, 3, 4], vec![5, 6, 7]), 0);
    }

    #[test]
    fn test_01_zero_capacity() {
        assert_eq!(knapsack_01(0, vec![1, 2, 3], vec![1, 2, 3]), 0);
    }

    #[test]
    fn test_01_greedy_trap() {
        // Greedy by value/weight ratio would pick wrong
        // Item 0: value=6, weight=3, ratio=2
        // Item 1: value=5, weight=2, ratio=2.5 <- greedy picks this first
        // Item 2: value=5, weight=2, ratio=2.5 <- greedy picks this
        // Greedy gets 10, but optimal is 6+5=11 or 5+5=10
        // Actually let's make a clearer trap:
        // capacity=4, items: (w=3,v=7), (w=2,v=4), (w=2,v=4)
        // Greedy by ratio: (w=3,v=7) ratio=2.33, (w=2,v=4) ratio=2
        // Greedy picks w=3 first, then can't fit w=2, gets 7
        // Optimal: pick both w=2 items, gets 8
        assert_eq!(knapsack_01(4, vec![3, 2, 2], vec![7, 4, 4]), 8);
    }

    // =========================================================================
    // Problem 2: Unbounded Knapsack Tests
    // =========================================================================

    #[test]
    fn test_unbounded_basic() {
        // Take weight=2 five times: 6*5=30
        assert_eq!(knapsack_unbounded(10, vec![2, 3, 5], vec![6, 8, 12]), 30);
    }

    #[test]
    fn test_unbounded_single_item() {
        // Take the item 3 times
        assert_eq!(knapsack_unbounded(9, vec![3], vec![5]), 15);
    }

    #[test]
    fn test_unbounded_exact_fit() {
        // capacity=6, weight=2 can be taken 3 times
        assert_eq!(knapsack_unbounded(6, vec![2], vec![3]), 9);
    }

    #[test]
    fn test_unbounded_mixed() {
        // capacity=5, weights=[1,2,3], values=[1,5,8]
        // Best: take w=2 twice (value=10) and w=1 once (value=1) = 11
        // Or: take w=3 once (8) and w=2 once (5) = 13? No, 3+2=5, value=13
        // Actually w=2 + w=3 = 5, value = 5+8 = 13
        assert_eq!(knapsack_unbounded(5, vec![1, 2, 3], vec![1, 5, 8]), 13);
    }

    #[test]
    fn test_unbounded_zero_capacity() {
        assert_eq!(knapsack_unbounded(0, vec![1, 2, 3], vec![1, 2, 3]), 0);
    }

    // =========================================================================
    // Problem 3: Subset Sum Tests
    // =========================================================================

    #[test]
    fn test_subset_sum_true() {
        assert!(subset_sum(vec![3, 34, 4, 12, 5, 2], 9)); // 3+4+2=9
    }

    #[test]
    fn test_subset_sum_false() {
        assert!(!subset_sum(vec![3, 34, 4, 12, 5, 2], 30));
    }

    #[test]
    fn test_subset_sum_single_element() {
        assert!(subset_sum(vec![5], 5));
        assert!(!subset_sum(vec![5], 3));
    }

    #[test]
    fn test_subset_sum_zero_target() {
        // Empty subset sums to 0
        assert!(subset_sum(vec![1, 2, 3], 0));
    }

    #[test]
    fn test_subset_sum_full_array() {
        assert!(subset_sum(vec![1, 2, 3, 4], 10)); // 1+2+3+4=10
    }

    #[test]
    fn test_subset_sum_impossible() {
        assert!(!subset_sum(vec![2, 4, 6], 5)); // All even, can't make odd
    }
}
