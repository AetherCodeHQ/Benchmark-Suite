package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func bubblePartial(a []int) {
	for i := 0; i < 100 && i < len(a); i++ {
		for j := i + 1; j < 100 && j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func sortSlice(a []int) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}

func linearSearch(a []int) {
	for _, v := range a {
		_ = v
	}
}

func mapLookup(a []int) {
	m := map[int]bool{}
	for _, v := range a {
		m[v] = true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	size := 10000
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size * 10)
	}
	fmt.Printf("Benchmark Suite (n=%d)\n", size)
	fmt.Println("=====================")
	algos := []struct {
		Name string
		Func func([]int)
	}{
		{"Bubble Sort (partial)", bubblePartial},
		{"Sort.Slice", sortSlice},
		{"Linear Search", linearSearch},
		{"Map Lookup", mapLookup},
	}
	for _, algo := range algos {
		cp := make([]int, len(data))
		for i := range data {
			cp[i] = data[i]
		}
		start := time.Now()
		algo.Func(cp)
		elapsed := time.Since(start)
		fmt.Printf("  %-25s %v\n", algo.Name, elapsed)
	}
}
