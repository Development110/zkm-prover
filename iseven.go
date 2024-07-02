package main

import (
	"os"
	"strconv"
)

func main() {
	if i, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
		if (i % 2 == 0) {
			os.Exit(0) // Success
		} else {
			os.Exit(1) // Fail
		}
	}
os.Exit(2) // Fail
}
