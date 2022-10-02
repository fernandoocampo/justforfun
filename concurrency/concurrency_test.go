package concurrency_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/justforfun/concurrency"
)

func BenchmarkProcessDataOne(b *testing.B) {
	values := generateValuesOne(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataOne(values)
	}
}

func BenchmarkProcessDataTwo(b *testing.B) {
	values := generateValuesTwo(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataTwo(values)
	}
}

func BenchmarkProcessDataThree(b *testing.B) {
	values := generateValuesThree(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataThree(context.TODO(), values)
	}
}

func BenchmarkProcessDataFour(b *testing.B) {
	values := generateValuesFour(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataFour(context.TODO(), values)
	}
}

func BenchmarkProcessDataFive(b *testing.B) {
	values := generateValuesFive(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataFive(context.TODO(), values)
	}
}

func BenchmarkProcessDataSix(b *testing.B) {
	values := generateValuesSix(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataSix(context.TODO(), values)
	}
}

func BenchmarkProcessDataSeven(b *testing.B) {
	values := generateValuesSeven(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataSeven(context.TODO(), values)
	}
}

func BenchmarkProcessDataEight(b *testing.B) {
	values := generateValuesEight(10)
	for i := 0; i < b.N; i++ {
		_ = concurrency.ProcessDataEight(context.TODO(), values)
	}
}
