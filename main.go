package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

type Suffix struct {
	index  int
	suffix string
}

func naive_suffix_array_construction(s string) []int {
	s += "$" // Add a sentinel character
	n := len(s)
	suffixes := make([]Suffix, n)

	// Create an array of suffixes
	for i := 0; i < n; i++ {
		suffixes[i] = Suffix{i, s[i:]}
	}

	/*
		Sort the suffixes lexicographically.
		For each pair of suffixes, compare them character by character.

		If the characters are equal, continue to the next character.

		If the characters are not equal, return the result of the comparison.

		If the characters are equal up to the length of the shorter suffix, return the result
		of the comparison of the lengths of the suffixes.
	*/
	sort.Slice(suffixes, func(i, j int) bool {
		for k := 0; k < len(suffixes[i].suffix) && k < len(suffixes[j].suffix); k++ {
			if suffixes[i].suffix[k] == suffixes[j].suffix[k] {
				continue
			} else if suffixes[i].suffix[k] < suffixes[j].suffix[k] {
				return true
			} else if suffixes[i].suffix[k] > suffixes[j].suffix[k] {
				return false
			}
		}
		return len(suffixes[i].suffix) < len(suffixes[j].suffix)
	})

	// Create the suffix array from the sorted suffixes
	sa := make([]int, n)
	for i, suf := range suffixes {
		sa[i] = suf.index
	}

	return sa
}

func print_suffix_array_of_string(s string) {
	suffixArray := naive_suffix_array_construction(s)
	fmt.Println("Suffix Array:", suffixArray)
}

func print_time_taken_to_build_suffix_array(s string) {
	start := time.Now()
	naive_suffix_array_construction(s)
	fmt.Println("Time taken to build suffix array:", time.Since(start))
}

func load_random_string(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return ""
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return ""
	}

	fileString := string(fileContent)
	return fileString
}

func get_a_string(size int) string {
	a_string := ""
	for i := 0; i < size; i++ {
		a_string += "a"
	}
	return a_string
}

func main() {
	// Example string
	print_suffix_array_of_string("banana")
	/*
		Expected output when creating the suffix array of "banana":

		Suffix Array: [6 5 3 1 0 4 2]
	*/

	// Load the 10000 length randomstring from the file
	random_string := load_random_string("randomstring.txt")

	// Print the time taken to build the suffix array of the random string
	print_time_taken_to_build_suffix_array(random_string)

	// Print the time taken to build the suffix array of a 10000 length string
	print_time_taken_to_build_suffix_array(get_a_string(10000))

	/*
		On my machine the average time taken to build the suffix array of the random string is 4.5ms
		and the average time taken to build the suffix array of 10000 a's is 90ms.

		The comparison function when comparing suffixes that are all a's will always return false,
		so the time taken to compare each string is always O(n) and with O(n logn) sorting,
		the time taken to build the suffix array of a string of all a's is always O(n^2 log n).

		The comparison function when comparing random strings will at some point find a difference
		between the characters of the suffixes, so the time taken to compare each string is way less -
		on average - than O(n).

		The worst case for a random string is still O(n^2 log n) but with string generated from
		4 characters, it's very unlikely that this will be hit.


		These facts are the reason why the time taken to build the suffix array of a random string
		is way less than the time taken to build the suffix array of a string of all a's.
	*/

}
