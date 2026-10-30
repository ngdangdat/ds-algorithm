pub struct Solution;

impl Solution {
    fn advance_to_next_space(start: usize, s: &Vec<char>) -> usize {
        let mut curr = start+1;
        while curr < s.len() {
            if s[curr] == ' ' {
                break;
            }
            curr += 1;
        }
        curr
    }
    pub fn reverse_words(s: &mut Vec<char>) {
        let (mut left, mut right) = (0, s.len()-1);
        while left < right {
            let tmp = s[left];
            s[left] = s[right];
            s[right] = tmp;
            left += 1;
            right -= 1;
        }
        let (mut left, mut right) = (0, 0);
        //  we stop when both reach last letter
        while right < (s.len()-1) {
            right = Solution::advance_to_next_space(right, s);
            let mut mid = right - 1;
            while left < mid {
                let tmp = s[left];
                s[left] = s[mid];
                s[mid] = tmp;
                left += 1;
                mid -= 1;
            }
            left = right + 1;
        }
    }
}

fn run_case(input: &[char], expected: &[char]) {
    let mut actual = input.to_vec();
    Solution::reverse_words(&mut actual);
    assert_eq!(actual, expected);
}

fn main() {
    run_case(
        &['t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'],
        &['b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'],
    );
    run_case(&['a'], &['a']);
    run_case(&['a', ' ', 'b'], &['b', ' ', 'a']);

    println!("All manual tests passed.");
}
