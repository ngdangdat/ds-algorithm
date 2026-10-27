/// Determines whether two strings are exactly one edit apart.
///
/// An edit is one insertion, deletion, or replacement of a character.
pub struct Solution;

impl Solution {
    pub fn is_one_edit_distance(s: String, t: String) -> bool {
        if ((s.len() as i32) - (t.len() as i32)).abs() >= 2 {
            return false;
        }
        if s.eq(&t) {
            return false;
        }
        let (mut si, mut ti, mut diff) = (0, 0, 0);
        let (sl, tl) = (s.len(), t.len());
        let schars: Vec<char> = s.chars().collect();
        let tchars: Vec<char> = t.chars().collect();
        while si < s.len() && ti < t.len() {
            let (cur_s_char, cur_t_char) = (schars[si], tchars[ti]);
            if cur_s_char == cur_t_char {
                si += 1;
                ti += 1;
                continue;
            }
            diff += 1;
            if diff > 1 {
                return false;
            }
            // lesson: check acceptable actions and 
            // separate to cases
            // - replace char in s
            // - delete from s
            // - add to s
            if sl == tl {
                si += 1;
                ti += 1;
            } else if sl > tl {
                si += 1;
            } else {
                ti += 1;
            }
        }
        true
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn insertion() {
        assert!(Solution::is_one_edit_distance("ab".to_owned(), "acb".to_owned()));
    }

    #[test]
    fn replacement() {
        assert!(Solution::is_one_edit_distance("ab".to_owned(), "ac".to_owned()));
    }

    #[test]
    fn deletion() {
        assert!(Solution::is_one_edit_distance("abc".to_owned(), "ab".to_owned()));
    }

    #[test]
    fn identical_strings_are_not_one_edit_apart() {
        assert!(!Solution::is_one_edit_distance("".to_owned(), "".to_owned()));
        assert!(!Solution::is_one_edit_distance("abc".to_owned(), "abc".to_owned()));
    }

    #[test]
    fn more_than_one_edit() {
        assert!(!Solution::is_one_edit_distance("abc".to_owned(), "axy".to_owned()));
    }
}
