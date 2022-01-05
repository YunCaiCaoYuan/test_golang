package main

import "testing"

func benchmarkEnqueue(b *testing.B, q *LKQueue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n ++ {
			q.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, q *LKQueue, size int) {
	for i := 0; i< b.N; i++ {
		for n := 0; n<size; n++ {
			q.Dequeue()
		}
	}
}

func setUp(q *LKQueue, size int) {
	for n := 0; n < size; n ++ {
		q.Enqueue(n)
	}
}

func BenchmarkLKQueue_Enqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	q := NewLKQueue()
	b.StartTimer()
	benchmarkEnqueue(b, q, size)
}
func BenchmarkLKQueue_Enqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	q := NewLKQueue()
	b.StartTimer()
	benchmarkEnqueue(b, q, size)
}
func BenchmarkLKQueue_Enqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	q := NewLKQueue()
	b.StartTimer()
	benchmarkEnqueue(b, q, size)
}
func BenchmarkLKQueue_Enqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	q := NewLKQueue()
	b.StartTimer()
	benchmarkEnqueue(b, q, size)
}

func BenchmarkLKQueue_Dequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	q := NewLKQueue()
	setUp(q, size)
	b.StartTimer()
	benchmarkDequeue(b, q, size)
}
func BenchmarkLKQueue_Dequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	q := NewLKQueue()
	setUp(q, size)
	b.StartTimer()
	benchmarkDequeue(b, q, size)
}
func BenchmarkLKQueue_Dequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	q := NewLKQueue()
	setUp(q, size)
	b.StartTimer()
	benchmarkDequeue(b, q, size)
}
func BenchmarkLKQueue_Dequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	q := NewLKQueue()
	setUp(q, size)
	b.StartTimer()
	benchmarkDequeue(b, q, size)
}
