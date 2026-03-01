package main

import "fmt"

// Version is set during the generation process.
const Version = "generation-2026-03-01-07-20-49"

func main() {
	fmt.Printf("%s version %s\n", "sdd-hello-world", Version)
}
