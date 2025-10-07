package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Current version of Go: %s\n", runtime.Version())
}
