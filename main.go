package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	username := os.Getenv("USERNAME")
	if username == "" {
		username = os.Getenv("USER")
	}
	fmt.Println("Username:", username)

	fmt.Println("CLI arguments:")
	if len(os.Args) <= 1 {
		fmt.Println("(none)")
	} else {
		for i, arg := range os.Args[1:] {
			fmt.Printf("  [%d] %s\n", i+1, arg)
		}
	}

	fmt.Println("Go version:", runtime.Version())
}
