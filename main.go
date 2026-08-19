package main

import (
	"fmt"
	"os"
)

// benchmark_suite - Performance benchmarking
func benchmark_suite(path string) {
	fmt.Println("========================================")
	fmt.Println("  Benchmark-Suite")
	fmt.Println("  Performance benchmarking")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	benchmark_suite(path)
}
