use chapter1::strings::unique_unicode::{get_random_string, unique_unicode_1, unique_unicode_2a, unique_unicode_3, UniqueUnicode3A, UniqueUnicode3B};
use criterion::{criterion_group, criterion_main, BatchSize, BenchmarkId, Criterion};

fn bench_unique_unicode(c: &mut Criterion) {
    let mut group = c.benchmark_group("unique_unicode");

    let mut solution3a = UniqueUnicode3A::new();
    let mut solution3b = UniqueUnicode3B::new();
    
    let functions: [(&str, Box<dyn FnMut(&str) -> bool>); 5] = [
        ("solution-1", Box::new(unique_unicode_1)),
        ("solution-2a", Box::new(unique_unicode_2a)),
        ("solution-3", Box::new(unique_unicode_3)),
        ("solution-3a", Box::new(|s| solution3a.solution(s))),
        ("solution-3b", Box::new(|s| solution3b.solution(s))),
    ];

    let lengths = [1, 2, 5, 10, 20, 50, 100, 200, 500, 1000];
    let num_strings = 1024;
    let data: Vec<(usize, Vec<String>)> = lengths
        .iter()
        .map(|&l| (l, (0..num_strings).map(|_| get_random_string(l)).collect()))
        .collect();

    for (name, mut f) in functions {
        for (l, strings) in &data {
            let mut idx = 0;
            group.bench_with_input(BenchmarkId::new(name, l), &l, |b, _| {
                b.iter_batched(|| {
                    let result = strings[idx].as_str();
                    idx = (idx + 1) % num_strings;
                    result
                }, |s| {
                    f(s);
                }, BatchSize::SmallInput);
            });
        }
    }

    group.finish();
}

criterion_group!(benches, bench_unique_unicode);
criterion_main!(benches);
