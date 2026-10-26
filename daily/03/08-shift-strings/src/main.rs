/// Solver interface for the string-shift problem.
///
/// The algorithm is intentionally left for the student to implement.
pub struct Solution;

impl Solution {
    /// Applies all requested shifts and returns the resulting string.
    ///
    /// Each shift is represented as `[direction, amount]`:
    /// - direction `0`: left shift
    /// - direction `1`: right shift
    pub fn string_shift(s: String, shift: Vec<Vec<i32>>) -> String {
        let mut total_left = 0;
        let mut total_right = 0;
        let sz = s.len() as i32;
        if sz == 0 {
            return s
        }
        for each_shift in &shift {
            let (dir, times) = (each_shift[0], each_shift[1]);
            if dir == 0 {
                total_left += times;
            } else {
                total_right += times;
            }
        }
        let blocks = (total_left-total_right) % sz;
        if blocks == 0 {
            return s
        }
        // if blocks > 0 -> shift left
        let mut char_index = blocks;
        if blocks < 0 {
            char_index = sz + blocks;
        }
        let mut char_arr: Vec<char> = Vec::new();
        let bytes: Vec<char> = s.chars().collect();
        while char_arr.len() < s.len() {
            char_arr.push(bytes[char_index as usize]);
            char_index += 1;
            if char_index == sz {
                char_index = 0;
            }
        }
        String::from_iter(char_arr)
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::Solution;

    // These tests are ignored until Solution::string_shift is implemented.
    #[test]
    fn example_one() {
        assert_eq!(
            Solution::string_shift("abc".to_owned(), vec![vec![0, 1], vec![1, 2]],),
            "cab"
        );
    }

    #[test]
    fn example_two() {
        assert_eq!(
            Solution::string_shift(
                "abcdefg".to_owned(),
                vec![vec![1, 1], vec![1, 1], vec![0, 2], vec![1, 3]],
            ),
            "efgabcd"
        );
    }

    #[test]
    fn no_operations() {
        assert_eq!(Solution::string_shift("abc".to_owned(), vec![]), "abc");
    }

    #[test]
    fn full_rotation() {
        assert_eq!(
            Solution::string_shift("abc".to_owned(), vec![vec![0, 3]]),
            "abc"
        );
    }

    #[test]
    fn amount_larger_than_length() {
        assert_eq!(
            Solution::string_shift("abc".to_owned(), vec![vec![1, 4]]),
            "bca"
        );
    }

    #[test]
    fn single_character() {
        assert_eq!(
            Solution::string_shift("a".to_owned(), vec![vec![0, 100], vec![1, 50]]),
            "a"
        );
    }

    #[test]
    fn empty_string() {
        assert_eq!(Solution::string_shift(String::new(), vec![vec![1, 5]]), "");
    }
}
