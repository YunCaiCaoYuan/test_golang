package test
// 基准测试

import (
	"fmt"
	"testing"
)

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

/*
func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
 */

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

/*
func BenchmarkBigLen(b *testing.B) {
	big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Len()
	}
}*/

/*
func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		// 每个 goroutine 有属于自己的 bytes.Buffer.
		var buf bytes.Buffer
		for pb.Next() {
			// 循环体在所有 goroutine 中总共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}*/
