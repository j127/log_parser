package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Usage:
// go run main.go < data/log.txt
func main() {
	var domains []string
	sum := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		domain := fields[0]
		visits, _ := strconv.Atoi(fields[1])

		// putting the domains in a slice keeps the output in order
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		sum[domain] += visits

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 41))

	// loop over the slice instead of the map
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
}
