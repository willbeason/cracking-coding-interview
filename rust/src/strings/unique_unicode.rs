use rand::Rng;
use rand::distr::StandardUniform;
use std::collections::HashSet;
use std::mem::take;

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
    let mut bitarray = [false; std::char::MAX as usize + 1];
    for c in s.chars() {
        if bitarray[c as usize] {
            return false;
        }
        bitarray[c as usize] = true;
    }
    true
}

/// UniqueUnicode3A is a solution that pre-allocates a bit array of all possible Unicode points
/// to use to identify duplicate characters in strings. Resets the bit array each time it is called.
///
/// Not threadsafe.
pub struct UniqueUnicode3A {
    bitarray: Box<[bool; std::char::MAX as usize + 1]>,
}

impl UniqueUnicode3A {
    pub fn new() -> Self {
        Self {
            bitarray: Box::new([false; std::char::MAX as usize + 1]),
        }
    }

    pub fn solution(&mut self, s: &str) -> bool {
        let mut result = true;
        let mut idx = 0;

        for c in s.chars() {
            if self.bitarray[c as usize] {
                result = false;
                break;
            }
            idx += 1;
            self.bitarray[c as usize] = true;
        }

        s.chars().take(idx).for_each(|c| {
            self.bitarray[c as usize] = false;
        });

        result
    }
}

/// UniqueUnicode3B is an AI-suggested modification to UniqueUnicode3B.
/// Additional changes have been manually made.
///
/// Not threadsafe.
pub struct UniqueUnicode3B {
    bitarray: Box<[bool; std::char::MAX as usize + 1]>,
    modified_positions: Vec<usize>,
}

impl UniqueUnicode3B {
    pub fn new() -> Self {
        Self {
            bitarray: Box::new([false; std::char::MAX as usize + 1]),
            modified_positions: Vec::new(),       
        }
    }

    pub fn solution(&mut self, s: &str) -> bool {
        // Keep track of the positions we've modified.
        self.modified_positions.clear();
        let mut result = true;

        for c in s.chars() {
            let pos = c as usize;
            if self.bitarray[pos] {
                result = false;
                break;
            }
            self.bitarray[pos] = true;
            self.modified_positions.push(pos);
        }

        // Reset all modifications.
        self.modified_positions.iter().for_each(|&pos| self.bitarray[pos] = false);

        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_unique_unicode() {
        // assert!(false, "{}", get_random_string(1000).len());

        let mut solution3a = UniqueUnicode3A::new();
        let mut solution3b = UniqueUnicode3B::new();

        let solutions: [(&str, Box<dyn FnMut(&str) -> bool>); 5] = [
            ("1", Box::new(unique_unicode_1)),
            ("2", Box::new(unique_unicode_2a)),
            ("3", Box::new(unique_unicode_3)),
            ("3a", Box::new(|s| solution3a.solution(s))),
            ("3b", Box::new(|s| solution3b.solution(s))),
        ];

        for (s, mut f) in solutions {
            assert!(f(""));
            assert!(f("a"), "Solution {}", s);
            assert!(!f("aa"), "Solution {}", s);
            assert!(f("ab"), "Solution {}", s);
            assert!(f("abc"), "Solution {}", s);
            assert!(!f("abca"), "Solution {}", s);
            assert!(f("abcdefghij"), "Solution {}", s);
            assert!(!f("abcdefghija"), "Solution {}", s);
            assert!(f("abcdefghijklmnopqrstuv"), "Solution {}", s);
            assert!(!f("abcdefghijklmnopqrstuva"), "Solution {}", s);
            assert!(!f("ää"), "Solution {}, f({}) = {}", s, "ää", f("ää"));
            assert!(f("aä"), "Solution {}, f({}) = {}", s, "aä", f("aä"));
            assert!(f("芼罡"), "Solution {}, f({}) = {}", s, "芼罡", f("芼罡"));
            assert!(!f("芼芼"), "Solution {}, f({}) = {}", s, "芼芼", f("芼芼"));
            assert!(f("ㅱ芼"), "Solution {}, f({}) = {}", s, "ㅱ芼", f("ㅱ芼"));
            assert!(!f("ㅱㅱ"), "Solution {}, f({}) = {}", s, "ㅱㅱ", f("ㅱㅱ"));
            // These cases are a mix of hieroglyphs and control characters. They are unlikely to display properly.
            assert!(!f("𓀃𓀃"), "Solution {}, f({}) = {}", s, "𓀃𓀃", f("𓀃𓀃"));
            assert!(f("󠇯𓀃"), "Solution {}, f({}) = {}", s, "󠇯𓀃", f("󠇯𓀃"));
            assert!(!f("󠇯󠇯󠇯"), "Solution {}, f({}) = {}", s, "󠇯󠇯", f("󠇯󠇯"));
        }

        assert!(solution3a.bitarray.iter().all(|&b| !b), "Bitarray not reset");
        assert!(solution3b.bitarray.iter().all(|&b| !b), "Bitarray not reset");
    }
}
