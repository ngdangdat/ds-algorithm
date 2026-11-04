pub struct Solution;

impl Solution {
    pub fn valid_path(n: i32, edges: Vec<Vec<i32>>, source: i32, destination: i32) -> bool {
        if source == destination {
            return true;
        }
        let mut es_by_v = vec![vec![]; n as usize];
        let mut visited: Vec<bool> = vec![false; n as usize];
        edges.iter().for_each(|edge| {
            let (u, v) = (edge[0], edge[1]);
            es_by_v[u as usize].push(v);
            es_by_v[v as usize].push(u);
        });
        let mut stack: Vec<i32> = vec![];
        stack.push(source);
        visited[source as usize] = true;
        while !stack.is_empty() {
            let u = stack.pop().unwrap();
            for v in &es_by_v[u as usize] {
                if *v == destination {
                    return true;
                }
                if visited[*v as usize] {
                    continue;
                }
                visited[*v as usize] = true;
                stack.push(*v);
            }
        }
        false
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn finds_a_path_in_a_connected_cycle() {
        assert!(Solution::valid_path(
            3,
            vec![vec![0, 1], vec![1, 2], vec![2, 0]],
            0,
            2,
        ));
    }

    #[test]
    fn returns_false_when_vertices_are_disconnected() {
        assert!(!Solution::valid_path(
            6,
            vec![vec![0, 1], vec![0, 2], vec![3, 5], vec![5, 4], vec![4, 3]],
            0,
            5,
        ));
    }

    #[test]
    fn accepts_a_path_from_a_vertex_to_itself() {
        assert!(Solution::valid_path(1, vec![], 0, 0));
    }

    #[test]
    fn handles_a_direct_edge() {
        assert!(Solution::valid_path(2, vec![vec![0, 1]], 0, 1));
    }
}
