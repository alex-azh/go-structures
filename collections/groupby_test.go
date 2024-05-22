// goos: windows
// goarch: amd64
// pkg: others/unsafe/field
// cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
// === RUN   Benchmark
// Benchmark
// Benchmark-16
//  2470384               482.9 ns/op           456 B/op         11 allocs/op

package collections_test

func Benchmark(b *testing.B) {
	first := MyStruct{ID: 0, Link: "1"}
	list := []*MyStruct{
		&first,
		{ID: 0, Link: "2"},
		{ID: 0, Link: "3"},
		{ID: 2, Link: "4"},
		{ID: 3, Link: "5"},
		{ID: 3, Link: "6"},
		{ID: 2, Link: "7"},
		{ID: 5, Link: "8"},
		{ID: 6, Link: "9"},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for range b.N {
		_ = GroupBy(&first, &first.ID, list)
	}
}