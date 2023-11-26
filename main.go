package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkoukk/tiktoken-go"
)

func main() {
	encoding := "cl100k_base"
	if len(os.Args) > 1 {
		encoding = os.Args[1]
	}

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	var totalTokens int
	for scanner.Scan() {
		totalTokens += len(tke.Encode(scanner.Text(), nil, nil))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(totalTokens)
}
