pub struct Solution;

impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        if n == 1 || trust.len() == 0 {
            return 1;
        }
        let mut trusting: Vec<bool> = vec![false; (n+1) as usize];
        let mut trust_count: Vec<i32> = vec![0; (n+1) as usize];
        for t in trust {
            let (u, v) = (t[0], t[1]);
            trusting[u as usize] = true;
            trust_count[v as usize] += 1;
        }
        for (index, t) in trusting.iter().enumerate() {
            if !*t {
                if trust_count[index] == (n-1) {
                    return index as i32;
                }
                break
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
    fn finds_the_judge_when_everyone_else_trusts_them() {
        assert_eq!(Solution::find_judge(2, vec![vec![1, 2]]), 2);
    }

    #[test]
    fn finds_the_judge_in_a_larger_town() {
        assert_eq!(Solution::find_judge(3, vec![vec![1, 3], vec![2, 3]]), 3);
    }

    #[test]
    fn returns_minus_one_when_the_candidate_is_not_a_judge() {
        assert_eq!(
            Solution::find_judge(3, vec![vec![1, 3], vec![2, 3], vec![3, 1]]),
            -1
        );
    }

    #[test]
    fn identifies_the_only_person_in_a_single_person_town() {
        assert_eq!(Solution::find_judge(1, vec![]), 1);
    }

    #[test]
    fn returns_minus_one_when_no_one_is_trusted_by_everyone_else() {
        assert_eq!(Solution::find_judge(3, vec![vec![1, 2], vec![2, 3]]), -1);
    }
}
