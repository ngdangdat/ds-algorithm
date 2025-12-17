/// Returns the longest common substring between two strings.
/// If there are multiple substrings of the same maximum length, return any one of them.
/// If there is no common substring, return an empty string.
pub fn longest_common_substring(s1: &str, s2: &str) -> String {
    let mut cells: Vec<Vec<i32>>= Vec::new();
    let char_s1: Vec<char> = s1.chars().collect();
    let char_s2: Vec<char> = s2.chars().collect();
    let mut max: (i32, i32, i32) = (0, 0, 0);

    for i in 0..char_s1.len() {
        cells.push(Vec::new());
        for j in 0..char_s2.len() {
            let mut val: i32 = 0;
            if char_s1[i] == char_s2[j] {
                if i == 0 || j == 0 {
                    val = 1;
                } else {
                    val = cells[i-1][j-1] + 1;
                }
            }
            cells[i].push(val);
            if val > max.2 {
                max = (i as i32, j as i32, val);
            }
        }
    }
    let (mut ci, _, mut cm) = max;
    let mut res_chars: Vec<char> = Vec::new();
    while cm > 0 {
        res_chars.push(char_s1[ci as usize]);
        cm -= 1;
        ci -= 1;
    }
    res_chars.reverse();
    res_chars.iter().cloned().collect::<String>()
}

fn main() {
    let s1 = "abcdxyz";
    let s2 = "xyzabcd";
    let result = longest_common_substring(s1, s2);
    println!("LCS of '{}' and '{}' is '{}'", s1, s2, result);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_common_substring_at_end() {
        // "abcd" appears at end of s1, start of s2
        let result = longest_common_substring("abcdxyz", "xyzabcd");
        assert!(result == "abcd" || result == "xyz"); // both are length 4 and 3, abcd is longer
    }

    #[test]
    fn test_common_substring_middle() {
        let result = longest_common_substring("GeeksforGeeks", "GeeksQuiz");
        assert_eq!(result, "Geeks");
    }

    #[test]
    fn test_empty_first_string() {
        let result = longest_common_substring("", "abc");
        assert_eq!(result, "");
    }

    #[test]
    fn test_empty_second_string() {
        let result = longest_common_substring("abc", "");
        assert_eq!(result, "");
    }

    #[test]
    fn test_both_empty() {
        let result = longest_common_substring("", "");
        assert_eq!(result, "");
    }

    #[test]
    fn test_no_common_substring() {
        let result = longest_common_substring("abc", "xyz");
        assert_eq!(result, "");
    }

    #[test]
    fn test_entire_string_is_substring() {
        let result = longest_common_substring("abc", "abc");
        assert_eq!(result, "abc");
    }

    #[test]
    fn test_one_is_substring_of_other() {
        let result = longest_common_substring("abc", "xabcy");
        assert_eq!(result, "abc");
    }

    #[test]
    fn test_single_char_match() {
        let result = longest_common_substring("a", "a");
        assert_eq!(result, "a");
    }

    #[test]
    fn test_single_char_no_match() {
        let result = longest_common_substring("a", "b");
        assert_eq!(result, "");
    }

    #[test]
    fn test_repeated_characters() {
        let result = longest_common_substring("aaaa", "aa");
        assert_eq!(result, "aa");
    }

    #[test]
    fn test_case_sensitive() {
        let result = longest_common_substring("ABC", "abc");
        assert_eq!(result, ""); // case sensitive, no match
    }
}
