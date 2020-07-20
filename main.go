package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Usage:
// go run main.go < data/log.txt
func main() {
	var domains []string
	var lines int

	total := 0
	sum := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines++

		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		// putting the domains in a slice keeps the output in order
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		total += visits
		sum[domain] += visits

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 41))

	// loop over the slice instead of the map to keep them in (sorted) order
	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
