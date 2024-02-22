mod ten;
fn main() {
    println!("{}", ten::contains_duplicates(vec![1, 1]));
    println!("{}", ten::valid_anagrams("one", "eon"));
    println!("{:?}", ten::two_sum(vec![1, 7, 3, 9], 10));
    println!("{:?}", ten::group_anagrams(vec![
        "eat".to_string(),
        "tea".to_string(),
        "tan".to_string(),
        "ate".to_string(),
        "nat".to_string(),
        "bat".to_string()]));
    println!("{:?}", ten::top_k_frequencies(vec![1, 1, 1, 2, 2, 2, 3], 1));
    println!("{}", ten::valid_palindrome("race a car".to_string()));
    println!("{}", ten::valid_palindrome("A man, a plan, a canal: Panama".to_string()));
}