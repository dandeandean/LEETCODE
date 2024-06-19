pub fn length_of_longest_substring(s: String) -> i32 {
    let mut used: std::collections::HashSet<char> = std::collections::HashSet::new();
    let (mut ret, mut l, mut _r): (usize, usize, usize) = (0, 0, 1);
    for r in 0..s.len() {
        let c = s.chars().nth(r).unwrap();
        while used.contains(&c) {
            used.remove(&s.chars().nth(l).unwrap());
            l += 1;
        }
        used.insert(c);
        ret = std::cmp::max(ret, r - l + 1);
    }
    ret as i32
}
fn main() {
    dbg!(
        length_of_longest_substring("abcabcbb".to_string()),
        length_of_longest_substring("bbbbb".to_string()),
        length_of_longest_substring("pwwkew".to_string()),
        length_of_longest_substring("".to_string()),
        length_of_longest_substring("dvdf".to_string()),
    );
}
