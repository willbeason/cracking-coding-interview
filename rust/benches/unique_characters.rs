use chapter1::strings::unique_characters::unique_characters_1;
use criterion::{BenchmarkId, Criterion, criterion_group, criterion_main};

fn bench_unique_characters(c: &mut Criterion) {
    let mut group = c.benchmark_group("unique_characters");
    for i in ["abc"].iter() {
        group.bench_with_input(BenchmarkId::new("unique_characters_1", i), i, |b, i| {
            b.iter(|| unique_characters_1(i))
        });
    }
    group.finish()
}

criterion_group!(benches, bench_unique_characters);
criterion_main!(benches);
