pub struct Solution;

impl Solution {
    pub fn ladder_length(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {
        let _ = (begin_word, end_word, word_list);
        todo!("Implement the shortest word transformation search.");
    }
}

fn run_case(begin_word: &str, end_word: &str, word_list: &[&str], expected: i32) {
    let words = word_list.iter().map(|word| (*word).to_owned()).collect();
    let actual = Solution::ladder_length(begin_word.to_owned(), end_word.to_owned(), words);
    assert_eq!(actual, expected);
}

fn main() {
    // Each transformation changes exactly one letter, and the result counts
    // both the beginWord and endWord.
    run_case("hit", "cog", &["hot", "dot", "dog", "lot", "log", "cog"], 5);

    // The endWord must be present in the dictionary.
    run_case("hit", "cog", &["hot", "dot", "dog", "lot", "log"], 0);

    // A direct one-letter transformation has sequence length two.
    run_case("a", "c", &["b", "c"], 2);

    println!("All manual tests passed.");
}
