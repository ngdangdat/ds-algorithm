struct Solution;

impl Solution {
    fn get_descent_periods_for(n: usize) -> i64 {
        (n*(n-1)/2) as i64
    }
    pub fn get_descent_periods(prices: Vec<i32>) -> i64 {
        if prices.len() <= 1 {
            return prices.len() as i64;
        }
        let mut total: i64 = prices.len() as i64;
        let mut start = 0;
        for (i, curr) in prices[1..].iter().enumerate() {
            let idx = i + 1;
            if prices[idx-1] - curr == 1 {
                // case: the last element
                if idx == (prices.len() - 1) {
                    total += Self::get_descent_periods_for(idx - start + 1);
                }
                continue
            }
            if (idx - start) > 1 {
                total += Self::get_descent_periods_for(idx - start);
            }
            start = idx;
        }
        total
    }
}

fn main() {
    let prices: Vec<i32> = vec![3,2,1,4];
    println!("prices: {:?}, result: {}", prices, Solution::get_descent_periods(prices.clone()));
}
