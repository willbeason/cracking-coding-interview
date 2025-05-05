use std::collections::HashSet;

/// Problem 1: Implement an algorithm to determine if all characters in a string are unique.
/// What if you cannot use additional data structures?

/// unique_characters_1 is the autocomplete suggestion for the solution
pub fn unique_characters_1(s: &str) -> bool {
    s.chars().collect::<Vec<_>>().len() == s.chars().collect::<HashSet<_>>().len()
}

/// unique_characters_2 is a second autocomplete suggestion for the solution
pub fn unique_characters_2(s: &str) -> bool {
    let mut bitarray = [0; 256];
    for c in s.chars() {
        bitarray[c as usize] += 1;
    }
    for b in bitarray {
        if b > 1 {
            return false;
        }
    }
    true
}

/// unique_characters_3 is an initial bitarray attempt.
pub fn unique_characters_3(s: &str) -> bool {
    let mut bitarray = [false; 128];
    // .bytes() assumes ASCII, but we can use .bytes() for UTF-8.
    for c in s.bytes() {
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
    use std::collections::HashSet;

    const SOLUTIONS: [(&str, fn(&str) -> bool); 3] = [
        ("1", unique_characters_1),
        ("2", unique_characters_2),
        ("3", unique_characters_3),
    ];

    #[test]
    fn test_unique_characters() {
        for (s, f) in SOLUTIONS {
            assert_eq!(f(""), true, "Solution {}", s);
            assert_eq!(f("a"), true, "Solution {}", s);
            assert_eq!(f("aa"), false, "Solution {}", s);
            assert_eq!(f("ab"), true, "Solution {}", s);
            assert_eq!(f("abc"), true, "Solution {}", s);
            assert_eq!(f("abca"), false, "Solution {}", s);
        }
    }
}

