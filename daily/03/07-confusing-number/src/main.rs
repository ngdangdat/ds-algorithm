pub struct Solution;

impl Solution {
    pub fn confusing_number(n: i32) -> bool {
        let mut digits: Vec<i32> = vec![];
        let mut check_n = n;
        while check_n > 0 {
            let mut digit = check_n % 10;
            if [2,3,4,5,7].contains(&digit) {
                return false;
            }
            if digit == 6 {
                digit = 9;
            } else if digit == 9 {
                digit = 6;
            }
            digits.push(digit);
            check_n = check_n / 10;
        }
        let mut changed_num: i32 = 0;
        for (index, digit) in digits.iter().enumerate() {
            let ten_exp: u32 = (digits.len()-index-1) as u32;
            changed_num += (10 as i32).pow(ten_exp)*digit;
        }
        changed_num != n
    }
}

fn main() {
    println!("confusing_number={:?}", Solution::confusing_number(11));
}
