use chapter1::strings::unique_characters::{unique_characters_1, unique_characters_2, unique_characters_3};
use criterion::{BenchmarkId, Criterion, criterion_group, criterion_main};

fn bench_unique_characters(c: &mut Criterion) {
    let mut group = c.benchmark_group("unique_characters");
    
    let data = [
        "a",
        "ab",
        "abcde",
        "abcdefghij",
        "abcdefghijklmnopqrstuv",
    ];
    let functions: [(&str, fn(&str) -> bool); 3] = [
        ("solution-1", unique_characters_1),
        ("solution-2", unique_characters_2),
        ("solution-3", unique_characters_3),
    ];
    for (name, f) in functions {
        for j in data {
            group.bench_with_input(BenchmarkId::new(name, j), j, |b, k| {
                b.iter(|| f(k))
            });
        }
    }

    group.finish()
}

criterion_group!(benches, bench_unique_characters);
criterion_main!(benches);
