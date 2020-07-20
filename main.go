package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Usage at this stage `./log_parser < data/enchiridion.txt`
// or pipe in a webpage from `curl`, etc.
func main() {
	args := os.Args[1:]

	// the words get lowercased
	pattern := regexp.MustCompile(`[^a-z]+`)

	if len(args) != 1 {
		fmt.Println("Please type a search word.")
		return
	}

	query := args[0]

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = pattern.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("The input does not contain %q.\n", query)
}
