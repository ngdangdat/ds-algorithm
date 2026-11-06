use std::collections::VecDeque;

pub struct Solution;

// 0: not traveled
// 1: visiting
// 2: completed
impl Solution {
    pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
        let n = num_courses as usize;
        let mut indegrees: Vec<i32> = vec![0; n];
        let mut graph: Vec<Vec<usize>> = vec![vec![]; n];
        for p in prerequisites {
            let v = p[0] as usize;
            let u = p[1] as usize;
            indegrees[v] += 1;
            graph[u].push(v);
        }
        let mut queue = VecDeque::new();
        let mut completed: usize = 0;
        for i in 0..n {
            if indegrees[i] == 0 {
                queue.push_back(i);
                completed += 1;
            }
        }
        while let Some(u) = queue.pop_front() {
            for &v in &graph[u] {
                indegrees[v] -= 1;
                if indegrees[v] == 0 {
                    completed += 1;
                    queue.push_back(v);
                }
            }
        }
        completed == n
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn accepts_a_valid_prerequisite_order() {
        assert!(Solution::can_finish(2, vec![vec![1, 0]]));
    }

    #[test]
    fn rejects_a_two_course_cycle() {
        assert!(!Solution::can_finish(2, vec![vec![1, 0], vec![0, 1]]));
    }

    #[test]
    fn accepts_courses_without_prerequisites() {
        assert!(Solution::can_finish(3, vec![]));
    }

    #[test]
    fn rejects_a_self_dependency() {
        assert!(!Solution::can_finish(1, vec![vec![0, 0]]));
    }

    #[test]
    fn accepts_disconnected_course_chains() {
        assert!(Solution::can_finish(
            5,
            vec![vec![1, 0], vec![2, 1], vec![4, 3]]
        ));
    }
}
