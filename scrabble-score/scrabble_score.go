package scrabble

import (
	"strings"
)

// // Score function takes in a word and calculates the scrabble score for that word
// func Score(word string) int {

// 	firstGroup := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
// 	secondGroup := []string{"D", "G"}
// 	thirdGroup := []string{"B", "C", "M", "P"}
// 	forthGroup := []string{"F", "H", "V", "W", "Y"}
// 	fifthGroup := []string{"K"}
// 	eigthGroup := []string{"J", "X"}
// 	tenthGroup := []string{"Q", "Z"}

// 	var scrabbleScore int
// 	for _, i := range strings.ToUpper(word) {
// 		for _, j := range firstGroup {
// 			if string(i) == j {
// 				scrabbleScore++
// 			}
// 		}
// 		for _, k := range secondGroup {
// 			if string(i) == k {
// 				scrabbleScore += 2
// 			}
// 		}
// 		for _, l := range thirdGroup {
// 			if string(i) == l {
// 				scrabbleScore += 3
// 			}
// 		}
// 		for _, m := range forthGroup {
// 			if string(i) == m {
// 				scrabbleScore += 4
// 			}
// 		}
// 		for _, n := range fifthGroup {
// 			if string(i) == n {
// 				scrabbleScore += 5
// 			}
// 		}
// 		for _, p := range eigthGroup {
// 			if string(i) == p {
// 				scrabbleScore += 8
// 			}
// 		}
// 		for _, q := range tenthGroup {
// 			if string(i) == q {
// 				scrabbleScore += 10
// 			}
// 		}
// 	}

// 	return scrabbleScore
// }

// Benchmarking

// I have made some suggestions that involve benchmarking (testing the performance and efficiency of your code). If you're already familiar with benchmarking you can skip this.

// Benchmarks already exist for this exercise. You can execute them with:

// go test -v --bench . --benchmem
// This will first run the tests, and then the benchmarks, producing something like this:

// goos: darwin
// goarch: amd64
// pkg: path/to/exercise
// BenchmarkName-8  100000  1859 ns/op  57 B/op  7 allocs/op
// PASS
// ok  path/to/exercise   2.042s
// For each benchmark there will be a line starting with the BenchmarkName followed by the number of cores (8) available. The next number (100000) indicates how many times the benchmark was run. Go will automatically work out how many times to run the benchmark to get a statistically useful result.

// The next three numbers indicate:

// the time it took to execute the benchmark, in ns/op (nanoseconds per operation)

// the memory usage, in B/op (bytes of memory allocated per operation)

// the number of memory allocations, in alloc/op

// For all these numbers, lower is better.

// The last number 2.042s indicates the total execution time for all tests and benchmarks. Ignore this. This is not significant for measuring speed! The benchmarking tool in Go executes faster benchmarks more often to produce more reliable results, possibly increasing the total execution time.

var letVal = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score function takes in a word and calculates the scrabble score for that word
func Score(word string) int {

	var scrabbleScore int
	for _, i := range strings.ToUpper(word) {
		scrabbleScore += letVal[i]
	}

	return scrabbleScore
}
