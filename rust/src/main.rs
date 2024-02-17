mod ten;
fn main() {
    println!("{}", ten::contains_duplicates(vec![1, 1]));
    println!("{}", ten::valid_anagrams("one", "eon"));
}