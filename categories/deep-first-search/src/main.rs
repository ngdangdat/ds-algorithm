use std::collections::{HashMap, HashSet};

type Graph = HashMap<i32, Vec<i32>>;

/// Perform DFS traversal starting from the given node.
/// Returns nodes in the order they were visited.
fn dfs(graph: &Graph, start: i32) -> Vec<i32> {
    let mut path: Vec<i32> = Vec::new();
    let mut visited: HashSet<i32> = HashSet::new();
    let mut stack: Vec<i32> = Vec::new();

    match graph.get(&start) {
        Some(v) => stack.extend(v.iter().rev()),
        None => return path,
    }
    visited.insert(start);
    path.push(start);

    while !stack.is_empty() {
        let u = stack.pop().unwrap();
        if visited.contains(&u) { continue; }
        match graph.get(&u) {
            Some(v) => stack.extend(v.iter().rev()),
            None => continue,
        }
        visited.insert(u);
        path.push(u);
    }

    path
}

fn main() {
    // Example graph:
    //     1
    //    / \
    //   2   3
    //  / \
    // 4   5
    let graph: Graph = HashMap::from([
        (1, vec![2, 3]),
        (2, vec![4, 5]),
        (3, vec![]),
        (4, vec![]),
        (5, vec![]),
    ]);

    let result = dfs(&graph, 1);
    println!("DFS traversal: {:?}", result);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_simple_tree() {
        let graph: Graph = HashMap::from([
            (1, vec![2, 3]),
            (2, vec![]),
            (3, vec![]),
        ]);
        let result = dfs(&graph, 1);
        assert_eq!(result, vec![1, 2, 3]);
    }

    #[test]
    fn test_deeper_tree() {
        let graph: Graph = HashMap::from([
            (1, vec![2, 3]),
            (2, vec![4, 5]),
            (3, vec![]),
            (4, vec![]),
            (5, vec![]),
        ]);
        let result = dfs(&graph, 1);
        assert_eq!(result, vec![1, 2, 4, 5, 3]);
    }

    #[test]
    fn test_single_node() {
        let graph: Graph = HashMap::from([(1, vec![])]);
        let result = dfs(&graph, 1);
        assert_eq!(result, vec![1]);
    }

    #[test]
    fn test_graph_with_cycle() {
        // Graph with cycle: 1 -> 2 -> 3 -> 1
        let graph: Graph = HashMap::from([
            (1, vec![2]),
            (2, vec![3]),
            (3, vec![1]),
        ]);
        let result = dfs(&graph, 1);
        assert_eq!(result, vec![1, 2, 3]);
    }

    #[test]
    fn test_disconnected_start() {
        let graph: Graph = HashMap::from([
            (1, vec![2]),
            (2, vec![]),
            (3, vec![4]),
            (4, vec![]),
        ]);
        // Starting from 1 should only visit 1 and 2
        let result = dfs(&graph, 1);
        assert_eq!(result, vec![1, 2]);
    }
}
