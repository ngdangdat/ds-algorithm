// https://leetcode.com/problems/maximum-profit-from-trading-stocks-with-discounts/description/

use std::collections::HashMap;

struct Solution;

impl Solution {
    fn traverse(start: i32, n: i32, present: Vec<i32>, future: Vec<i32>, graph: HashMap<i32, Vec<i32>>, budget: i32) -> i32 {
        
        todo!()
    }

    pub fn max_profit(n: i32, present: Vec<i32>, future: Vec<i32>, hierarchy: Vec<Vec<i32>>, budget: i32) -> i32 {
        let mut graph: HashMap<i32, Vec<i32>> = HashMap::new();
        for edge in hierarchy {
            graph.entry(edge[0]).or_default().push(edge[1]);
        }
        let max_profits: HashMap<i32, i32> = HashMap::new();
        todo!()
    }
}

fn main() {
    // Example 1: n=2, present=[1,2], future=[4,3], hierarchy=[[1,2]], budget=3
    // Expected: 5
    let result = Solution::max_profit(
        2,
        vec![1, 2],
        vec![4, 3],
        vec![vec![1, 2]],
        3,
    );
    println!("Example 1: {result} (expected: 5)");

    // Example 4: Chain 1->2->3, all buy with discounts
    // Expected: 12
    let result = Solution::max_profit(
        3,
        vec![5, 2, 3],
        vec![8, 5, 6],
        vec![vec![1, 2], vec![2, 3]],
        7,
    );
    println!("Example 4: {result} (expected: 12)");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example_1() {
        // Employee 1 buys at 1, profit 3. Employee 2 gets discount, buys at 1, profit 2.
        // Total cost: 2, total profit: 5
        assert_eq!(
            Solution::max_profit(
                2,
                vec![1, 2],
                vec![4, 3],
                vec![vec![1, 2]],
                3
            ),
            5
        );
    }

    #[test]
    fn test_example_2() {
        // Only Employee 2 buys at 4, sells at 8, profit 4
        assert_eq!(
            Solution::max_profit(
                2,
                vec![3, 4],
                vec![5, 8],
                vec![vec![1, 2]],
                4
            ),
            4
        );
    }

    #[test]
    fn test_example_3() {
        // Employees 1 and 3 buy, total cost 4+6=10 (wait, should be 4+8=12?)
        // Actually: Employee 1 buys at 4 (profit 3), Employee 3 at 8 (but discounted to 4, profit 7)
        // Total cost: 4+4=8, profit: 3+7=10
        assert_eq!(
            Solution::max_profit(
                3,
                vec![4, 6, 8],
                vec![7, 9, 11],
                vec![vec![1, 2], vec![1, 3]],
                10
            ),
            10
        );
    }

    #[test]
    fn test_example_4() {
        // Chain: 1 -> 2 -> 3
        // Employee 1 buys at 5 (profit 3)
        // Employee 2 buys discounted at 1 (profit 4)
        // Employee 3 buys discounted at 1 (profit 5)
        // Total cost: 5+1+1=7, profit: 3+4+5=12
        assert_eq!(
            Solution::max_profit(
                3,
                vec![5, 2, 3],
                vec![8, 5, 6],
                vec![vec![1, 2], vec![2, 3]],
                7
            ),
            12
        );
    }

    #[test]
    fn test_single_employee() {
        // Only CEO, buy at 5, sell at 10, profit 5
        assert_eq!(
            Solution::max_profit(
                1,
                vec![5],
                vec![10],
                vec![],
                5
            ),
            5
        );
    }

    #[test]
    fn test_no_profit_possible() {
        // future <= present, no profit from buying
        assert_eq!(
            Solution::max_profit(
                2,
                vec![10, 10],
                vec![5, 5],
                vec![vec![1, 2]],
                20
            ),
            0
        );
    }

    #[test]
    fn test_budget_too_small() {
        // Budget 0, can't buy anything
        assert_eq!(
            Solution::max_profit(
                2,
                vec![1, 2],
                vec![4, 3],
                vec![vec![1, 2]],
                0
            ),
            0
        );
    }
}
