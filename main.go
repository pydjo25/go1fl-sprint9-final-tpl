package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range size {
		data[i] = rand.Intn(size * 10)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	results := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := range CHUNKS {
		go func(idx int) {
			defer wg.Done()

			start := idx * len(data) / CHUNKS
			end := (idx + 1) * len(data) / CHUNKS

			if start >= end {
				results[idx] = math.MinInt
				return
			}

			chunk := data[start:end]
			chunkMax := chunk[0]
			for _, num := range chunk {
				if num > chunkMax {
					chunkMax = num
				}
			}
			results[idx] = chunkMax
		}(i)
	}

	wg.Wait()
	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("\nИщем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("\nМаксимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
