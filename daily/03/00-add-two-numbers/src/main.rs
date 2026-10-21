#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    pub fn new(val: i32) -> Self {
        Self { val, next: None }
    }
}

pub struct Solution;

impl Solution {
    pub fn add_two_numbers(
        mut l1: Option<Box<ListNode>>,
        mut l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        // this is a virtual node
        let mut head = ListNode{
            next: None,
            val: -1,
        };
        let mut res = &mut head;
        let mut remainder: i32 = 0;
        while l1.is_some() || l2.is_some() {
            let mut sum_val: i32 = remainder;
            if l1.is_some() {
                sum_val += l1.as_ref().unwrap().val;
                l1 = l1.unwrap().next;
            }
            if l2.is_some() {
                sum_val += l2.as_ref().unwrap().val;
                l2 = l2.unwrap().next;
            }
            remainder = sum_val / 10;
            let node = ListNode{
                next: None,
                val: sum_val%10,
            };
            res.next = Some(Box::from(node));
            res = res.next.as_mut().unwrap();
        }
        if remainder > 0 {
            res.next = Some(Box::from(ListNode{
                next: None,
                val: 1,
            }));
        }
        head.next
    }
}

fn list(values: &[i32]) -> Option<Box<ListNode>> {
    values
        .iter()
        .rev()
        .fold(None, |next, &val| Some(Box::new(ListNode { val, next })))
}

fn run_case(left: &[i32], right: &[i32], expected: &[i32]) {
    let actual = Solution::add_two_numbers(list(left), list(right));
    assert_eq!(actual, list(expected));
}

fn main() {
    // Digits are stored in reverse order: 342 + 465 = 807.
    run_case(&[2, 4, 3], &[5, 6, 4], &[7, 0, 8]);

    // Zero values.
    run_case(&[0], &[0], &[0]);

    // Carry that creates an additional digit: 999 + 1 = 1000.
    run_case(&[9, 9, 9], &[1], &[0, 0, 0, 1]);

    // Different lengths with carries: 99 + 1 = 100.
    run_case(&[9, 9], &[1], &[0, 0, 1]);

    println!("All manual tests passed.");
}
