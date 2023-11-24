package main

import (
	"fmt"
	"os"

	"github.com/containerd/console"
)

func main() {
	// 1. generate console from stderr
	c, err := console.ConsoleFromFile(os.Stderr)
	if err != nil {
		fmt.Println("failed to generate console from stderr", err)
		os.Exit(1)
	}

	// 2. output to stderr, stdout and console
	fmt.Fprintln(os.Stderr, "via os.Stderr")
	fmt.Fprintln(os.Stdout, "via os.Stdout")
	fmt.Fprintln(c, "via stderr console")
}
