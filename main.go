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
	sum := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		domain := fields[0]
		visits, _ := strconv.Atoi(fields[1])
		sum[domain] += visits

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 41))
	for domain, visits := range sum {
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
}
