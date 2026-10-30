use std::collections::HashMap;
/// Data structure for adding integers and checking whether any pair sums to a target.
pub struct TwoSum {
    nums: Vec<i32>,
    freq: HashMap<i32, i32>,
}

impl TwoSum {
    pub fn new() -> Self {
        Self {
            nums: Vec::new(),
            freq: HashMap::new(),
        }
    }

    pub fn add(&mut self, number: i32) {
        self.nums.push(number);
        *(self.freq.entry(number).or_insert(0)) += 1;
    }

    pub fn find(&self, value: i32) -> bool {
        for num in &self.nums {
            let comp = value - *num;
            let freq = *(self.freq.get(&comp).unwrap_or(&0));
            if comp == *num {
                if freq > 1 {
                    return true;
                }
            } else {
                if freq > 0 {
                    return true;
                }
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
    use super::TwoSum;

    #[test]
    fn finds_a_pair_added_to_the_data_structure() {
        let two_sum = TwoSum::new();
        two_sum.add(1);
        two_sum.add(3);
        two_sum.add(5);

        assert!(two_sum.find(4));
        assert!(two_sum.find(6));
        assert!(!two_sum.find(7));
    }

    #[test]
    fn allows_using_the_same_value_twice_when_added_twice() {
        let two_sum = TwoSum::new();
        two_sum.add(3);

        assert!(!two_sum.find(6));

        two_sum.add(3);
        assert!(two_sum.find(6));
    }

    #[test]
    fn handles_negative_numbers_and_zero() {
        let two_sum = TwoSum::new();
        two_sum.add(-4);
        two_sum.add(0);
        two_sum.add(4);

        assert!(two_sum.find(0));
        assert!(two_sum.find(-4));
        assert!(!two_sum.find(8));
    }

    #[test]
    fn returns_false_when_the_data_structure_is_empty() {
        let two_sum = TwoSum::new();

        assert!(!two_sum.find(0));
    }
}
