use rand::Rng;
use rand::distr::StandardUniform;
use std::collections::HashSet;

pub fn get_random_string(len: usize) -> String {
    rand::rng()
        .sample_iter::<char, _>(StandardUniform)
        .take(len)
        .collect()
}

/// unique_unicode_1 is the autocomplete suggestion for the solution.
pub fn unique_unicode_1(s: &str) -> bool {
    let mut seen = HashSet::new();
    for c in s.chars() {
        if !seen.insert(c) {
            return false;
        }
    }
    true
}

/// unique_unicode_2 is another autocomplete suggestion for the solution.
/// Fails for non-ascii as len() returns the length in bytes, not the number of
/// characters.
pub fn unique_unicode_2(s: &str) -> bool {
    s.chars().collect::<HashSet<_>>().len() == s.len()
}

/// unique_unicode_2a is a manual correction of the above.
pub fn unique_unicode_2a(s: &str) -> bool {
    s.chars().collect::<HashSet<_>>().len() == s.chars().count()
}

/// unique_unicode_3 is an initial bitarray attempt.
pub fn unique_unicode_3(s: &str) -> bool {
    let mut bitarray = [false; 1024*1024];
    for c in s.chars() {
        if bitarray[c as usize] {
            return false;
        }
        bitarray[c as usize] = true;
    }
    true
}
#[cfg(test)]
mod tests {
    use super::*;

    const SOLUTIONS: [(&str, fn(&str) -> bool); 3] = [
        ("1", unique_unicode_1),
        ("2", unique_unicode_2a),
        ("3", unique_unicode_3),
    ];

    #[test]
    fn test_unique_unicode() {
        for (s, f) in SOLUTIONS {
            assert!(f(""));
            assert!(f("a"), "Solution {}", s);
            assert!(!f("aa"), "Solution {}", s);
            assert!(f("ab"), "Solution {}", s);
            assert!(f("abc"), "Solution {}", s);
            assert!(!f("abca"), "Solution {}", s);
            assert!(!f("ää"), "Solution {}, f({}) = {}", s, "ää", f("ää"));
            assert!(f("aä"), "Solution {}, f({}) = {}", s, "aä", f("aä"));
            assert!(f("芼罡"), "Solution {}, f({}) = {}", s, "芼罡", f("芼罡"));
            assert!(!f("芼芼"), "Solution {}, f({}) = {}", s, "芼芼", f("芼芼"));
            assert!(f("ㅱ芼"), "Solution {}, f({}) = {}", s, "ㅱ芼", f("ㅱ芼"));
            assert!(!f("ㅱㅱ"), "Solution {}, f({}) = {}", s, "ㅱㅱ", f("ㅱㅱ"));
            assert!(!f("𓀃𓀃"), "Solution {}, f({}) = {}", s, "𓀃𓀃", f("𓀃𓀃"));
            assert!(!f("𓀃𓀃"), "Solution {}, f({}) = {}", s, "𓀃𓀃", f("𓀃𓀃"));
            assert!(f("󠇯𓀃"), "Solution {}, f({}) = {}", s, "󠇯𓀃", f("󠇯𓀃"));
            assert!(!f("󠇯󠇯󠇯"), "Solution {}, f({}) = {}", s, "󠇯󠇯", f("󠇯󠇯"));
        }
    }
}
