/// Data structure for maintaining the median of a stream of integers.
pub struct MedianFinder {
    // Add your chosen data structures here.
    nums: Vec<i32>,
    is_sorted: bool,
}

impl MedianFinder {
    pub fn new() -> Self {
        Self {
            nums: Vec::new(),
            is_sorted: false,
        }
    }

    pub fn add_num(&mut self, num: i32) {
        self.nums.push(num);
        self.is_sorted = false;
    }

    pub fn find_median(&mut self) -> f64 {
        if !self.is_sorted {
            self.nums.sort();
            self.is_sorted = true;
        }
        let sz = self.nums.len();
        let mid = sz / 2;
        if sz % 2 == 0 {
            return (f64::from(self.nums[mid-1]) + f64::from(self.nums[mid])) / 2_f64;
        }
        f64::from(self.nums[mid])
    }
}

fn main() {
    // Add manual examples here while developing the solution.
}

#[cfg(test)]
mod tests {
    use super::MedianFinder;

    #[test]
    fn finds_median_for_odd_number_of_values() {
        let mut finder = MedianFinder::new();
        finder.add_num(1);
        finder.add_num(2);
        finder.add_num(3);

        assert!((finder.find_median() - 2.0).abs() < 1e-5);
    }

    #[test]
    fn finds_median_for_even_number_of_values() {
        let mut finder = MedianFinder::new();
        finder.add_num(1);
        finder.add_num(2);

        assert!((finder.find_median() - 1.5).abs() < 1e-5);
    }
}
