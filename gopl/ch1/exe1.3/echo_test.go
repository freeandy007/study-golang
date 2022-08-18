package echo

import (
	"testing"
	"time"
)

const Times = 1000000

func BenchmarkEcho1(b *testing.B) {
	b.N = Times
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo1()
	}
	b.Log(time.Since(now))
}

func BenchmarkEcho2(b *testing.B) {
	b.N = Times
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo2()
	}
	b.Log(time.Since(now))
}

func BenchmarkEcho3(b *testing.B) {
	b.N = Times
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo3()
	}
	b.Log(time.Since(now))
}
