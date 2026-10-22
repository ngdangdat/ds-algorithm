use std::{cmp};
struct Solution{}

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut seen: [bool; 500] = [false; 500];
        let mut longest = 0;
        let mut left: usize = 0;
        let schars: Vec<char> = s.chars().collect();
        for right in 0..schars.len() {
            let right_cidx = schars[right] as usize;
            while seen[right_cidx] {
                let left_cidx = schars[left] as usize;
                seen[left_cidx] = false;
                left += 1;
            }
            seen[right_cidx] = true;
            longest = cmp::max((right - left + 1) as i32, longest);
        }
        longest
    }
}

fn run_case(input: &str, expected: i32) {
    let actual = Solution::length_of_longest_substring(input.to_owned());
    assert_eq!(actual, expected, "input: {input:?}");
}

fn main() {
    run_case("", 0);
    run_case("a", 1);
    run_case("abcabcbb", 3);
    run_case("bbbbb", 1);
    run_case("pwwkew", 3);
    run_case("au", 2);
    run_case("dvdf", 3);
    run_case("abba", 2);
    run_case("a b", 3);
    run_case("你好世界你", 4);

    println!("All manual tests passed.");
}
