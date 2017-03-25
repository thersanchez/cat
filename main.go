package main

import (
	"fmt"
	"os"
)

func main() {
	fileName, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(fileName)
	os.Exit(0)
}

func parseArgs() (string, error) {
	return "", fmt.Errorf("TODO")
}
