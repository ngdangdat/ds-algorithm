pub struct Solution;

impl Solution {
    fn first_successful(spell: i32, potions: &Vec<i32>, l: usize, r: usize, success: i64) -> i32 {
        if l <= r {
            let m = l + (r-l)/2;
            let p_m = i64::from(potions[m]) * i64::from(spell);
            if p_m < success {
                return Solution::first_successful(spell, potions, m+1, r, success);
            } else {
                if m == 0 || i64::from(potions[m-1]) * i64::from(spell) < success {
                    return m as i32;
                }
                return Solution::first_successful(spell, potions, l, m-1, success);
            }
        }
        -1
    }

    fn count_successful(spell: i32, potions: &Vec<i32>, success: i64) -> i32 {
        let sz = potions.len() as i32;
        let first_successful = Solution::first_successful(spell, potions, 0, potions.len() - 1, success);
        if first_successful == -1 {
            return 0;
        }
        sz - first_successful
    }
    pub fn successful_pairs(spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
        potions.sort();
        let mut res = vec![];
        for spell in spells {
            res.push(Solution::count_successful(spell, &potions, success));
        }
        res
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn handles_the_first_example() {
        assert_eq!(
            Solution::successful_pairs(vec![5, 1, 3], vec![1, 2, 3, 4, 5], 7),
            vec![4, 0, 3]
        );
    }

    #[test]
    fn handles_the_second_example() {
        assert_eq!(
            Solution::successful_pairs(vec![3, 1, 2], vec![8, 5, 8], 16),
            vec![2, 0, 2]
        );
    }

    #[test]
    fn handles_no_successful_pairs() {
        assert_eq!(
            Solution::successful_pairs(vec![1, 2], vec![1, 1, 1], 10),
            vec![0, 0]
        );
    }
}
