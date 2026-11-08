use std::collections::VecDeque;

pub struct Solution;

impl Solution {
    pub fn valid_tree(n: i32, edges: Vec<Vec<i32>>) -> bool {
        let num = n as usize;
        let mut graph: Vec<Vec<usize>> = vec![vec![]; num];
        let mut parents: Vec<Option<usize>> = vec![None; num];
        edges.iter().for_each(|e| {
            let u = e[0] as usize;
            let v = e[1] as usize;
            graph[u].push(v);
            graph[v].push(u);
        });
        let mut stack: VecDeque<usize> = VecDeque::new();
        let mut visited: Vec<bool> = vec![false; num];
        let mut visited_cnt: i32 = 0;
        stack.push_back(0);
        visited[0] = true;
        while let Some(u) = stack.pop_front() {
            visited_cnt += 1;
            for &v in &graph[u] {
                if visited[v] {
                    if parents[u].unwrap() != v {
                        return false;
                    }
                } else {
                    visited[v] = true;
                    parents[v] = Some(u);
                    stack.push_back(v);
                }
            }
        }
        visited_cnt == n
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn accepts_a_connected_acyclic_graph() {
        assert!(Solution::valid_tree(
            5,
            vec![vec![0, 1], vec![0, 2], vec![0, 3], vec![1, 4]]
        ));
    }

    #[test]
    fn rejects_a_graph_with_a_cycle() {
        assert!(!Solution::valid_tree(
            5,
            vec![vec![0, 1], vec![1, 2], vec![2, 3], vec![1, 3], vec![1, 4]]
        ));
    }

    #[test]
    fn rejects_a_disconnected_graph() {
        assert!(!Solution::valid_tree(4, vec![vec![0, 1], vec![2, 3]]));
    }

    #[test]
    fn accepts_a_single_vertex_without_edges() {
        assert!(Solution::valid_tree(1, vec![]));
    }

    #[test]
    fn rejects_a_graph_with_too_many_edges() {
        assert!(!Solution::valid_tree(
            3,
            vec![vec![0, 1], vec![1, 2], vec![0, 2]]
        ));
    }
}
