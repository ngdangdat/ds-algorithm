fn _quicksort_sort(arr: &mut [i32], low: usize, high: usize) -> isize {
    let pivot = arr[low];
    let mut index = high;
    for i in (low+1..=high).into_iter().rev() {
        if arr[i] >= pivot {
            arr.swap(i, index);
            index -= 1;
        }
    }
    arr.swap(low, index);
    index as isize
}

fn _quicksort(arr: &mut [i32], low: usize, high: usize) {
    if low < high {
        let pivot = _quicksort_sort(arr, low, high);
        if pivot >= 1 {
            _quicksort(arr, low, (pivot-1) as usize);
        }
        _quicksort(arr, (pivot+1) as usize, high);
    }
}

fn quicksort(arr: &mut [i32]) {
    if arr.len() > 1 {
        _quicksort(arr, 0, arr.len() - 1);
    }
}

fn main() {
    let mut arr = Vec::from([3,6,8,10,1,2,1]);
    println!("Before: {:?}", arr);
    quicksort(&mut arr);
    println!("After:  {:?}", arr);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_empty_array() {
        let mut arr: [i32; 0] = [];
        quicksort(&mut arr);
        assert_eq!(arr, []);
    }

    #[test]
    fn test_single_element() {
        let mut arr = [42];
        quicksort(&mut arr);
        assert_eq!(arr, [42]);
    }

    #[test]
    fn test_two_elements_sorted() {
        let mut arr = [1, 2];
        quicksort(&mut arr);
        assert_eq!(arr, [1, 2]);
    }

    #[test]
    fn test_two_elements_unsorted() {
        let mut arr = [2, 1];
        quicksort(&mut arr);
        assert_eq!(arr, [1, 2]);
    }

    #[test]
    fn test_already_sorted() {
        let mut arr = [1, 2, 3, 4, 5];
        quicksort(&mut arr);
        assert_eq!(arr, [1, 2, 3, 4, 5]);
    }

    #[test]
    fn test_reverse_sorted() {
        let mut arr = [5, 4, 3, 2, 1];
        quicksort(&mut arr);
        assert_eq!(arr, [1, 2, 3, 4, 5]);
    }

    #[test]
    fn test_random_order() {
        let mut arr = [3, 6, 8, 10, 1, 2, 1];
        quicksort(&mut arr);
        assert_eq!(arr, [1, 1, 2, 3, 6, 8, 10]);
    }

    #[test]
    fn test_duplicates() {
        let mut arr = [5, 2, 5, 2, 5, 2];
        quicksort(&mut arr);
        assert_eq!(arr, [2, 2, 2, 5, 5, 5]);
    }

    #[test]
    fn test_all_same() {
        let mut arr = [7, 7, 7, 7, 7];
        quicksort(&mut arr);
        assert_eq!(arr, [7, 7, 7, 7, 7]);
    }

    #[test]
    fn test_negative_numbers() {
        let mut arr = [-3, -1, -4, -1, -5, -9, -2, -6];
        quicksort(&mut arr);
        assert_eq!(arr, [-9, -6, -5, -4, -3, -2, -1, -1]);
    }

    #[test]
    fn test_mixed_positive_negative() {
        let mut arr = [3, -2, 5, -8, 0, 1, -1];
        quicksort(&mut arr);
        assert_eq!(arr, [-8, -2, -1, 0, 1, 3, 5]);
    }

    #[test]
    fn test_large_array() {
        let mut arr = [64, 34, 25, 12, 22, 11, 90, 45, 33, 21, 50, 40, 30, 20, 10];
        quicksort(&mut arr);
        assert_eq!(arr, [10, 11, 12, 20, 21, 22, 25, 30, 33, 34, 40, 45, 50, 64, 90]);
    }
}
