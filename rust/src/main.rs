mod ten;
fn main() {
    println!("{}", ten::contains_duplicates(vec![1, 1]));
    println!("{}", ten::valid_anagrams("one", "eon"));
    println!("{:?}", ten::two_sum(vec![1, 7, 3, 9], 10));
}