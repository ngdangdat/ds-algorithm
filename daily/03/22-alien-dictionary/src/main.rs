use std::{char, collections::VecDeque, cmp};

pub struct Solution;

const N: usize = 125;

impl Solution {
    pub fn alien_order(words: Vec<String>) -> String {
        let mut chars_track: [usize; N] = [0; N];
        words.iter().for_each(|w| {
            w.chars().for_each(|c| {
                chars_track[c as usize] = 1;
            });
        });
        if words.len() == 0 {
            return String::from("");
        } else if words.len() == 1 {
            return ('a'..='z')
                .filter_map(
                    |c| (chars_track[c as usize] == 1)
                    .then(|| c)
                ).collect::<String>();
        }
        let mut graph: Vec<Vec<usize>> = vec![vec![]; N];
        let mut indegrees = vec![0; N];
        for sindex in 1..words.len() {
            let ps: Vec<char> = words[sindex-1].chars().collect();
            let cs: Vec<char> = words[sindex].chars().collect();
            let mut done: bool = false;
            // first different chars
            for char_index in 0..cmp::min(ps.len(), cs.len()) {
                let pc = ps[char_index];
                let cc = cs[char_index];
                if pc != cc {
                    graph[pc as usize].push(cc as usize);
                    indegrees[cc as usize] += 1;
                    done = true;
                    break;
                }
            }
            if !done && ps.len() > cs.len() {
                return String::from("");
            }
        }
        let mut stack: VecDeque<u8> = VecDeque::new();
        for c in 'a'..='z' {
            let index = c as usize;
            if chars_track[index] == 0 {
                continue;
            }
            if indegrees[index] == 0 {
                stack.push_back(index as u8);
            }
        }
        let mut res: Vec<char> = vec![];
        while let Some(u) = stack.pop_front() {
            res.push(u as char);
            let ui = u as usize;
            chars_track[ui] = 0;
            for &v in &graph[ui] {
                indegrees[v] -= 1;
                if indegrees[v] == 0 {
                    stack.push_back(v as u8);
                }
            }
        }
        for c in 'a'..='z' {
            let index = c as usize;
            if chars_track[index] == 1 {
                res.push(c);
            }
            if indegrees[index] > 0 {
                return String::from("");
            }
        }
        res.iter().collect::<String>()
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn derives_the_order_from_the_first_example() {
        let words = ["wrt", "wrf", "er", "ett", "rftt"]
            .into_iter()
            .map(String::from)
            .collect();
        assert_eq!(Solution::alien_order(words), "wertf");
    }

    #[test]
    fn handles_two_ordered_letters() {
        let words = ["z", "x"].into_iter().map(String::from).collect();
        assert_eq!(Solution::alien_order(words), "zx");
    }

    #[test]
    fn rejects_a_cycle() {
        let words = ["z", "x", "z"].into_iter().map(String::from).collect();
        assert_eq!(Solution::alien_order(words), "");
    }

    #[test]
    fn rejects_an_invalid_prefix_order() {
        let words = ["abc", "ab"].into_iter().map(String::from).collect();
        assert_eq!(Solution::alien_order(words), "");
    }

    #[test]
    fn handles_a_single_word() {
        let words = ["abc"].into_iter().map(String::from).collect();
        assert_eq!(Solution::alien_order(words), "abc");
    }
}
