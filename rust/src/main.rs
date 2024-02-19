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
        "bat".to_string()]))
}