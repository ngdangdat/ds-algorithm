
pub struct Solution;

impl Solution {
    fn dfs(
        precourse: usize, graph: &Vec<Vec<usize>>, state: &mut Vec<usize>, path: &mut Vec<i32>
    ) -> bool {
        if state[precourse] == 1 {
            return false;
        }
        if state[precourse] == 2 {
            return true;
        }
        state[precourse] = 1;
        for &course in &graph[precourse] {
            if !Solution::dfs(course, graph, state, path) {
                return false;
            }
        }
        state[precourse] = 2;
        path.push(precourse as i32);
        true
    }
    pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        let n = num_courses as usize;
        let mut graph: Vec<Vec<usize>> = vec![vec![]; n];
        prerequisites.iter().for_each(|p| {
            let course = p[0] as usize;
            let pre = p[1] as usize;
            graph[pre].push(course);
        });
        let mut state: Vec<usize> = vec![0; n];
        let mut path: Vec<i32> = vec![];
        for course in 0..n {
            if !Solution::dfs(course, &graph, &mut state, &mut path) {
                return vec![];
            }
        }
        path.reverse();
        path
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    fn is_valid_order(num_courses: i32, prerequisites: &[Vec<i32>], order: &[i32]) -> bool {
        if order.len() != num_courses as usize {
            return false;
        }

        let mut positions = vec![0; num_courses as usize];
        for (position, &course) in order.iter().enumerate() {
            if course < 0 || course >= num_courses {
                return false;
            }
            positions[course as usize] = position;
        }

        prerequisites
            .iter()
            .all(|pair| positions[pair[1] as usize] < positions[pair[0] as usize])
    }

    #[test]
    fn returns_a_valid_order_for_a_simple_chain() {
        let prerequisites = vec![vec![1, 0]];
        let order = Solution::find_order(2, prerequisites.clone());
        assert!(is_valid_order(2, &prerequisites, &order));
    }

    #[test]
    fn returns_a_valid_order_for_branching_dependencies() {
        let prerequisites = vec![vec![1, 0], vec![2, 0], vec![3, 1], vec![3, 2]];
        let order = Solution::find_order(4, prerequisites.clone());
        assert!(is_valid_order(4, &prerequisites, &order));
    }

    #[test]
    fn handles_a_single_course() {
        assert_eq!(Solution::find_order(1, vec![]), vec![0]);
    }

    #[test]
    fn returns_an_empty_order_for_a_cycle() {
        assert!(Solution::find_order(2, vec![vec![1, 0], vec![0, 1]]).is_empty());
    }
}
