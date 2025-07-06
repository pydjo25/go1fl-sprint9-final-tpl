package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateRandomElements(t *testing.T) {
	dataTests := []struct {
		name string
		size int
	}{
		{"Zero size", 0},
		{"Size 1", 1},
		{"Small size", 100},
		{"Large size", 10000},
	}

	for _, dt := range dataTests {
		t.Run(dt.name, func(t *testing.T) {
			slice := generateRandomElements(dt.size)

			if len(slice) != dt.size {
				t.Errorf("Expected length %d, got %d", dt.size, len(slice))
			}

			if dt.size > 0 {
				for i, v := range slice {
					if v < 0 {
						t.Errorf("Negative value at index %d: %d", i, v)
					}
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	dataTests := []struct {
		name   string
		slice  []int
		expect int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{5}, 5},
		{"All same", []int{3, 3, 3, 3}, 3},
		{"Positive numbers", []int{1, 2, 5, 4, 3}, 5},
		{"With negatives", []int{-1, -5, -3, -2}, -1},
		{"First is max", []int{10, 2, 5}, 10},
		{"Last is max", []int{1, 2, 15}, 15},
	}

	for _, dt := range dataTests {
		t.Run(dt.name, func(t *testing.T) {
			result := maximum(dt.slice)
			if result != dt.expect {
				t.Errorf("Expected %d, got %d", dt.expect, result)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {

	dataTests := []struct {
		name   string
		slice  []int
		expect int
	}{
		{
			name:   "Empty slice",
			slice:  []int{},
			expect: 0,
		},
		{
			name:   "Single element",
			slice:  []int{42},
			expect: 42,
		},
		{
			name:   "Size less than CHUNKS",
			slice:  []int{3, 1, 4, 1, 5, 9, 2}, // 7 элементов
			expect: 9,
		},
		{
			name:   "Exact multiple of CHUNKS",
			slice:  make([]int, 16), // 16 элементов
			expect: 15,
		},
		{
			name:   "Not divisible by CHUNKS",
			slice:  make([]int, 19), // 19 элементов
			expect: 18,
		},
		{
			name:   "Max in first chunk",
			slice:  make([]int, 24),
			expect: 999,
		},
		{
			name:   "Max in last chunk",
			slice:  make([]int, 24),
			expect: 999,
		},
		{
			name:   "Max in middle chunk",
			slice:  make([]int, 24),
			expect: 999,
		},
		{
			name:   "All negative values",
			slice:  make([]int, 16),
			expect: -1,
		},
		{
			name:   "Large random data",
			slice:  make([]int, 1000),
			expect: 0, // Будет пересчитано
		},
	}

	for i := range dataTests[3].slice {
		dataTests[3].slice[i] = i
	}
	for i := range dataTests[4].slice {
		dataTests[4].slice[i] = i
	}
	dataTests[5].slice[0] = 999
	dataTests[6].slice[23] = 999
	dataTests[7].slice[10] = 999
	for i := range dataTests[8].slice {
		dataTests[8].slice[i] = -i - 10
	}
	dataTests[8].slice[5] = -1

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxVal := math.MinInt
	for i := range dataTests[9].slice {
		val := r.Intn(1000000)
		dataTests[9].slice[i] = val
		if val > maxVal {
			maxVal = val
		}
	}
	dataTests[9].expect = maxVal

	for _, dt := range dataTests {
		t.Run(dt.name, func(t *testing.T) {
			result := maxChunks(dt.slice)
			if result != dt.expect {
				t.Errorf("Expected %d, got %d", dt.expect, result)
			}
		})
	}
}

func BenchmarkMaximum(b *testing.B) {
	data := generateRandomElements(1_000_000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maximum(data)
	}
}

func BenchmarkMaxChunks(b *testing.B) {
	data := generateRandomElements(1_000_000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxChunks(data)
	}
}
