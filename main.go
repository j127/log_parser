package main

import (
	"bufio"
	"fmt"
	"os"
)

// Usage at this stage `./log_parser < data/enchiridion.txt`
func main() {
	in := bufio.NewScanner(os.Stdin)
	lines := 0

	for in.Scan() {
		// fmt.Println("> ", in.Text(), " => ", in.Bytes())
		lines++
	}

	// e.g., this line at the top of the file would cause an error:
	// `os.Stdin.Close()`
	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("lines: ", lines)

}
